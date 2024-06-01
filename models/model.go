package models

// Table describes a table in the database
type Table struct {
	TableName string `json:"table_name"`

}

// Column describes a column within a table
type Column struct {
	ColumnName string `json:"column_name"`
	DataType   string `json:"data_type"`
	IsNullable bool   `json:"is_nullable"`
	IsPrimary  bool   `json:"is_primary"`
}


// Index model
type Index struct {
	IndexName  string `json:"index_name"`
	TableName  string `json:"table_name"`
	IndexDef   string `json:"index_def"`
}
// Constraint model
type Constraint struct {
	ConstraintName string `json:"constraint_name"`
	TableName      string `json:"table_name"`
	ConstraintType string `json:"constraint_type"`
}

// Schema model
type Schema struct {
	SchemaName string `json:"schema_name"`
}
// View model
type View struct {
	ViewName string `json:"view_name"`
	ViewDefinition string `json:"view_definition"`
}

// Sequence model
type Sequence struct {
	SequenceName string `json:"sequence_name"`
	DataType string `json:"data_type"`
	StartValue int `json:"start_value"`
	MinimumValue int `json:"minimum_value"`
	MaximumValue int `json:"maximum_value"`
	IncrementBy int `json:"increment_by"`
}
// Function model
type Function struct {
	FunctionName      string `json:"function_name"`
	FunctionDefinition string `json:"function_definition"`
}

// Trigger model
type Trigger struct {
	TriggerName      string `json:"trigger_name"`
	EventManipulation string `json:"event_manipulation"`
	EventObjectTable  string `json:"event_object_table"`
	ActionStatement   string `json:"action_statement"`
}


type Role struct {
	Rolname  string `json:"rolname"`
	Rolsuper bool   `json:"rolsuper"`
	Rolinherit bool `json:"rolinherit"`
	Rolcreaterole bool `json:"rolcreaterole"`
	Rolcreatedb bool `json:"rolcreatedb"`
	Rolcanlogin bool `json:"rolcanlogin"`
	Rolreplication bool `json:"rolreplication"`
	Rolbypassrls bool `json:"rolbypassrls"`
}
