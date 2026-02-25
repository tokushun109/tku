package mysqlutil

import "database/sql"

// NullStringOrEmpty は sql.NullString を string に変換します。
// 値が NULL の場合は空文字列を返します。
func NullStringOrEmpty(v sql.NullString) string {
	if !v.Valid {
		return ""
	}
	return v.String
}
