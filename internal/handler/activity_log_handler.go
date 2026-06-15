package handler

import (
	"umkm-odod/helper"
	"umkm-odod/internal/constants"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/service"

	"github.com/gin-gonic/gin"
)

// no interface, langsung struct implementasi
type ActivityLogHandler struct {
	service service.ActivityLogService
}

// consturctor
func NewActivityLogHandler(service service.ActivityLogService) *ActivityLogHandler {
	return &ActivityLogHandler{
		service: service,
	}
}

// struct method
func (h *ActivityLogHandler) GetAllActivityLog(c *gin.Context) {
	// query URL
	var query dto.GetAllActivityLogResponseQuery
	err := c.ShouldBindQuery(&query)

	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// pagination with helper
	helper.NormalizePagination(&query.Page, &query.Limit)

	logs, total, err := h.service.GetAllActivityLog(c.Request.Context(), query)
	if err != nil {
		helper.ErrorResponse(c, constants.ErrorGetData, err)
		return
	}

	helper.SuccessGetAllLogsPerTenant(c, logs, int(total), query.Page, query.Limit)
}
