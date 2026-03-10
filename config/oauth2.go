package config

// OAuth2 defines configuration required for OAuth2 authentication.
type OAuth2 struct {
	// GrantType defines the OAuth2 grant type.
	GrantType string `mapstructure:"grant_type" yaml:"grant_type"`

	// TokenURL is the endpoint used to obtain or refresh tokens.
	TokenURL string `mapstructure:"token_url" yaml:"token_url"`

	// ClientID identifies the OAuth2 client.
	ClientID string `mapstructure:"client_id" yaml:"client_id"`

	// ClientSecret authenticates the OAuth2 client.
	ClientSecret string `mapstructure:"client_secret" yaml:"client_secret"`

	// RefreshToken is used when the grant type is refresh_token.
	RefreshToken string `mapstructure:"refresh_token,omitempty" yaml:"refresh_token,omitempty"`

	// Scope defines optional OAuth2 scopes.
	Scope string `mapstructure:"scope,omitempty" yaml:"scope,omitempty"`

	// ContentType defines the request content type used when requesting
	// tokens from the OAuth2 token endpoint. If empty, the default
	// "application/x-www-form-urlencoded" is used.
	ContentType string `mapstructure:"content_type,omitempty" yaml:"content_type,omitempty"`
}
