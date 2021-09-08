package util

import (
	"errors"
	"os"
	"path"
	"syscall"
	"unsafe"
	"bingWallpaper/constant"
)

func MakeDir(data string) (string,error) {
	dir := path.Join(constant.BingDir,data)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}
	return dir,nil
}

// SetWindowsWallpaper 设置windows壁纸
func SetWindowsWallpaper(imagePath string) error {
	dll, err := syscall.LoadDLL("user32.dll")
	if err != nil {
		Error.Println(err.Error())
	}
	proc, err := dll.FindProc("SystemParametersInfoW")
	if err != nil {
		Error.Println(err.Error())
	}
	//_t, _ := syscall.UTF16PtrFromString(imagePath)
	//ret, _, _ := proc.Call(IntPtr(20), IntPtr(1), StrPtr(imagePath), IntPtr(3))
	_t, _ := syscall.UTF16PtrFromString(imagePath)
	ret, _, err := proc.Call(20, 1, uintptr(unsafe.Pointer(_t)), 3)
	if err != nil {
		Error.Println(err.Error())
	}
	if ret != 1 {
		return errors.New("系统调用失败")
	}
	return nil
}

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func StrPtr(s string) uintptr {
	p, _ := syscall.UTF16PtrFromString(s)
	return uintptr(unsafe.Pointer(p))
}