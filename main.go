package main

import (
	"fmt"

	"github.com/signintech/gopdf"
)

func main() {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("arial", "./fonts/Noto_Sans_JP/NotoSansJP-VariableFont_wght.ttf")

	if err != nil {
		fmt.Println("Error loading font:", err)
		return
	}
	err = pdf.SetFont("arial", "", 14)
	if err != nil {
		fmt.Println("Error setting font:", err)
		return
	}

	pdf.Cell(nil, "すべての")
	err = pdf.Image("./pdf.jpeg", 10, 50, nil) // Replace "image.jpg" with the actual image file path
	if err != nil {
		fmt.Println("Error adding image:", err)
		return
	}
	pdf.WritePdf("hello.pdf")

}
