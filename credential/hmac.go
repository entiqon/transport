package credential

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/entiqon/transport/auth"
)

// hmacCredential implements HMAC request signing using SHA256.
type hmacCredential struct {
	key    string
	secret string
}

// HMAC creates a credential strategy that signs outgoing HTTP
// requests using HMAC-SHA256.
//
// The credential injects the following headers:
//
//	X-Key: <key>
//	X-Timestamp: <unix timestamp>
//	X-Signature: <signature>
//
// The signature is computed as:
//
//	HMAC-SHA256(secret, METHOD + "\n" + PATH + "\n" + TIMESTAMP)
func HMAC(key, secret string) auth.Credential {
	return &hmacCredential{
		key:    key,
		secret: secret,
	}
}

// Apply signs the outgoing HTTP request and injects HMAC headers.
func (c *hmacCredential) Apply(ctx context.Context, req *http.Request) error {

	if c.key == "" {
		return errors.New("credential: missing HMAC key")
	}

	if c.secret == "" {
		return errors.New("credential: missing HMAC secret")
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	canonical := req.Method + "\n" + req.URL.Path + "\n" + timestamp

	mac := hmac.New(sha256.New, []byte(c.secret))
	mac.Write([]byte(canonical))

	signature := hex.EncodeToString(mac.Sum(nil))

	req.Header.Set("X-Key", c.key)
	req.Header.Set("X-Timestamp", timestamp)
	req.Header.Set("X-Signature", signature)

	return nil
}
