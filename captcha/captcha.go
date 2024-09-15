package captcha

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = base64Captcha.DefaultMemStore

func DriverAudioFunc() (id, b64s string, answer string, err error) {
	driver := base64Captcha.DefaultDriverAudio
	captcha := base64Captcha.NewCaptcha(driver, store)
	return captcha.Generate()
}

func DriverStringFunc() (id, b64s string, answer string, err error) {
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
	captcha := base64Captcha.NewCaptcha(driver.ConvertFonts(), store)
	return captcha.Generate()
}

func DriverDigitFunc() (id, b64s string, answer string, err error) {
	driver := base64Captcha.DefaultDriverDigit
	captcha := base64Captcha.NewCaptcha(driver, store)
	return captcha.Generate()
}

// Verify captcha answer directly
func Verify(id, answer string, clear bool) bool {
	return store.Verify(id, answer, clear)
}
