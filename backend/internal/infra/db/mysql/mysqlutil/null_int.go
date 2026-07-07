package mysqlutil

import "database/sql"

// NullInt64ToPtr は sql.NullInt64 を *int に変換します。
// 値が NULL の場合は nil を返します。
func NullInt64ToPtr(v sql.NullInt64) *int {
	if !v.Valid {
		return nil
	}
	value := int(v.Int64)
	return &value
}
