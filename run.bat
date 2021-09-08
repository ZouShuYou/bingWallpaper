@echo off 
if "%1" == "h" goto begin 
mshta vbscript:createobject("wscript.shell").run("%~nx0 h",0)(window.close)&&exit 
:begin 
path = %path%;.\;
bingWallpaper.exe -install
@echo "准备停止服务程序..."
sc stop BingWallpaper
@echo "设置允许与桌面进行交互方式允许"
sc config BingWallpaper type= interact type= own
@echo "正在重新启动服务..."
sc start BingWallpaper
@echo "启动服务成功！"
pause