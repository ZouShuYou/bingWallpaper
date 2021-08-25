package main

import (
	"bingWallpaper/constant"
	"bingWallpaper/util"
	"bingWallpaper/wallpaper"
	"flag"
	"fmt"
	"runtime"
	"github.com/kardianos/service"
)

var (
	version 	bool
	help 		bool
	start 		bool
	stop		bool
	restart		bool
	remove		bool
)

func init()  {
	flag.BoolVar(&version,"v",false,"show current version of bingWallpaper and exit")
	flag.BoolVar(&help,"h",false,"this help")
	flag.BoolVar(&start,"start",false,"start bingWallpaper windows service")
	flag.BoolVar(&stop,"stop",false,"stop bingWallpaper windows service")
	flag.BoolVar(&restart,"restart",false,"restart bingWallpaper windows service")
	flag.BoolVar(&remove,"remove",false,"remove bingWallpaper windows service")

	flag.Parse()
}

func main() {
	serviceConfig := &service.Config{
		Name:        "bingWallpaper",
		DisplayName: "bingWallpaper service",
		Description: "每日同步微软bing壁纸",
	}

	if version {
		fmt.Printf("bingWallpaper %s on %s %s with %s %s\n", constant.Version, runtime.GOOS, runtime.GOARCH, runtime.Version(), constant.BuildTime)
		return
	}

	if help {
		flag.PrintDefaults()
	}

	imagePath, _ := wallpaper.FetchAndWrite()
	util.SetWindowsWallpaper(imagePath)
}