package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

// 写文件，以单元格的值进行写入

//func main() {
//	f := excelize.NewFile()
//	defer func() {
//		if err := f.Close(); err != nil {
//			fmt.Println(err)
//		}
//	}()
//	// Create a new sheet.
//	index, err := f.NewSheet("Sheet2")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	// Set value of a cell.
//	f.SetCellValue("Sheet2", "A2", "Hello world.")
//	f.SetCellValue("Sheet1", "B2", 100)
//	// Set active sheet of the workbook.
//	f.SetActiveSheet(index)
//	// Save spreadsheet by the given path.
//	if err := f.SaveAs("Book1.xlsx"); err != nil {
//		fmt.Println(err)
//	}
//}

// 读文件，可以按值，也可以读全部

func main() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and cell reference.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
