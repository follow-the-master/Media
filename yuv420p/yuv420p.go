package main

import (
	"fmt"
	"os"
)

// Splityuv420p is split yuv data of a yuv file
// save yuv data to 3 file
// todo:
// refine: 重构减少代码量
// refine: 支持读多帧
// refine: 丰富错误处理(eg:num范围异常/yuv EOF支持)
func Splityuv420p(yuvfile string, w, h, num int) (err error) {
	yuv, err := os.Open(yuvfile)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}

	defer yuv.Close()

	y, err := os.OpenFile(yuvfile + ".y.yuv", os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("create y file error:", err)
		return
	}

	defer y.Close()

	u, err := os.OpenFile(yuvfile + ".u.yuv", os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("create u file error:", err)
		return
	}

	defer u.Close()

	v, err := os.OpenFile(yuvfile + ".v.yuv", os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("create v file error:", err)
		return
	}

	defer v.Close()

	yBuf := make([]byte, w*h, w*h)
	uBuf := make([]byte, w*h/4, w*h/4)
	vBuf := make([]byte, w*h/4, w*h/4)

	if _, err = yuv.Read(yBuf); err != nil {
		fmt.Println("read yuv file error:", err)
		return
	}

	if _, err = y.Write(yBuf); err != nil {
		fmt.Println("write y file error:", err)
		return
	}

		if _, err = yuv.Read(uBuf); err != nil {
		fmt.Println("read yuv file error:", err)
		return
	}

	if _, err = u.Write(uBuf); err != nil {
		fmt.Println("write u file error:", err)
		return
	}

	if _, err = yuv.Read(vBuf); err != nil {
		fmt.Println("read yuv file error:", err)
		return
	}

	if _, err = v.Write(vBuf); err != nil {
		fmt.Println("write v file error:", err)
		return
	}

	return nil
}

// SplitOneFrameyuv420p is split 1 frame of Split
func SplitOneFrameyuv420p(yuvfile string, w,h int)error{
	return Splityuv420p(yuvfile, w,h, 1)
}

func main(){
	SplitOneFrameyuv420p("./a.yuv", 1920, 1200)
}