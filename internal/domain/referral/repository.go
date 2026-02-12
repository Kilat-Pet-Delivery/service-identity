package referral

import (
	"context"

	"github.com/google/uuid"
)

// ReferralRepository defines persistence operations for referrals.
type ReferralRepository interface {
	Save(ctx context.Context, r *Referral) error
	Update(ctx context.Context, r *Referral) error
	FindByReferrerID(ctx context.Context, referrerID uuid.UUID) ([]*Referral, error)
	FindByReferralCode(ctx context.Context, code string) (*Referral, error)
	FindByRefereeID(ctx context.Context, refereeID uuid.UUID) (*Referral, error)
	CountByReferrerID(ctx context.Context, referrerID uuid.UUID) (int64, error)
	SaveUserReferralCode(ctx context.Context, userID uuid.UUID, code string) error
	GetUserReferralCode(ctx context.Context, userID uuid.UUID) (string, error)
	FindUserIDByReferralCode(ctx context.Context, code string) (uuid.UUID, error)
}
