package main

import (
	"bingWallpaper/bing"
	"bingWallpaper/constant"
	"flag"
	"github.com/kardianos/service"
	"log"
	"runtime"
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
		DisplayName: "bing Wallpaper service",
		Description: "每日同步微软bing壁纸",
	}

	prg := &bing.Program{}

	s, err := service.New(prg, serviceConfig)
	if err != nil {
		log.Fatal(err)
	}

	if version {
		log.Printf("bingWallpaper %s on %s %s with %s %s\n", constant.Version, runtime.GOOS, runtime.GOARCH, runtime.Version(), constant.BuildTime)
		return
	}

	if help {
		flag.PrintDefaults()
	}

	if start {
		s.Start()
	}

	if stop {
		s.Stop()
	}

}