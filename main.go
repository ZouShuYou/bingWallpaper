package  main

import (
	"bingWallpaper/util"
	"bingWallpaper/wallpaper"
)


func main() {
	imagePath, _ := wallpaper.FetchAndWrite()
	util.SetWindowsWallpaper(imagePath)
}