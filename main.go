package main

import (
	"bingWallpaper/bing"
	"bingWallpaper/constant"
	"bingWallpaper/util"
	"flag"
	"github.com/kardianos/service"
	"runtime"
)

var (
	version 	bool
	help 		bool
	run 		bool
	stop		bool
	install		bool
	remove		bool
	start		bool
	restart		bool
)

func init()  {
	flag.BoolVar(&version,"v",false,"show current version of bingWallpaper and exit")
	flag.BoolVar(&help,"h",false,"this help")
	flag.BoolVar(&run,"run",false,"start bingWallpaper windows service")
	flag.BoolVar(&stop,"stop",false,"stop bingWallpaper windows service")
	flag.BoolVar(&install,"install",false,"restart bingWallpaper windows service")
	flag.BoolVar(&remove,"remove",false,"remove bingWallpaper windows service")
	flag.BoolVar(&start,"start",false,"start bingWallpaper windows service")
	flag.BoolVar(&restart,"restart",false,"restart bingWallpaper windows service")

	flag.Parse()
}

func main() {
	arguments := []string{"-run"}
	option := make(service.KeyValue)
	option["Interactive"] = true

	serviceConfig := &service.Config{
		Name:        	"BingWallpaper",
		DisplayName: 	"Bing Wallpaper Service",
		Description: 	"每日同步微软bing壁纸",
		Executable: 	constant.Executable,
		Arguments:		arguments,
		Option: 		option,
	}

	prg := &bing.Program{}
	s, err := service.New(prg, serviceConfig)
	prg.Service = s

	if err != nil {
		util.Info.Fatal(err)
	}

	if version {
		util.Info.Printf("bingWallpaper %s on %s %s with %s %s\n", constant.Version, runtime.GOOS, runtime.GOARCH, runtime.Version(), constant.BuildTime)
		return
	}

	if help {
		flag.PrintDefaults()
	}

	if run {
		s.Run()
	}

	if stop {
		s.Stop()
	}

	if install {
		s.Install()
		s.Start()
	}

	if remove {
		s.Stop()
		s.Uninstall()
	}

	if start {
		s.Start()
	}

	if restart {
		s.Restart()
	}
}