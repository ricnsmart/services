package services

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

/*
@param titles map{{"序列号":"A"},{"设备类型":"B"}}
*/
func NewXLSX(filePath string, data []map[string]interface{}, titles map[string]string) error {
	f := excelize.NewFile()
	sheet := "Sheet1"
	// Create a new sheet.
	index := f.NewSheet(sheet)

	for title, column := range titles {
		axis := fmt.Sprintf(`%v1`, column)
		f.SetCellValue(sheet, axis, title)
	}

	for rowNo, rowData := range data {
		for key, value := range rowData {
			axis := fmt.Sprintf(`%v%v`, titles[key], rowNo+2)
			f.SetCellValue(sheet, axis, value)
		}
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	if err := f.SaveAs(filePath); err != nil {
		return err
	}

	return nil
}
