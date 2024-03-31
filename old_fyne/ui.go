package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GetBrowseModsCard(data MMModData) fyne.Widget {
	bDownload := widget.NewButtonWithIcon("Download", theme.ContentAddIcon(), func() {
		return
	})
	var iThumb *canvas.Image
	iThumb = canvas.NewImageFromFile("placeholder-grey.png")
	card := widget.NewCard("", "", canvas.NewText(string(data.FileSize)+" bytes", color.White))
	card.SetImage(iThumb)
	go func() {
		thumbUri, err := storage.ParseURI(data.ThumbUrl)
		if err != nil {
			logger.Printf("Error parsing mod thumbnail URL for %s: %+v", data.ThumbUrl, err)
			return
		}
		iThumb = canvas.NewImageFromURI(thumbUri)
		card.SetImage(iThumb)
		card.Image.FillMode = canvas.ImageFillContain
	}()

	lbName := widget.NewLabel(data.Name)
	lbAuthor := widget.NewLabel(data.Author)
	lbName.Wrapping = fyne.TextWrapWord
	lbAuthor.Wrapping = fyne.TextWrapWord

	loVert := container.NewVBox(
		lbName,
		lbAuthor,
		widget.NewLabel(fmt.Sprintf("Size (kb): %d", data.FileSize/1024)),
		bDownload,
	)

	card.SetContent(loVert)
	return card
}
func getInstalledModsCard(data MMModData) fyne.Widget {
	return nil
}

func getInstalledTabUi() fyne.CanvasObject {
	return widget.NewLabel("Retrieving information from the API...")
}

func getBrowseTabUi() fyne.CanvasObject {
	g := container.NewGridWrap(fyne.NewSize(240, 300))
	s := container.NewScroll(g)

	cont := container.NewMax(s)

	go func() {
		mods, _, _ := GetModsFromRemote(0)
		g.RemoveAll()
		for _, mod := range mods {
			g.Add(GetBrowseModsCard(mod))

		}

		s.Refresh()
		cont.Refresh()

	}()

	return cont
}
