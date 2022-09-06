package tools

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	pkgstrings "bls/pkg/strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// CompareHashAndPassword 比较hash
func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetGID 获取当前协程id
func GetGID() (n uint64, err error) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err = strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return
	}
	return
}

// GetLocalNetDeviceNames 获取网卡名
func GetLocalNetDeviceNames() (names []string, err error) {
	baseNicPath := "/sys/class/net/"
	cmd := exec.Command("ls", baseNicPath)
	buf, err := cmd.Output()
	if err != nil {
		return
	}
	output := string(buf)
	str := ""
	for _, device := range strings.Split(output, "\n") {
		if len(device) > 1 {
			if device != "lo" {
				str += device + "|"
			}
		}
	}
	names = strings.Split(str[:len(str)-1], "|")
	return
}

func Hmac(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func IsStringEmpty(str string) bool {
	return strings.Trim(str, " ") == ""
}

func GetUUID() string {
	u := uuid.New()
	return strings.ReplaceAll(u.String(), "-", "")
}

// func PathExists(path string) bool {
// 	_, err := os.Stat(path)
// 	if err == nil {
// 		return true
// 	}

// 	if os.IsNotExist(err) {
// 		return false
// 	}

// 	return false
// }

func Base64ToImage(imageBase64 string) ([]byte, error) {
	image, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func GetDirFiles(dir string) ([]string, error) {
	dirList, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	filesRet := make([]string, 0)
	for _, file := range dirList {
		if file.IsDir() {
			files, err := GetDirFiles(dir + string(os.PathSeparator) + file.Name())
			if err != nil {
				return nil, err
			}

			filesRet = append(filesRet, files...)
		} else {
			filesRet = append(filesRet, dir+string(os.PathSeparator)+file.Name())
		}
	}
	return filesRet, nil
}

//slice去重
func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

func IdsStrToIdsIntGroup(keys string) []int {
	IDS := make([]int, 0)
	ids := strings.Split(keys, ",")
	for i := 0; i < len(ids); i++ {
		ID, _ := pkgstrings.StringToInt(ids[i])
		IDS = append(IDS, ID)
	}
	return IDS
}

//截取字符串-将字符串前面的0去掉
func SubByZero(str string) string {
	arr := []byte(str)
	for i, str1 := range arr {
		if str1 != 48 {
			str = string(arr[i:])
			break
		}
	}
	return str
}

//精度处理
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

//浮点数精度处理转换整数
func Float2Int(f float64, decimal int) uint64 {
	value := ChangeNumber(f, 8)
	split := strings.Split(value, ".")
	if len(split) == 2 {
		decimal -= len(split[1])
	}
	value = strings.ReplaceAll(value, ".", "")
	for i := 0; i < decimal; i++ {
		value += "0"
	}
	i, _ := strconv.ParseInt(value, 10, 64)
	return uint64(i)
}

//保留小数位
func ChangeNumber(f float64, m int) string {
	n := strconv.FormatFloat(f, 'f', -1, 64)
	if n == "" {
		return ""
	}
	if m >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if len(newn) < 2 || m >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:m]
}

//将字符串float进行精度处理
func Float2String(value string, decimal int) string {
	split := strings.Split(value, ".")
	if len(split) > 1 {
		decimal -= len(split[1])
		value = strings.ReplaceAll(value, ".", "")
	}
	for decimal > 0 {
		value += "0"
		decimal -= 1
	}
	return SubByZero(value)
}

//精度处理
func ChangeInt2Float(f string, m int) string {
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

//精度处理
func ChangeFloat2Float(f string, m int) string {
	if strings.Contains(f, ".") {
		index := strings.Index(f, ".")
		sub := f[index+1 : len(f)]
		if len(sub) > m {
			f = f[0 : index+m+1]
		} else {
			for m > len(sub) {
				f += "0"
				m--
			}
		}
	} else {
		f += "."
		for m > 0 {
			f += "0"
			m--
		}
	}
	return f
}

func HexToBigInt(hex string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(hex[2:], 16)
	return n
}

func ToBigFloat(str string) *big.Float {
	f, _, _ := big.ParseFloat(str, 10, 256, big.ToNearestEven)
	return f
}

func BigIntToETHBigFloat(f *big.Int) *big.Float {
	v := ToBigFloat(f.String())
	z := big.NewFloat(0)
	z.Quo(v, big.NewFloat(math.Pow10(18)))
	return z
}

func BigIntToUsdtBigFloat(f *big.Int) *big.Float {
	v := ToBigFloat(f.String())
	z := big.NewFloat(0)
	z.Quo(v, big.NewFloat(math.Pow10(6)))
	return z
}

func ToHex(b []byte) string {
	hex := hex.EncodeToString(b)
	if len(hex) == 0 {
		hex = "0"
	}
	return "0x" + hex
}

//小数点位数校验
func CheckFloat(value string, decimal int) bool {
	values := strings.Split(value, ".")
	if len(values) == 2 {
		if len(values[1]) > decimal {
			return false
		}
	}
	return true
}

func GetOrderid() string {
	var t = time.Now().Unix()
	var s = time.Unix(t, 0).Format("20060102150405") + fmt.Sprint(time.Now().UnixNano())
	return s
}

//结构体转为map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
