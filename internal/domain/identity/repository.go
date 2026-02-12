package identity

import (
	"context"

	"github.com/google/uuid"
)

// UserRepository defines persistence operations for User aggregates.
type UserRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Save(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	ListAll(ctx context.Context, page, limit int) ([]*User, int64, error)
	CountByRole(ctx context.Context) (map[string]int64, error)
}

// TokenRepository defines persistence operations for RefreshToken entities.
type TokenRepository interface {
	Save(ctx context.Context, token *RefreshToken) error
	FindByToken(ctx context.Context, token string) (*RefreshToken, error)
	RevokeAllForUser(ctx context.Context, userID uuid.UUID) error
}
