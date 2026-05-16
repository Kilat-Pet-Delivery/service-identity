package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Kilat-Pet-Delivery/lib-common/domain"
	"github.com/Kilat-Pet-Delivery/service-identity/internal/domain/identity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PasswordResetModel is the GORM model for the password_resets table.
type PasswordResetModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	Token     string     `gorm:"type:text;uniqueIndex;not null"`
	ExpiresAt time.Time  `gorm:"not null"`
	UsedAt    *time.Time `gorm:""`
	CreatedAt time.Time  `gorm:"not null;default:now()"`
}

// TableName specifies the table name for GORM.
func (PasswordResetModel) TableName() string {
	return "password_resets"
}

// toDomain converts a PasswordResetModel to a domain PasswordReset.
func (m *PasswordResetModel) toDomain() *identity.PasswordReset {
	return identity.ReconstructPasswordReset(
		m.ID,
		m.UserID,
		m.Token,
		m.ExpiresAt,
		m.UsedAt,
		m.CreatedAt,
	)
}

// fromDomainPasswordReset converts a domain PasswordReset to a PasswordResetModel.
func fromDomainPasswordReset(p *identity.PasswordReset) *PasswordResetModel {
	return &PasswordResetModel{
		ID:        p.ID(),
		UserID:    p.UserID(),
		Token:     p.Token(),
		ExpiresAt: p.ExpiresAt(),
		UsedAt:    p.UsedAt(),
		CreatedAt: p.CreatedAt(),
	}
}

// GormPasswordResetRepository is a GORM-based implementation of PasswordResetRepository.
type GormPasswordResetRepository struct {
	db *gorm.DB
}

// NewGormPasswordResetRepository creates a new GormPasswordResetRepository.
func NewGormPasswordResetRepository(db *gorm.DB) *GormPasswordResetRepository {
	return &GormPasswordResetRepository{db: db}
}

// Create persists a new password reset token to the database.
func (r *GormPasswordResetRepository) Create(ctx context.Context, reset *identity.PasswordReset) error {
	model := fromDomainPasswordReset(reset)
	return r.db.WithContext(ctx).Create(model).Error
}

// FindByToken retrieves a non-expired password reset token by its token string.
// Returns domain.ErrNotFound if the token does not exist or has expired.
func (r *GormPasswordResetRepository) FindByToken(ctx context.Context, token string) (*identity.PasswordReset, error) {
	var model PasswordResetModel
	err := r.db.WithContext(ctx).
		Where("token = ? AND expires_at > ?", token, time.Now().UTC()).
		First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return model.toDomain(), nil
}
