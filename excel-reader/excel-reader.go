package excel_reader

import (
	"fmt"
	"github.com/xuri/excelize"
)

var path = "/Users/IansIpad/Projects/goworkspace/src/go-playground/go-playground/text-files/"
var sheet = "Sheet1"

func ExcelReader() {
	fmt.Println("Things are off to a good start...")
	f, err := excelize.OpenFile(path + "creditcard.xlsx")

	if err != nil {
		fmt.Println("OMG Error!!!!")
	}

	err = f.Close()
	if err != nil {
		fmt.Println("OMG Error!!")
	}

	c1, err := f.GetCellValue(sheet, "A1")
	if err != nil {
		fmt.Println("OMG Error!!")
	}

	rows, err := f.GetRows(sheet)
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

	fmt.Println(c1)
	fmt.Println("We made it through the function...")
}
