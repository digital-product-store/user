package token

import "context"

type TokenGenerator interface {
	Generate(ctx context.Context, data Data) (string, error)
}
