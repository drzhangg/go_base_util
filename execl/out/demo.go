package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	//f, err := excelize.OpenFile("Book1.xlsx")
	f, err := excelize.OpenFile("./execl/out/资产维修报表.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("资产维修报表", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("资产维修报表")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Println(colCell)
			//fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
