package identity

import (
	"time"

	"github.com/google/uuid"
)

// RefreshToken represents a refresh token entity linked to a user.
type RefreshToken struct {
	id        uuid.UUID
	userID    uuid.UUID
	token     string
	expiresAt time.Time
	revoked   bool
	createdAt time.Time
}

// NewRefreshToken creates a new RefreshToken.
func NewRefreshToken(userID uuid.UUID, token string, expiresAt time.Time) *RefreshToken {
	return &RefreshToken{
		id:        uuid.New(),
		userID:    userID,
		token:     token,
		expiresAt: expiresAt,
		revoked:   false,
		createdAt: time.Now().UTC(),
	}
}

// ReconstructRefreshToken rebuilds a RefreshToken from persistence data.
func ReconstructRefreshToken(
	id, userID uuid.UUID,
	token string,
	expiresAt time.Time,
	revoked bool,
	createdAt time.Time,
) *RefreshToken {
	return &RefreshToken{
		id:        id,
		userID:    userID,
		token:     token,
		expiresAt: expiresAt,
		revoked:   revoked,
		createdAt: createdAt,
	}
}

// --- Getters ---

// ID returns the token's unique identifier.
func (t *RefreshToken) ID() uuid.UUID { return t.id }

// UserID returns the owning user's ID.
func (t *RefreshToken) UserID() uuid.UUID { return t.userID }

// Token returns the token string.
func (t *RefreshToken) Token() string { return t.token }

// ExpiresAt returns the expiration timestamp.
func (t *RefreshToken) ExpiresAt() time.Time { return t.expiresAt }

// Revoked returns whether the token has been revoked.
func (t *RefreshToken) Revoked() bool { return t.revoked }

// CreatedAt returns the creation timestamp.
func (t *RefreshToken) CreatedAt() time.Time { return t.createdAt }

// --- Behavior ---

// Revoke marks the token as revoked.
func (t *RefreshToken) Revoke() {
	t.revoked = true
}

// IsExpired checks whether the token has expired.
func (t *RefreshToken) IsExpired() bool {
	return time.Now().UTC().After(t.expiresAt)
}

// IsValid returns true if the token is neither revoked nor expired.
func (t *RefreshToken) IsValid() bool {
	return !t.revoked && !t.IsExpired()
}
