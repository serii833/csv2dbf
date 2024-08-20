package fileLoader

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"csv2dbf/column_types"
	"csv2dbf/models"
)

func parseToByte(s string) byte {
	val, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		panic(err)
	}
	return byte(val)
}

func parseheaderColumns(headerLine string, sep string) []models.ColumnInfo {
	columns := make([]models.ColumnInfo, 0)
	fields := strings.Split(headerLine, sep)
	for _, f := range fields {
		typeInfoStartIndex := strings.IndexRune(f, '[')
		typeInfoEndIndex := strings.IndexRune(f, ']')

		typeInfoString := f[typeInfoStartIndex+1 : typeInfoEndIndex]

		colInfo := models.ColumnInfo{}
		colInfo.Name = strings.Trim(f[0:typeInfoStartIndex], " ")

		typeInfoParts := strings.Split(typeInfoString, " ")

		switch strings.ToUpper(typeInfoParts[0]) {
		case "C":
			colInfo.Type = columnTypes.String
			colInfo.Size = parseToByte(typeInfoParts[1])
		case "N":
			colInfo.Type = columnTypes.Number
			colInfo.Size = parseToByte(typeInfoParts[1])
			colInfo.Precision = parseToByte(typeInfoParts[2])
		}
		columns = append(columns, colInfo)
	}
	return columns
}

func LoadFile(pathToFile string, sep string) ([]models.ColumnInfo, []string) {
	fileData, err := os.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}

	fileText := string(fileData)

	linesTmp := strings.Split(fileText, "\n")
	lines := slices.DeleteFunc(linesTmp, func(c string) bool { return c == "" })

	columns := parseheaderColumns(lines[0], sep)
	dataLines := lines[1:]

	return columns, dataLines
}
