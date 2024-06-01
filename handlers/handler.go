package handlers

import (
	"context"
	"net/http"
	"postgres-metadata/services"

	"github.com/gin-gonic/gin"
)

// MetadataHandler handles requests for database metadata
type MetadataHandler struct {
	service *services.MetadataService
}

// NewMetadataHandler creates a new handler instance
func NewMetadataHandler(service *services.MetadataService) *MetadataHandler {
	return &MetadataHandler{service: service}
}

func (h *MetadataHandler) GetTables(ctx *gin.Context) {
	schema := ctx.Param("schema")
	tables, err := h.service.GetTables(context.Background(), schema)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tables)
}

func (h *MetadataHandler) GetColumns(ctx *gin.Context) {
	schema := ctx.Param("schema")
	table := ctx.Param("table")

	column, err := h.service.GetTableColumns(context.Background(), schema, table)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, column)

}

func (h *MetadataHandler) GetIndexes(ctx *gin.Context) {
	schema := ctx.Param("schema")

	indexes, err := h.service.GetIndexes(context.Background(), schema)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, indexes)
}

func (h *MetadataHandler) GetConstraints(ctx *gin.Context) {
	schema := ctx.Param("schema")
	table := ctx.Param("table")

	constraints, err := h.service.GetConstraints(context.Background(), schema, table)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, constraints)
}

func (h *MetadataHandler) GetSchemas(ctx *gin.Context) {
	schemas, err := h.service.GetSchemas(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, schemas)
}

func (h *MetadataHandler) GetViews(ctx *gin.Context) {
	schema := ctx.Param("schema")

	views, err := h.service.GetViews(context.Background(), schema)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, views)
}

func (h *MetadataHandler) GetSequences(ctx *gin.Context) {
	schema := ctx.Param("schema")

	sequences, err := h.service.GetSequences(context.Background(), schema)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sequences)
}

func (h *MetadataHandler) GetFunctions(ctx *gin.Context) {
	schema := ctx.Param("schema")

	functions, err := h.service.GetFunctions(context.Background(), schema)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, functions)
}

func (h *MetadataHandler) GetTriggers(ctx *gin.Context) {
	schema := ctx.Param("schema")

	triggers, err := h.service.GetTriggers(context.Background(), schema)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, triggers)
}

func (h *MetadataHandler) GetRoles(ctx *gin.Context) {
	roles, err := h.service.GetRoles(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, roles)
}
