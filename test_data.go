package main

import (
	columntypes "dbf_exporter/column_types"
	"dbf_exporter/models"
)

func getTestColumns() []models.ColumnInfo {
	ret := make([]models.ColumnInfo, 2)
	ret[0] = models.ColumnInfo{Type: columntypes.Number, Name: "id", Size: 10, Precision: 0}
	ret[1] = models.ColumnInfo{Type: columntypes.String, Name: "name", Size: 100}
	return ret
}

func getTestData() [][]string {
	ret := make([][]string, 5)

	ret[0] = make([]string, 2)
	ret[0][0] = "1"
	ret[0][1] = "name1"

	ret[1] = make([]string, 2)
	ret[1][0] = "2"
	ret[1][1] = "імя2"

	ret[2] = make([]string, 2)
	ret[2][0] = "3"
	ret[2][1] = "name3"

	ret[3] = make([]string, 2)
	ret[3][0] = "4"
	ret[3][1] = "name4"

	ret[4] = make([]string, 2)
	ret[4][0] = "5"
	ret[4][1] = "name5"

	return ret
}
