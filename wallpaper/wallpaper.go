package wallpaper

import (
	"bingWallpaper/util"
	"bingWallpaper/constant"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

const (

)

type Wallpaper struct {
	Copyright  string `json:"copyright"`
	Url    string `json:"url"`
	Date   string `json:"enddate"`
}

type BingWallpaperResult struct {
	Images	[]*Wallpaper
}

//获取图片json信息
func GetImageInfo() (*BingWallpaperResult, error){
	resp, err := http.Get(constant.ImgUrl)
	if err != nil {
		return nil,err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("get from bing wallpaper failed: %s", resp.Status)
	}

	var result BingWallpaperResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

//获取并写入图片
func FetchAndWrite() (string,error){
	wallpaperResult, _ := GetImageInfo()
	image := wallpaperResult.Images[0]
	url := constant.BingUrl + image.Url
	resp, err := http.Get(url)
	if err != nil {
		return "",err
	}
	if resp.StatusCode != http.StatusOK{
		resp.Body.Close()
		return "",fmt.Errorf("get image from bing failed: %s", resp.Status)
	}
	dir, err := util.MakeDir(image.Date)
	if err != nil {
		return "",err
	}
	index := strings.Index(image.Copyright,"(")
	fileName := image.Copyright[:index]+".jpg"
	filePath := path.Join(dir,fileName)
	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	err = ioutil.WriteFile(filePath, imageData, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return "",err
	}
	return filePath,err
}