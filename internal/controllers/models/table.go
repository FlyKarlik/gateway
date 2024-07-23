package models

// Table struct for table
type Table struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Alias              string `json:"alias"`
	IsRelated          bool   `json:"is_related"`
	IsVersioned        bool   `json:"is_versioned"`
	IsArchived         bool   `json:"is_archived"`
	IsGeometryNullable bool   `json:"is_geometry_nullable"`
	GeometryType       string `json:"geometry_type"`
	SRID               int32  `json:"srid"`
	TableType          string `json:"table_type"`
}
