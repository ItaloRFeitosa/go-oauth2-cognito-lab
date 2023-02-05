package auth

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"

	"github.com/MicahParks/keyfunc"
)

const (
	FiberJWTContextKey = "token"
)

var (
	OAuth2JWKsEndpoint = os.Getenv("OAUTH2_JWKS_ENDPOINT")
)

func FiberMiddleware() fiber.Handler {
	options := keyfunc.Options{
		Ctx: context.Background(),
		RefreshErrorHandler: func(err error) {
			log.Printf("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}

	jwks, err := keyfunc.Get(OAuth2JWKsEndpoint, options)
	if err != nil {
		log.Fatalf("Failed to create JWKS from resource at the given URL.\nError: %s", err.Error())
	}

	return jwtware.New(jwtware.Config{
		ContextKey: FiberJWTContextKey,
		KeyFunc:    jwks.Keyfunc,
		Claims:     &CustomClaims{},
	})
}

func FiberEnsureScopes(scopes ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := FiberGetToken(c)

		// log.Printf("token.Raw: %#+v\n", token.Raw)

		claims := token.Claims.(*CustomClaims)
		if !claims.HasScopes(scopes...) {
			return fiber.NewError(http.StatusForbidden, "insufficient scopes")
		}

		return c.Next()
	}
}

func FiberInjectSubject(c *fiber.Ctx) error {
	c.SetUserContext(NewContextWithSubject(c.UserContext(), FiberGetSubject(c)))

	return c.Next()
}

func FiberGetToken(c *fiber.Ctx) *jwt.Token {
	return c.Locals(FiberJWTContextKey).(*jwt.Token)
}

func FiberGetSubject(c *fiber.Ctx) string {
	return FiberGetToken(c).Claims.(*CustomClaims).Subject
}
