package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sqweek/dialog"
)

type LaunchFlags struct {
	DataPath string
	GamePath string
	ModDir   string
	ApiKey   string
	UseJson  bool
}

var MMApplication fyne.App
var Flags LaunchFlags
var logger = log.Default()

func init() {
}

func main() {
	Flags = parseFlags()
	MMApplication = app.New()

	LoadMMData(Flags.DataPath)

	go startUI()

	MMApplication.Run()
}

func startUI() {

	wMain := MMApplication.NewWindow("MainWindow")
	wMain.SetContent(container.NewVBox(
		widget.NewLabel("Waiting for API Key..."),
	),
	)

	res, _ := fyne.LoadResourceFromPath("logo.png")

	wMain.SetIcon(res)

	key_chan := make(chan string)
	wMain.SetMaster()

	wMain.Resize(fyne.Size{
		Width:  600,
		Height: 400,
	})
	wMain.CenterOnScreen()
	wMain.SetFixedSize(true)
	wMain.SetTitle("CC Mod Manager")
	wMain.Show()

	wMain.Hide()
	if Data.ApiKey == "" {
		logger.Printf("First run!")
		AskForApiKey(key_chan)
		key := <-key_chan
		logger.Printf("Grabbed key! Saving...")
		GetMMData().ApiKey = key
		SaveMMData(Flags.DataPath)
	}

	if Data.GamePath == "" || Data.ModDir == "" {
		AskForModDir(key_chan)
		dir := <-key_chan
		logger.Printf("Set CC mod directory to %s", dir)
		GetMMData().ModDir = dir
		SaveMMData(Flags.DataPath)
	}

	populateMainWindow(wMain)

	wMain.SetMainMenu(getMainMenu())

	wMain.Show()

}

func AskForModDir(dir chan string) {
	w := MMApplication.NewWindow("Mod Directory")

	l1 := widget.NewLabel("Please specify your Cortex Command Mods Folder")
	l2 := widget.NewLabel("Path: <None>")

	path := ""
	bBrowse := widget.NewButton("Browse", func() {
		path, _ = dialog.Directory().Title("Cortex Command Mod Folder").Browse()
		l2.SetText(fmt.Sprintf("Path: %s", path))
	})

	bSubmit := widget.NewButton("Submit", func() {
		w.Close()
	})

	cont := container.NewVBox(
		l1, l2, bBrowse, bSubmit,
	)

	w.SetContent(cont)
	w.Show()

	w.SetOnClosed(func() {
		dir <- path
	})

}

func getMainMenu() *fyne.MainMenu {
	mSettings := fyne.NewMenu(
		"File",
		fyne.NewMenuItem(
			"Settings", openSettings),
	)
	return fyne.NewMainMenu(mSettings)
}

func getInstalledTabUi() fyne.CanvasObject {
	return widget.NewLabel("TODO")
}

func getBrowseTabUi() fyne.CanvasObject {
	return widget.NewLabel("TODO")
}

func populateMainWindow(win fyne.Window) {

	tabs := container.NewAppTabs(
		container.NewTabItem("Installed mods", getInstalledTabUi()),
		container.NewTabItem("Browse mods", getBrowseTabUi()),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	win.SetContent(tabs)

}

func AskForApiKey(done chan string) {

	APIKeyPagePath, _ := url.Parse("https://mod.io/me/access")

	w := MMApplication.NewWindow("API Key")
	l1 := widget.NewLabel("In order to use this application you need a mod.io API key")

	l2 := widget.NewHyperlink("Mod.io API Key page", APIKeyPagePath)
	in := widget.NewEntry()
	in.SetPlaceHolder("API Key")
	key := ""
	sub := widget.NewButton("Submit", func() {
		key = strings.TrimSpace(in.Text)
		w.Close()
	})

	cont := container.NewVBox(
		l1, l2, in, sub,
	)

	w.SetContent(cont)
	w.Show()

	w.SetOnClosed(func() {
		done <- key
	})
}

func openSettings() {

}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func parseFlags() LaunchFlags {
	f := LaunchFlags{
		DataPath: *flag.String("d", "mm.dat", "override the default data file path"),
		GamePath: *flag.String("g", "", "override the set game directory path"),
		ModDir:   *flag.String("m", "", "override the mod dir path (will use gamepth mod dir if unspecified)"),
		ApiKey:   *flag.String("a", "", "override the api key"),
		UseJson:  *flag.Bool("json", false, "uses json instead of gob"),
	}

	flag.Parse()

	return f

}
