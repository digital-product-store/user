package storage

import "context"

type Storage interface {
	GetByUUID(ctx context.Context, uuid string) *User
	GetByUsername(ctx context.Context, username string) *User
}
