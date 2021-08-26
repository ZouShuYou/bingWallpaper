package bing

import (
	"bingWallpaper/util"
	"bingWallpaper/wallpaper"
	"github.com/kardianos/service"
	"log"
	"time"
)

var logger service.Logger

type Program struct{
	exit chan struct{}
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
	ticker := time.NewTicker(24 * time.Hour)
	for  {
		select {
		case tm := <- ticker.C:
			log.Printf("Bing Wallpaper Service is still running at %v",tm)
			imagePath, _ := wallpaper.FetchAndWrite()
			log.Printf("get wallpaper and save in %s",imagePath)
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
