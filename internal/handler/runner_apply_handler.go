package handler

import (
	"context"

	"github.com/Kilat-Pet-Delivery/lib-common/response"
	"github.com/Kilat-Pet-Delivery/lib-proto/dto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RunnerApplyService defines the application-layer contract the handler depends on.
type RunnerApplyService interface {
	Apply(ctx context.Context, req dto.RunnerApplicationRequest) (string, error)
}

// RunnerApplyHandler handles POST /runners/apply.
type RunnerApplyHandler struct {
	service RunnerApplyService
	logger  *zap.Logger
}

// NewRunnerApplyHandler creates a new RunnerApplyHandler.
func NewRunnerApplyHandler(service RunnerApplyService, logger *zap.Logger) *RunnerApplyHandler {
	return &RunnerApplyHandler{
		service: service,
		logger:  logger,
	}
}

// RegisterRoutes registers runner apply routes on the given router group.
// No auth middleware — this is a public endpoint.
func (h *RunnerApplyHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.Group("/runners").POST("/apply", h.Apply)
}

// Apply handles POST /runners/apply.
func (h *RunnerApplyHandler) Apply(c *gin.Context) {
	var req dto.RunnerApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	applicationID, err := h.service.Apply(c.Request.Context(), req)
	if err != nil {
		h.logger.Warn("runner apply failed", zap.Error(err))
		response.Error(c, err)
		return
	}

	response.Created(c, gin.H{"applicationId": applicationID})
}
