package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
)

func main() {

	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("mset brouser")
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.NewVBox(true, 0)
	vbox.SetBorderWidth(5)
	drawingArea := gtk.NewDrawingArea()

	var pixmap *gdk.Pixmap
	var gc *gdk.GC

	drawingArea.Connect("configure-event", func() {

		if pixmap != nil {
			pixmap.Unref()
		}
		allocation := drawingArea.GetAllocation()
		pixmap = gdk.NewPixmap(drawingArea.GetWindow().GetDrawable(), allocation.Width, allocation.Height, 24)

		fmt.Printf("[%d %d]\n", allocation.Width, allocation.Height)

		gc = gdk.NewGC(pixmap.GetDrawable())
		gc.SetRgbFgColor(gdk.NewColor("white"))
		pixmap.GetDrawable().DrawRectangle(gc, true, 0, 0, -1, -1)
		gc.SetRgbFgColor(gdk.NewColor("black"))
		gc.SetRgbBgColor(gdk.NewColor("white"))
	})

	drawingArea.Connect("expose-event", func() {
		if pixmap != nil {
			d := pixmap.GetDrawable()

			gc.SetRgbFgColor(gdk.NewColorRGB(255, 0, 0))
			for y := 0; y < 100; y++ {
				for x := 0; x < 200; x++ {
					d.DrawPoint(gc, x, y)
				}
			}

			gc.SetRgbFgColor(gdk.NewColorRGB(0, 127, 0))
			d.DrawLine(gc, 0, 0, 200, 200)

			drawingArea.GetWindow().GetDrawable().DrawDrawable(gc, pixmap.GetDrawable(), 0, 0, 0, 0, -1, -1)
		}
	})

	drawingArea.SetEvents(int(gdk.BUTTON_PRESS_MASK))

	vbox.Add(drawingArea)
	window.Add(vbox)

	window.SetSizeRequest(512, 512)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.ShowAll()

	gtk.Main()
}
