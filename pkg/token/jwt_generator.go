package token

import (
	"context"
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.elastic.co/apm/v2"
)

type JWTGenerator struct {
	privateKey *rsa.PrivateKey
	ttl        time.Duration
	issuer     string
	subject    string
}

func NewJWTGenerator(privateKey []byte, issuer string, subject string, ttl time.Duration) (*JWTGenerator, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, err
	}

	gen := &JWTGenerator{
		privateKey: key,
		issuer:     issuer,
		subject:    subject,
		ttl:        ttl,
	}
	return gen, nil
}

func (g *JWTGenerator) Generate(ctx context.Context, data Data) (string, error) {
	span, _ := apm.StartSpan(ctx, "Generate", "TokenGenerator")
	defer span.End()

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["iss"] = g.issuer
	claims["sub"] = g.subject
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(g.ttl).Unix()
	claims["dat"] = data.toMap()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(g.privateKey)
	return token, err
}
