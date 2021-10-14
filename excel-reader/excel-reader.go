package excel_reader

import (
	"fmt"
	"github.com/xuri/excelize"
	"log"
)

var path = "/Users/IansIpad/Projects/goworkspace/src/go-playground/go-playground/text-files/"

func ExcelReader() {
	fmt.Println("Things are off to a good start...")
	f, err := excelize.OpenFile(path + "creditcard.xlsx")

	if err != nil {
		fmt.Println("OMG Error!!!!")
	}
	error1 := f.Close()
	if error1 != nil {
		fmt.Println("OMG Error!!")
	}

	c1, error2 := f.GetCellValue("Sheet1", "A1")

	if error2 != nil {
		log.Fatal(error2)
	}

	fmt.Println(c1)
	fmt.Println("We made it through the function...")
}
