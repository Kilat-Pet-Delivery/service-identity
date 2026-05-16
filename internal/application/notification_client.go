package application

import (
	"context"

	"go.uber.org/zap"
)

// PasswordResetNotifier sends password reset notifications.
// TODO: replace with Kafka-backed notifier publishing to identity.events when that topic exists.
type PasswordResetNotifier interface {
	SendPasswordResetEmail(ctx context.Context, email, token string) error
}

// LogOnlyPasswordResetNotifier is a stub notifier that logs the event without sending.
type LogOnlyPasswordResetNotifier struct {
	logger *zap.Logger
}

// NewLogOnlyPasswordResetNotifier creates a new LogOnlyPasswordResetNotifier.
func NewLogOnlyPasswordResetNotifier(logger *zap.Logger) *LogOnlyPasswordResetNotifier {
	return &LogOnlyPasswordResetNotifier{logger: logger}
}

// SendPasswordResetEmail logs the reset email event without sending.
func (n *LogOnlyPasswordResetNotifier) SendPasswordResetEmail(ctx context.Context, email, token string) error {
	n.logger.Info("password reset email enqueued (log-only)", zap.String("email", email))
	return nil
}
