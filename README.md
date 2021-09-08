# bingWallpaper
golang语言编写，设置windows桌面壁纸，图片来自于[必应网站](https://cn.bing.com)  

使用[kardianos/service](https://github.com/kardianos/service/blob/master/example/runner/runner.json)
注册windows服务

每24小时自动下载并设置桌面壁纸

支持命令行操作,首次注册windows服务可以在cmd命令行（管理员身份打开）运行
```
bingWallpaper -install
```

目前遇到一个bug，不支持以后台服务的形式切换壁纸，暂时还不知道如何解决，原以为勾上"允许服务与桌面交互"会解决，然而并没有什么卵用，google了下好像要以隐藏的GUI程序切换壁纸才能生效。

目前仅支持运行bat文件，以后台exe的方式切换壁纸。

