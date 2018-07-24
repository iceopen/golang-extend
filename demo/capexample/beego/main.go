package main

import (
	"github.com/mojocn/base64Captcha"
	"fmt"
	"encoding/json"
)

type ConfigJsonBody struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

func main() {
	var postParameters ConfigJsonBody
	str:=`{"CaptchaType":"character","Id":"","VerifyValue":"","ConfigAudio":{"CaptchaLen":6,"Language":"zh"},"ConfigCharacter":{"Height":60,"Width":240,"Mode":2,"ComplexOfNoiseText":0,"ComplexOfNoiseDot":0,"IsUseSimpleFont":true,"IsShowHollowLine":false,"IsShowNoiseDot":false,"IsShowNoiseText":false,"IsShowSlimeLine":false,"IsShowSineLine":false,"CaptchaLen":6},"ConfigDigit":{"Height":80,"Width":240,"CaptchaLen":5,"MaxSkew":0.7,"DotCount":80}}`
	json.Unmarshal([]byte(str), &postParameters)
	fmt.Println(postParameters)
	var config interface{}
	config = postParameters.ConfigCharacter

	captchaId, captcaInterfaceInstance := base64Captcha.GenerateCaptcha(postParameters.Id, config)
	base64blob := base64Captcha.CaptchaWriteToBase64Encoding(captcaInterfaceInstance)
	fmt.Println(captchaId)
	fmt.Println(base64blob)
}
