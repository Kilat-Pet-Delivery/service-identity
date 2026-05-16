package identity

import (
	"time"

	"github.com/google/uuid"
)

// PasswordReset represents a single-use password reset token linked to a user.
type PasswordReset struct {
	id        uuid.UUID
	userID    uuid.UUID
	token     string
	expiresAt time.Time
	usedAt    *time.Time
	createdAt time.Time
}

// NewPasswordReset creates a new PasswordReset for a given user.
func NewPasswordReset(userID uuid.UUID, token string, expiresAt time.Time) *PasswordReset {
	return &PasswordReset{
		id:        uuid.New(),
		userID:    userID,
		token:     token,
		expiresAt: expiresAt,
		usedAt:    nil,
		createdAt: time.Now().UTC(),
	}
}

// ReconstructPasswordReset rebuilds a PasswordReset from persistence data.
func ReconstructPasswordReset(
	id, userID uuid.UUID,
	token string,
	expiresAt time.Time,
	usedAt *time.Time,
	createdAt time.Time,
) *PasswordReset {
	return &PasswordReset{
		id:        id,
		userID:    userID,
		token:     token,
		expiresAt: expiresAt,
		usedAt:    usedAt,
		createdAt: createdAt,
	}
}

// --- Getters ---

// ID returns the reset token's unique identifier.
func (p *PasswordReset) ID() uuid.UUID { return p.id }

// UserID returns the owning user's ID.
func (p *PasswordReset) UserID() uuid.UUID { return p.userID }

// Token returns the reset token string.
func (p *PasswordReset) Token() string { return p.token }

// ExpiresAt returns the expiration timestamp.
func (p *PasswordReset) ExpiresAt() time.Time { return p.expiresAt }

// UsedAt returns when the token was used, or nil if unused.
func (p *PasswordReset) UsedAt() *time.Time { return p.usedAt }

// CreatedAt returns the creation timestamp.
func (p *PasswordReset) CreatedAt() time.Time { return p.createdAt }

// --- Behavior ---

// IsExpired checks whether the token has expired.
func (p *PasswordReset) IsExpired() bool {
	return time.Now().UTC().After(p.expiresAt)
}

// IsUsed returns true if the token has already been used.
func (p *PasswordReset) IsUsed() bool {
	return p.usedAt != nil
}

// IsValid returns true if the token is neither used nor expired.
func (p *PasswordReset) IsValid() bool {
	return !p.IsUsed() && !p.IsExpired()
}

// MarkUsed sets the usedAt timestamp to the current UTC time.
func (p *PasswordReset) MarkUsed() {
	now := time.Now().UTC()
	p.usedAt = &now
}
