package auth

import (
	"github.com/2637309949/micro/v3/service/auth"
)

const (
	// BearerScheme used for Authorization header
	BearerScheme = "Bearer "
	// TokenCookieName is the name of the cookie which stores the auth token
	TokenCookieName = "micro-token"
)

// SystemRules are the default rules which are applied to the runtime services
var SystemRules = []*auth.Rule{
	{
		ID:       "default",
		Scope:    auth.ScopePublic,
		Resource: &auth.Resource{Type: "*", Name: "*", Endpoint: "*"},
		Access:   auth.AccessGranted,
		Priority: 0,
	},
}
