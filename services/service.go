package services

import (
	"context"
	"postgres-metadata/models"
	"postgres-metadata/repository"

)

// MetadataService provides methods for accessing database metadata
type MetadataService struct {
  repo repository.PostgresRepository
}

// NewMetadataService creates a new service instance
func NewMetadataService(repo repository.PostgresRepository) *MetadataService {
  return &MetadataService{repo: repo}
}

// func (s *MetadataService) GetTables(ctx context.Context) ([]models.Table, error) {
// 	return s.repo.GetTables(ctx)
//   }
func (s *MetadataService) GetTables(ctx context.Context, schema string) ([]models.Table, error) {
	return s.repo.GetTables(ctx, schema)
}

  func (s *MetadataService) GetTableColumns(ctx context.Context, schema, table string) ([]models.Column, error) {
	return s.repo.GetColumns(ctx, schema, table)
  }

  func (s *MetadataService) GetIndexes(ctx context.Context, schema string) ([]models.Index, error) {
	return s.repo.GetIndexes(ctx, schema)
}

func (s *MetadataService) GetConstraints(ctx context.Context, schema, table string) ([]models.Constraint, error) {
	return s.repo.GetConstraints(ctx, schema, table)
}

func (s *MetadataService) GetSchemas(ctx context.Context) ([]models.Schema, error) {
	return s.repo.GetSchemas(ctx)
}

func (s *MetadataService) GetViews(ctx context.Context, schema string) ([]models.View, error) {
	return s.repo.GetViews(ctx, schema)
}

func (s *MetadataService) GetSequences(ctx context.Context, schema string) ([]models.Sequence, error) {
	return s.repo.GetSequences(ctx, schema)
}

func (s *MetadataService) GetFunctions(ctx context.Context, schema string) ([]models.Function, error) {
	return s.repo.GetFunctions(ctx, schema)
}

func (s *MetadataService) GetTriggers(ctx context.Context, schema string) ([]models.Trigger, error) {
	return s.repo.GetTriggers(ctx, schema)
}

func (s *MetadataService) GetRoles(ctx context.Context) ([]models.Role, error) {
	return s.repo.GetRoles(ctx)
}