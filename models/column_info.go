package models

import (
	"csv2dbf/column_types"
)

type ColumnInfo struct {
	Name      string
	Type      columnTypes.ColumnDataType
	Size      byte
	Precision byte
}
