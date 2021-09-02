# bingWallpaper
golang语言编写，设置windows桌面壁纸，图片来自于[必应网站](https://cn.bing.com)  

使用[kardianos/service](https://github.com/kardianos/service/blob/master/example/runner/runner.json)
注册windows服务

每24小时自动下载并设置桌面壁纸

支持命令行操作,首次注册windows服务需在bingWallpaper目录下，使用cmd命令行（管理员身份打开）运行
```
bingWallpaper -install
```

壁纸文件保存路径在url.go文件中，可根据自己需求更改

