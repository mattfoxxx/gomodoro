package gomodoro

import (
	"fmt"
	"fyne.io/fyne/v2"
	app "fyne.io/fyne/v2/app"
	_ "fyne.io/fyne/v2/container"
	_ "fyne.io/fyne/v2/layout"
	_ "fyne.io/fyne/v2/widget"
	"fyne.io/systray"
	"fyne.io/systray/example/icon"
	"time"
)

var gomodoroApp fyne.App
var mainWindow fyne.Window
var splitsWindow fyne.Window

// Launch starts the pomodoro system tray app
func Launch() {
	gomodoroApp = app.NewWithID("com.github.mattfoxxx.gomodoro")
	systrayStart, systrayStop := systray.RunWithExternalLoop(onReady, onExit)

	mainWindow = gomodoroApp.NewWindow("Main Application Window")
	mainWindow.SetMaster()
	mainWindow.Resize(fyne.NewSize(1920, 1024))
	mainWindow.CenterOnScreen()
	mainWindow.SetCloseIntercept(func() {
		systrayStop()
		gomodoroApp.Quit()
	})
	mainWindow.Hide()

	splitsWindow = gomodoroApp.NewWindow("Hidden Pomodoro Splits Window")
	splitsWindow.Resize(fyne.NewSize(640, 480))
	splitsWindow.CenterOnScreen()
	splitsWindow.SetCloseIntercept(func() {
		splitsWindow.Hide()
	})
	splitsWindow.Hide()

	systrayStart()
	gomodoroApp.Run()
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Gomodoro")
	systray.SetTooltip("gomodoro - your pomodoro timer")
	mNotify := systray.AddMenuItem("Notify", "Send a notification")
	mSplits := systray.AddMenuItem("Intervals", "Configure intervals")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Sets the icon of gomodoro_app menu item. Only available on Mac and Windows.
	mQuit.SetIcon(icon.Data)

	/*	go func() {
		for {
			systray.SetTitle(getClockTime(timezone))
			systray.SetTooltip(timezone + " timezone")
			time.Sleep(1 * time.Second)
		}
	}()*/
	go func() {
		for {
			select {
			case <-mSplits.ClickedCh:
				fmt.Println("Intervals was clicked!")
				go showWindow(splitsWindow)
			case <-mNotify.ClickedCh:
				fmt.Println("Notify was clicked!")
				go showNotification(gomodoroApp)
			case <-mQuit.ClickedCh:
				fmt.Println("User requested to quit the application!")
				gomodoroApp.Quit()
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// clean up here
	gomodoroApp.Quit()
}

func showNotification(a fyne.App) {
	time.Sleep(time.Second * 2)
	a.SendNotification(fyne.NewNotification("Example Title", "Example Content"))
}

func showWindow(w fyne.Window) {
	fmt.Println("In showWindow function")
	w.Show()
}
