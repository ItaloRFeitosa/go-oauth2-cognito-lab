package auth

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	jwt.StandardClaims

	ClientID string      `json:"client_id,omitempty"`
	Scopes   *ScopeClaim `json:"scope,omitempty"`
}

func (c *CustomClaims) HasScopes(scopes ...string) bool {
	return c.Scopes.HasScopes(scopes...)
}

type ScopeClaim struct {
	scopes map[string]bool
}

func (claim *ScopeClaim) MarshalJSON() ([]byte, error) {
	return []byte(claim.String()), nil
}

func (claim *ScopeClaim) String() string {
	var concatened string
	for scope := range claim.scopes {
		concatened = fmt.Sprintf("%s %s", concatened, scope)
	}

	return concatened
}

func (s *ScopeClaim) UnmarshalJSON(b []byte) error {
	var asString string

	s.scopes = make(map[string]bool)

	if err := json.Unmarshal(b, &asString); err != nil {
		return err
	}

	for _, scope := range strings.Split(asString, " ") {
		if scope != "" {
			s.scopes[scope] = true
		}
	}

	return nil
}

func (s *ScopeClaim) HasScopes(scopes ...string) bool {
	for _, scope := range scopes {
		if !s.HasScope(scope) {
			return false
		}
	}

	return true
}

func (s *ScopeClaim) HasScope(scope string) bool {
	_, ok := s.scopes[scope]

	return ok
}
