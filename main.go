package AutoBorder

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var qr = QR_String{bitmask: 0}

var comitteeSelector widget.Select
var amountEntry widget.Entry
var messageEntry widget.Entry
var amountLock widget.Check
var generateButton widget.Button
var qrImg canvas.Image
var saveFile dialog.FileDialog

var a = app.New()
var w = a.NewWindow("Auto-Border")
var c = w.Canvas()
var saveFileWindow = a.NewWindow("Save file")
var saveFileCanvas = saveFileWindow.Canvas()

func main() {

	comitteeSelector = selectCommittee()
	amountEntry = enterAmount()
	messageEntry = enterMessage()
	generateButton = saveFileBtn()

	c.SetContent(widget.NewLabel("Hello World!"))
	w.Resize(fyne.NewSize(400, 500))
	saveFileWindow.Resize(fyne.NewSize(1000, 700))

	canvas.Refresh(&qrImg)

	content := container.NewCenter(
		container.NewGridWithRows(
			2,
			container.NewStack(&qrImg),
			container.NewVBox(
				&amountEntry,
				container.NewGridWithColumns(2, &comitteeSelector, &messageEntry),
				&generateButton)))
	c.SetContent(content)

	w.ShowAndRun()

	tidyUp()

}

func tidyUp() {
	fmt.Println("Exited")
}
