package tests

import (
	"github.com/jwijenbergh/puregotk/v4/adw"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

func baseTestApp(setupWebView func() *gtk.Widget) {
	adw.Init()

	app := adw.NewApplication("app.example.test", gio.GApplicationFlagsNoneValue)
	defer app.Unref()

	actcb := func(_ gio.Application) {
		window := adw.NewApplicationWindow(&app.Application)
		mainBox := gtk.NewBox(gtk.OrientationVerticalValue, 0)

		headerBar := adw.NewHeaderBar()
		headerBar.AddCssClass("flat")
		headerBar.AddCssClass("compact")
		headerBar.SetShowTitle(false)
		headerBar.SetShowEndTitleButtons(true)
		mainBox.Append(&headerBar.Widget)

		webViewWidget := setupWebView()
		webViewWidget.SetVexpand(true)
		webViewWidget.SetHexpand(true)
		webViewWidget.Show()

		mainBox.Append(webViewWidget)
		window.SetContent(&mainBox.Widget)
		window.SetDefaultSize(800, 600)
		window.Present()
	}
	app.ConnectActivate(&actcb)

	app.Run(1, []string{"app.example.test"})
}
