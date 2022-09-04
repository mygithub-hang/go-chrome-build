package go_chrome_build

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 控制台颜色文字变量
var (
	Green   = string([]byte{27, 91, 51, 50, 109})
	White   = string([]byte{27, 91, 51, 55, 109})
	Yellow  = string([]byte{27, 91, 51, 51, 109})
	Red     = string([]byte{27, 91, 51, 49, 109})
	Blue    = string([]byte{27, 91, 51, 52, 109})
	Magenta = string([]byte{27, 91, 51, 53, 109})
	Cyan    = string([]byte{27, 91, 51, 54, 109})
	Reset   = string([]byte{27, 91, 48, 109})
)

// IsExist 判断文件是否存在
func IsExist(fileAddr string) bool {
	// 读取文件信息，判断文件是否存在
	_, err := os.Stat(fileAddr)
	if err != nil {
		if os.IsExist(err) { // 根据错误类型进行判断
			return true
		}
		return false
	}
	return true
}

// IsDir 判断目录是否存在
func IsDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// EchoError 彩色打印错误信息
func EchoError(i interface{}) {
	// fmt.Println(green, i, reset)
	// fmt.Println(white, i, reset)
	// fmt.Println(yellow, i, reset)
	// fmt.Println(red, i, reset)
	// fmt.Println(blue, i, reset)
	// fmt.Println(magenta, i, reset)
	// fmt.Println(cyan, i, reset)
	fmt.Println(Magenta, i, Reset)
}

// EchoSuccess 彩色打印成功信息
func EchoSuccess(i interface{}) {
	fmt.Println(Green, i, Reset)
}

// StringToArray 字符串转切片
func StringToArray(str string, sep string) []string {
	return strings.Split(str, sep)
}

func ArrayToString(arr []string, sep string) string {
	str := ""
	if len(arr) == 0 {
		return str
	}
	for _, v := range arr {
		if str == "" {
			str = v
		} else {
			str += sep + v
		}
	}
	return str
}

func StrFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
}

// AskForConfirmation 控制台询问
func AskForConfirmation(s string, tips ...string) (string, error) {
	msg := "[y/*]:"
	if len(tips) > 0 {
		msg = tips[0]
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s%s", s, msg)
		response, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		response = strings.TrimRight(response, "\n")
		response = strings.TrimRight(response, "\r")
		response = strings.TrimRight(response, "\n")
		return response, err
	}
}

// ToUnderScore 驼峰命名转下划线命名
func ToUnderScore(s string) string {
	newStr := ""
	for k, v := range s {
		if k == 0 {
			if v >= 65 && v <= 90 {
				newStr += strings.ToLower(string(v))
			} else {
				newStr += string(v)
			}
		} else {
			if v >= 65 && v <= 90 {
				newStr += "_" + strings.ToLower(string(v))
			} else {
				newStr += string(v)
			}
		}
	}
	return newStr
}

// GetCurrentPath 获取当前文件位置
func GetCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// GetExcPath 获取执行位置
func GetExcPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return strings.Replace(ret, "\\", "/", -1)
}
