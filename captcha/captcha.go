package captcha

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = base64Captcha.DefaultMemStore

func DriverAudioFunc() (id, b64s string, err error) {
	driver := base64Captcha.DefaultDriverAudio
	cap := base64Captcha.NewCaptcha(driver, store)
	return cap.Generate()
}

func DriverStringFunc() (id, b64s string, err error) {
	height := 46
	width := 140
	noiseCount := 2
	showLineOptions := 2
	length := 4
	source := "234567890abcdefghjkmnpqrstuvwxyz"
	bgColor := &color.RGBA{240, 240, 246, 246}
	var fontsStorage base64Captcha.FontsStorage = base64Captcha.DefaultEmbeddedFonts
	var fonts []string = []string{"wqy-microhei.ttc"}
	driver := base64Captcha.NewDriverString(height, width, noiseCount, showLineOptions, length, source, bgColor, fontsStorage, fonts)
	cap := base64Captcha.NewCaptcha(driver.ConvertFonts(), store)
	return cap.Generate()
}

func DriverDigitFunc() (id, b64s string, err error) {
	driver := base64Captcha.DefaultDriverDigit
	cap := base64Captcha.NewCaptcha(driver, store)
	return cap.Generate()
}

// Verify captcha's answer directly
func Verify(id, answer string, clear bool) bool {
	return store.Verify(id, answer, clear)
}
