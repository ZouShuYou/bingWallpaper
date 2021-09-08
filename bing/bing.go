package bing

import (
	"bingWallpaper/constant"
	"bingWallpaper/util"
	"bingWallpaper/wallpaper"
	"github.com/kardianos/service"
	"log"
	"time"
)


type Program struct{
	exit    chan struct{}
	Service service.Service
}

func (p *Program) Start(s service.Service) error  {
	if service.Interactive() {
		log.Println("Bing Wallpaper is Running in terminal.")
	}else {
		log.Println("Bing Wallpaper is Running under service manager.")
	}
	p.exit = make(chan struct{})
	go p.run()
	return nil
}

func (p *Program) run() {
	util.Info.Println("Bing Wallpaper Service is running......")
	imagePath, _ := wallpaper.FetchAndWrite()
	util.Info.Printf("get wallpaper and save in %s",imagePath)
	err := util.SetWindowsWallpaper(imagePath)
	if err != nil {
		util.Error.Println(err.Error())
	}
	ticker := time.NewTicker(constant.Time)
	for  {
		select {
		case tm := <- ticker.C:
			util.Info.Printf("Bing Wallpaper Service is still running at %v",tm)
			imagePath, _ := wallpaper.FetchAndWrite()
			util.Info.Printf("get wallpaper and save in %s",imagePath)
			util.SetWindowsWallpaper(imagePath)
		case <-p.exit:
			ticker.Stop()
		}
	}
}

func (p *Program) Stop(s service.Service) error  {
	log.Println("Bing Wallpaper is Stopping.")
	close(p.exit)
	return nil
}
