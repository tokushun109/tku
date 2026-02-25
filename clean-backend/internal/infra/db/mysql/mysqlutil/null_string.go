package mysqlutil

import "database/sql"

// NullStringOrEmpty converts sql.NullString to string.
// It returns empty string when value is NULL.
func NullStringOrEmpty(v sql.NullString) string {
	if !v.Valid {
		return ""
	}
	return v.String
}
