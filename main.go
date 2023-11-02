package main

import (
	"embed"
	_ "github.com/icza/mjpeg"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"log"
	_ "net/http/pprof"
	"qh-tool/proxy"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {

	// Create an instance of the app structure
	// 创建一个App结构体实例
	app := NewApp()
	// Create application with options
	// 使用选项创建应用
	err := wails.Run(&options.App{
		Title:             "",
		Width:             480,
		Height:            700,
		MinWidth:          480,
		MinHeight:         700,
		DisableResize:     true,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		Menu:              nil,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnBeforeClose:     app.beforeClose,
		OnShutdown:        app.shutdown,
		WindowStartState:  options.Normal,
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Handler:    nil,
			Middleware: nil,
		},
		Bind: []interface{}{
			app,
			proxy.NewProxy(),
		},
		// Windows platform specific options
		// Windows平台特定选项
		Windows: &windows.Options{
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:           0xffffff,
				DarkModeTitleBarInactive:   0xffffff,
				DarkModeTitleText:          0xffffff,
				DarkModeTitleTextInactive:  0xffffff,
				DarkModeBorder:             0xffffff,
				DarkModeBorderInactive:     0xffffff,
				LightModeTitleBar:          0xffffff,
				LightModeTitleBarInactive:  0xffffff,
				LightModeTitleText:         0xffffff,
				LightModeTitleTextInactive: 0xffffff,
				LightModeBorder:            0xffffff,
				LightModeBorderInactive:    0xffffff,
			},
			DisableWindowIcon:                 true,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.Light,
		},
		// Mac platform specific options
		// Mac平台特定选项
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            true,
				UseToolbar:                 false,
				HideToolbarSeparator:       false,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "穿云箭",
				Message: "穿云箭-内网究极神器",
				Icon:    icon,
			},
		},
		Linux: &linux.Options{
			Icon: icon,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
