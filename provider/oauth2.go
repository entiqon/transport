package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/entiqon/transport/auth"
	"github.com/entiqon/transport/config"
	"github.com/entiqon/transport/credential"
)

// oauth2Provider resolves OAuth2 authentication configuration into a
// credential usable by the transport client.
//
// The provider retrieves OAuth2 access tokens from the configured token
// endpoint and returns a BearerToken credential that can be applied to
// outgoing HTTP requests. Tokens are cached and automatically refreshed
// when they expire.
type oauth2Provider struct {
	http *http.Client

	mu         sync.Mutex
	credential auth.Credential
	expiresAt  time.Time
}

// OAuth2 creates an OAuth2 authentication provider.
//
// The returned provider resolves OAuth2 configuration into credentials
// during request execution. If client is nil, http.DefaultClient is used.
func OAuth2(client *http.Client) auth.Provider {

	if client == nil {
		client = http.DefaultClient
	}

	return &oauth2Provider{
		http: client,
	}
}

// Refresh invalidates the cached OAuth2 credential.
//
// Does not immediately request a new token. Instead, it clears the
// cached expiration so that the next call to Resolve triggers a token
// refresh from the authorization server.
//
// This method is typically invoked when a transport client receives an
// HTTP 401 Unauthorized response and needs to force credential renewal.
func (p *oauth2Provider) Refresh(ctx context.Context) error {

	p.mu.Lock()
	defer p.mu.Unlock()

	p.expiresAt = time.Time{}
	p.credential = nil

	return nil
}

// Resolve implements auth.Provider.
//
// Resolve returns a credential constructed from an OAuth2 access token.
// The provider caches the credential and refreshes it automatically when
// the token expires.
func (p *oauth2Provider) Resolve(
	ctx context.Context,
	cfg config.Auth,
) (auth.Credential, error) {

	if cfg.OAuth2 == nil {
		return nil, config.InvalidOAuth2ConfigError
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	now := time.Now()

	if p.credential == nil || now.After(p.expiresAt) {

		cred, exp, refreshToken, err := p.requestToken(ctx, cfg.OAuth2)
		if err != nil {
			return nil, err
		}

		p.credential = cred
		p.expiresAt = exp

		if refreshToken != "" {
			cfg.OAuth2.RefreshToken = refreshToken
		}
	}

	return p.credential, nil
}

// requestToken resolves the OAuth2 token based on the configured grant type.
//
// The grant type acts as a discriminator that determines which OAuth2
// flow should be executed.
func (p *oauth2Provider) requestToken(
	ctx context.Context,
	cfg *config.OAuth2,
) (auth.Credential, time.Time, string, error) {

	switch cfg.GrantType {

	case "refresh_token":
		return p.refreshToken(ctx, cfg)

	case "client_credentials":
		return p.clientCredentials(ctx, cfg)

	default:
		return nil, time.Time{}, "", config.InvalidOAuth2ConfigError
	}
}

// refreshToken implements the OAuth2 refresh_token grant flow.
func (p *oauth2Provider) refreshToken(
	ctx context.Context,
	cfg *config.OAuth2,
) (auth.Credential, time.Time, string, error) {

	form := url.Values{}
	form.Set("grant_type", cfg.GrantType)
	form.Set("refresh_token", cfg.RefreshToken)
	form.Set("client_id", cfg.ClientID)
	form.Set("client_secret", cfg.ClientSecret)

	return p.request(ctx, cfg, form)
}

// clientCredentials implements the OAuth2 client_credentials grant flow.
func (p *oauth2Provider) clientCredentials(
	ctx context.Context,
	cfg *config.OAuth2,
) (auth.Credential, time.Time, string, error) {

	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	form.Set("client_id", cfg.ClientID)
	form.Set("client_secret", cfg.ClientSecret)

	if cfg.Scope != "" {
		form.Set("scope", cfg.Scope)
	}

	return p.request(ctx, cfg, form)
}

// request performs the OAuth2 token HTTP request to the configured
// authorization server.
func (p *oauth2Provider) request(
	ctx context.Context,
	cfg *config.OAuth2,
	form url.Values,
) (auth.Credential, time.Time, string, error) {

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		cfg.TokenURL,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return nil, time.Time{}, "", err
	}

	contentType := cfg.ContentType
	if contentType == "" {
		contentType = "application/x-www-form-urlencoded"
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := p.http.Do(req)
	if err != nil {
		return nil, time.Time{}, "", err
	}
	defer resp.Body.Close()

	var r oauth2Response

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, time.Time{}, "", err
	}

	cred := credential.BearerToken(r.AccessToken)

	buffer := 5
	if r.ExpiresIn <= buffer {
		buffer = 1
	}

	expires := time.Now().Add(time.Duration(r.ExpiresIn-buffer) * time.Second)

	return cred, expires, r.RefreshToken, nil
}

// oauth2Response represents the OAuth2 token endpoint response returned
// by the authorization server.
type oauth2Response struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
