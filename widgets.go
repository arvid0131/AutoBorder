package AutoBorder

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sqweek/dialog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func selectCommittee() widget.Select {
	return widget.Select{
		Options: []string{
			"Café",
			"AktU",
			"D-shopen",
			"D-chip",
			"Sexet - Alkoholfritt",
			"Sexet - Öl",
			"Sexet - Cider",
			"Sexet - Vin",
			"Sexet - Drink",
			"Sexet - Mat",
			"Sexet - Biljett",
			"UtEDischot",
		},
		OnChanged: func(s string) {
			parseCommittee(s)
			genCode()
		}}
}

func enterAmount() widget.Entry {
	return widget.Entry{
		PlaceHolder: "Amount",
		Validator: func(e string) error {
			if len(e) == 0 {
				return nil
			} else if _, err := strconv.ParseFloat(e, 32); err != nil {
				return errors.New("Invalid input, must be a number without leading zeroes!")
			}
			return nil
		},
		OnChanged: func(s string) {
			if _, err := strconv.ParseFloat(s, 32); err == nil {
				qr.amt, _ = strconv.ParseFloat(s, 64)
			}
			genCode()
		}}
}

func enterMessage() widget.Entry {
	return widget.Entry{
		PlaceHolder: "Message",
		Validator: func(e string) error {
			if strings.Contains(e, ";") {
				return errors.New("Message cannot contain semi-colon!")
			}
			return nil
		},
		OnChanged: func(s string) {
			if !strings.Contains(s, ";") {
				qr.itm = s
			}
			genCode()
		}}
}

func saveFileBtn() widget.Button {
	return *widget.NewButton("Save file...", func() {

		name, cancelled := dialog.File().Filter("png files", "png").Title("Save as png").Save()
		if cancelled == dialog.Cancelled {
			return
		} else {
			f, _ := os.Create(name + ".png")
			_, err := f.Write(generateQRBytes(generate_msg(qr)))
			if err != nil {
				log.Fatal(err)
			}
		}
	})
}

func generatedImage(s string) canvas.Image {
	return *canvas.NewImageFromImage(generateQrImg(s))
}

func genCode() {
	qrImg = generatedImage((generate_msg(qr)))
	qrImg.Resize(fyne.Size{
		Width:  256,
		Height: 256})
	qrImg.Move(fyne.NewPos(0, -128))
	canvas.Refresh(&qrImg)
}
