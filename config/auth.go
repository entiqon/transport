package config

// Auth defines authentication configuration used to resolve credentials.
type Auth struct {
	// Kind identifies the authentication mechanism.
	//
	// Examples:
	//   oauth2
	//   basic
	//   apikey
	//   bearer
	Kind string `mapstructure:"kind" yaml:"kind"`

	// OAuth2 contains OAuth2-specific configuration when Kind is "oauth2".
	OAuth2 *OAuth2 `mapstructure:"oauth2,omitempty" yaml:"oauth2,omitempty"`
}
