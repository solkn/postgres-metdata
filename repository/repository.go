package repository

import (
	"context"
	"postgres-metadata/models"

	"gorm.io/gorm"
)

type PostgresRepository interface {
	GetTables(ctx context.Context, schema string) ([]models.Table, error)
	GetColumns(ctx context.Context, schema, table string) ([]models.Column, error)
	GetIndexes(ctx context.Context, schema string) ([]models.Index, error)
	GetConstraints(ctx context.Context, schema, table string) ([]models.Constraint, error)
	GetSchemas(ctx context.Context) ([]models.Schema, error)
	GetViews(ctx context.Context, schema string) ([]models.View, error)
	GetSequences(ctx context.Context, schema string) ([]models.Sequence, error)
	GetFunctions(ctx context.Context, schema string) ([]models.Function, error)
	GetTriggers(ctx context.Context, schema string) ([]models.Trigger, error)
	GetRoles(ctx context.Context) ([]models.Role, error)
}

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) PostgresRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetTables(ctx context.Context, schema string) ([]models.Table, error) {
	var tables []models.Table
	result := r.db.WithContext(ctx).Raw(`
	  SELECT table_name, table_schema
	  FROM information_schema.tables
	  WHERE table_schema = $1;
	`, schema).Scan(&tables)
	if result.Error != nil {
		return nil, result.Error
	}

	return tables, nil
}

func (r *postgresRepository) GetColumns(ctx context.Context, schema, table string) ([]models.Column, error) {
	var columns []models.Column
	result := r.db.WithContext(ctx).Raw(`
	  SELECT 
		c.column_name, 
		c.data_type, 
		CASE WHEN c.is_nullable = 'YES' THEN true ELSE false END AS is_nullable,
		CASE WHEN kcu.column_name IS NOT NULL THEN true ELSE false END AS is_primary
	  FROM 
		information_schema.columns AS c
	  LEFT JOIN 
		information_schema.key_column_usage AS kcu
	  ON 
		c.table_schema = kcu.table_schema 
		AND c.table_name = kcu.table_name 
		AND c.column_name = kcu.column_name
		AND kcu.constraint_name IN (
			SELECT constraint_name
			FROM information_schema.table_constraints
			WHERE table_schema = c.table_schema
			AND table_name = c.table_name
			AND constraint_type = 'PRIMARY KEY'
		)
	  WHERE 
		c.table_schema = $1 AND c.table_name = $2;
	`, schema, table).Scan(&columns)
	if result.Error != nil {
		return nil, result.Error
	}

	return columns, nil
}

func (r *postgresRepository) GetIndexes(ctx context.Context, schema string) ([]models.Index, error) {
	var indexes []models.Index
	result := r.db.WithContext(ctx).Raw(`
	  SELECT indexname AS index_name, tablename AS table_name, indexdef AS index_def
	  FROM pg_indexes
	  WHERE schemaname = $1;
	`, schema).Scan(&indexes)
	if result.Error != nil {
		return nil, result.Error
	}

	return indexes, nil
}

func (r *postgresRepository) GetConstraints(ctx context.Context, schema, table string) ([]models.Constraint, error) {
	var constraints []models.Constraint
	result := r.db.WithContext(ctx).Raw(`
	  SELECT constraint_name, table_name, constraint_type
	  FROM information_schema.table_constraints
	  WHERE table_schema = $1 AND table_name = $2;
	`, schema, table).Scan(&constraints)
	if result.Error != nil {
		return nil, result.Error
	}

	return constraints, nil
}

func (r *postgresRepository) GetSchemas(ctx context.Context) ([]models.Schema, error) {
	var schemas []models.Schema
	result := r.db.WithContext(ctx).Raw(`
	  SELECT schema_name
	  FROM information_schema.schemata;
	`).Scan(&schemas)
	if result.Error != nil {
		return nil, result.Error
	}

	return schemas, nil
}

func (r *postgresRepository) GetViews(ctx context.Context, schema string) ([]models.View, error) {
	var views []models.View
	result := r.db.WithContext(ctx).Raw(`
	  SELECT table_name AS view_name, view_definition
	  FROM information_schema.views
	  WHERE table_schema = $1;
	`, schema).Scan(&views)
	if result.Error != nil {
		return nil, result.Error
	}

	return views, nil
}

func (r *postgresRepository) GetSequences(ctx context.Context, schema string) ([]models.Sequence, error) {
	var sequences []models.Sequence
	result := r.db.WithContext(ctx).Raw(`
	  SELECT sequence_name, data_type, start_value::int, minimum_value::int, maximum_value::int, increment_by::int
	  FROM information_schema.sequences
	  WHERE sequence_schema = $1;
	`, schema).Scan(&sequences)
	if result.Error != nil {
		return nil, result.Error
	}

	return sequences, nil
}

func (r *postgresRepository) GetFunctions(ctx context.Context, schema string) ([]models.Function, error) {
	var functions []models.Function
	result := r.db.WithContext(ctx).Raw(`
	  SELECT routine_name AS function_name, routine_definition AS function_definition
	  FROM information_schema.routines
	  WHERE routine_schema = $1 AND routine_type = 'FUNCTION';
	`, schema).Scan(&functions)
	if result.Error != nil {
		return nil, result.Error
	}

	return functions, nil
}

func (r *postgresRepository) GetTriggers(ctx context.Context, schema string) ([]models.Trigger, error) {
	var triggers []models.Trigger
	result := r.db.WithContext(ctx).Raw(`
	  SELECT trigger_name, event_manipulation, event_object_table, action_statement
	  FROM information_schema.triggers
	  WHERE trigger_schema = $1;
	`, schema).Scan(&triggers)
	if result.Error != nil {
		return nil, result.Error
	}

	return triggers, nil
}

func (r *postgresRepository) GetRoles(ctx context.Context) ([]models.Role, error) {
	var roles []models.Role
	result := r.db.WithContext(ctx).Raw(`
	  SELECT rolname, rolsuper, rolinherit, rolcreaterole, rolcreatedb, rolcanlogin, rolreplication, rolbypassrls
	  FROM pg_roles;
	`).Scan(&roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil

}
