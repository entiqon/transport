package config

import "github.com/entiqon/transport/errors"

// InvalidAuthConfigError indicates that the authentication
// configuration is invalid or incomplete.
var InvalidAuthConfigError = &errors.Error{
	Code:    "auth_config",
	Message: "invalid authentication configuration",
}

// InvalidOAuth2ConfigError indicates that the OAuth2 configuration
// is missing required fields or is otherwise invalid.
var InvalidOAuth2ConfigError = &errors.Error{
	Code:    "oauth2_config",
	Message: "OAuth2 configuration is invalid",
}
