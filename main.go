package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/LindsayBradford/go-dbf/godbf"

	columnTypes "csv2dbf/column_types"
	"csv2dbf/file_loader"
	"csv2dbf/models"
)

func createColumns(table *godbf.DbfTable, columns []models.ColumnInfo) {
	for _, column := range columns {
		switch column.Type {
		case columnTypes.Number:
			err := table.AddNumberField(column.Name, column.Size, column.Precision)
			if err != nil {
				panic(err)
			}
		case columnTypes.String:
			err := table.AddTextField(column.Name, column.Size)
			if err != nil {
				panic(err)
			}
		default:
			panic(fmt.Sprintf("not supported column type: %s\n", column.Type))
		}
	}
}

func main() {
	pathToFilePtr := flag.String("in", "", "path to input csv")
	delimiterPtr := flag.String("delimiter", ",", "delimiter")
	pathToOutFilePtr := flag.String("out", "out.dbf", "path to out file")
	encodingPtr := flag.String("encoding", "win1251", "encoding (win1251, koi8r)")

	flag.Parse()

	// encodings
	// "UTF8", "Windows-1251", "koi8r"

	delimiter := *delimiterPtr
	pathToFile := *pathToFilePtr
	pathToOutFile := *pathToOutFilePtr
	encoding := "Windows-1251"
	if *encodingPtr == "koi8r" {
		encoding = *encodingPtr
	}

	fmt.Println("--- config ---")
	fmt.Println("delimiter:", *delimiterPtr)
	fmt.Println("encoding:", encoding)
	fmt.Println("path to input file:", pathToFile)
	fmt.Println("path to out file:", pathToOutFile)
	fmt.Println("")

	columns, dataLines := fileLoader.LoadFile(pathToFile, delimiter)

	table := godbf.New(encoding)
	createColumns(table, columns)

	for _, row := range dataLines {
		rowIdx, err := table.AddNewRecord()
		if err != nil {
			panic(err)
		}

		vals := strings.Split(row, delimiter)
		for colIdx, col := range columns {
			table.SetFieldValueByName(rowIdx, col.Name, vals[colIdx])
		}
	}

	err := godbf.SaveToFile(table, pathToOutFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}

// func test2() {
// 	table := godbf.New("UTF8")
//
// 	columns := getTestColumns()
// 	createColumns(table, columns)
//
// 	data := getTestData()
// 	for _, row := range data {
// 		rowIdx, err := table.AddNewRecord()
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		for colIdx, col := range columns {
// 			table.SetFieldValueByName(rowIdx, col.Name, row[colIdx])
// 		}
// 	}
//
// 	err := godbf.SaveToFile(table, "test.dbf")
// 	if err != nil {
// 		panic(err)
// 	}
// }

func test1() {
	table := godbf.New("UTF8")
	// table := godbf.New("Windows-1251")
	// table := godbf.New("koi8r")

	err := table.AddNumberField("id", 10, 0)
	if err != nil {
		panic(err)
	}

	err = table.AddTextField("name", 100)
	if err != nil {
		panic(err)
	}

	rowIdx, err := table.AddNewRecord()
	if err != nil {
		panic(err)
	}

	table.SetFieldValueByName(rowIdx, "id", "111")
	table.SetFieldValueByName(rowIdx, "name", "John")

	rowIdx, err = table.AddNewRecord()
	if err != nil {
		panic(err)
	}

	table.SetFieldValueByName(rowIdx, "id", "222")
	table.SetFieldValueByName(rowIdx, "name", "Велера")

	err = godbf.SaveToFile(table, "test.dbf")
	if err != nil {
		panic(err)
	}
}
