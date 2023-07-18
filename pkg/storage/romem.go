package storage

import (
	"context"

	"go.elastic.co/apm/v2"
)

type ROMemStorage struct {
	users []User
}

func (r ROMemStorage) GetByUUID(ctx context.Context, uuid string) *User {
	span, ctx := apm.StartSpan(ctx, "GetByUUID", "storage")
	defer span.End()

	for _, user := range r.users {
		if user.Uuid == uuid {
			return &user
		}
	}

	return nil
}

func (r ROMemStorage) GetByUsername(ctx context.Context, username string) *User {
	span, ctx := apm.StartSpan(ctx, "GetByUsername", "storage")
	defer span.End()

	for _, user := range r.users {
		if user.Username == username {
			return &user
		}
	}

	return nil
}

func NewROMemStorage(users []User) ROMemStorage {
	return ROMemStorage{
		users: users,
	}
}
