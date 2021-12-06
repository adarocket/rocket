package main

import (
	"adarocket/rocket/client"
	"google.golang.org/grpc"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var redColor = color.RGBA{R: 255, G: 0, B: 0, A: 0}
var greenColor = color.RGBA{R: 0, G: 255, B: 0, A: 0}
var orangeColor = color.RGBA{R: 255, G: 165, B: 0, A: 0}
var grayColor = color.RGBA{R: 220, G: 220, B: 220, A: 0}

const preferenceCurrentTutorial = "currentTutorial"

var topWindow fyne.Window

func main() {
	clientConn, err := grpc.Dial("178.124.167.214:5300", grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	authClient = client.NewAuthClient(clientConn)

	a := app.New()
	a.SetIcon(theme.FyneLogo())
	// a.Settings().SetTheme(theme.DarkTheme())

	w := a.NewWindow("Ada Rocket")
	topWindow = w

	w.SetMaster()

	w.SetContent(loginScreen(w, a))
	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}

func makeNav(setTutorial func(menu *MenuField), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return MenuIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := MenuIndex[uid]
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return canvas.NewText("Collection Widgets", greenColor)
			// return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			m, ok := MenuFields[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*canvas.Text).Text = m.Title

			switch m.Status {
			case "OK":
				obj.(*canvas.Text).Color = greenColor
			case "WARNING":
				obj.(*canvas.Text).Color = orangeColor
			case "ERROR":
				obj.(*canvas.Text).Color = redColor
			default:
				obj.(*canvas.Text).Color = grayColor
			}

			// obj.(*widget.Label).SetText(m.Title)
		},
		OnSelected: func(uid string) {
			if m, ok := MenuFields[uid]; ok {
				a.Preferences().SetString(preferenceCurrentTutorial, uid)
				setTutorial(m)
			}
		},
	}

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
		tree.Select(currentPref)
	}

	a.Settings().SetTheme(theme.DarkTheme())

	// FIXME: на светлой теме весь текст белый
	// themes := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
	// 	widget.NewButton("Dark", func() {
	// 		a.Settings().SetTheme(theme.DarkTheme())
	// 	}),
	// 	widget.NewButton("Light", func() {
	// 		a.Settings().SetTheme(theme.LightTheme())
	// 	}),
	// )
	// return container.NewBorder(nil, themes, nil, nil, tree)

	return container.NewBorder(nil, nil, nil, nil, tree)
}
