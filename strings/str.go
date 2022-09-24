package strings

import (
	"encoding/hex"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func StringToBool(e string) (bool, error) {
	return strconv.ParseBool(e)
}

func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func FormatTimeStr(timeStr string) (string, error) {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02T15:04:05.000Z", timeStr, loc)
	return theTime.Format("2006/01/02 15:04:05"), err
}

// Str2Str 精度处理
func Str2Str(f string, m int) string {
	i := len(f)
	if i <= m {
		for i < m {
			f = "0" + f
			i++
		}
		return "0." + f
	} else {
		sliceTep := make([]string, i)
		for index, value := range f {
			if index == i-m {
				sliceTep = append(sliceTep, ".")
			}
			sliceTep = append(sliceTep, string(value))
		}
		return strings.Join(sliceTep, "")
	}
}

// String2Hex 将字符串转十六进制
func String2Hex(value string) string {
	int := new(big.Int)
	int.SetString(value, 0)
	return hex.EncodeToString(int.Bytes())
}
