package main

import (
	"cyj/proxy"
	"embed"
	_ "github.com/icza/mjpeg"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	//_ "net/http/pprof"
)

const lockUniqueId = "C04B4B88-2B22-3ADE-9DFE-551EBA430D19"

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
		OnStartup: app.startup,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               lockUniqueId,               // 单实例锁
			OnSecondInstanceLaunch: app.onSecondInstanceLaunch, // 如果重复打开就再次显示窗口
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
			DisableWindowIcon:                 false,
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

func (a *App) onSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
	runtime.WindowUnminimise(a.ctx)
	runtime.Show(a.ctx)
}
