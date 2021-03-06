package main

import (
	"fmt"
	. "github.com/windosx/face-engine/v3"
	"github.com/windosx/face-engine/v3/util"
)

var width, height = util.GetImageWidthAndHeight("./test3.jpg")
var imageData = util.GetResizedBGR("./test3.jpg")

func main() {
	// 激活SDK
	if err := Activation("Your App ID", "Your SDK key"); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	// 初始化引擎
	engine, err := NewFaceEngine(DetectModeImage,
		OrientPriority0,
		12,
		50,
		EnableFaceDetect|EnableFaceRecognition|EnableFace3DAngle|EnableLiveness|EnableIRLiveness|EnableAge|EnableGender)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	// 检测人脸
	info, err := engine.DetectFaces(width-width%4, height, ColorFormatBGR24, imageData)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	// 处理人脸数据
	if err = engine.Process(width-width%4, height, ColorFormatBGR24, imageData, info, EnableAge|EnableGender|EnableFace3DAngle|EnableLiveness); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	// 获取年龄
	ageInfo, err := engine.GetAge()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("ageInfo: %v\n", ageInfo)
	// 销毁引擎
	if err = engine.Destroy(); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
