package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("github.com/xxxserxxx/fyneissue2")

	w := a.NewWindow("Fyne Issue 2")

	items := make([]string, 20)
	for i := 0; i < 20; i++ {
		items[i] = fmt.Sprintf("Item #%d", i)
	}

	itemList := widget.NewList(
		func() int {
			return len(items)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(it widget.ListItemID, o fyne.CanvasObject) {
			c, _ := o.(*widget.Label)
			c.Text = items[it]
			c.Refresh()
		},
	)
	addItem := widget.NewEntry()
	/*
		lambda := func(val string) {
			items = append([]string{val}, items...)
			itemList.Refresh()
		}
		addItem.OnSubmitted = lambda
		addButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() { lambda(addItem.Text) })
	*/
	addButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {})
	addRow := container.New(layout.NewBorderLayout(nil, nil, nil, addButton), addItem, addButton)

	c := container.New(layout.NewBorderLayout(addRow, nil, nil, nil), addRow, itemList)
	itemList.Refresh()
	w.SetContent(c)
	w.ShowAndRun()
}
