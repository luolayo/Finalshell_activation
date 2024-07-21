package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/ebfe/keccak"
	"os"
	"strings"
)

var newEntry = "127.0.0.1 *.hostbuf.com"

// SetHosts Set hosts
func SetHosts() {
	var hostsFilePath string
	switch IsWindows() {
	case true:
		hostsFilePath = "C:\\Windows\\System32\\drivers\\etc\\hosts"
		// windows
		fmt.Println("Windows")
	case false:
		fmt.Println("Linux")
		hostsFilePath = "/etc/hosts"
	// linux
	default:
		hostsFilePath = "/etc/hosts"
		fmt.Println("get hosts file path error")
	}
	file, err := os.Open(hostsFilePath)
	if err != nil {
		fmt.Printf("无法打开hosts文件: %v\n", err)
		return
	}
	scanner := bufio.NewScanner(file)
	entryExists := false
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == newEntry {
			entryExists = true
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("读取hosts文件时出错: %v\n", err)
		return
	}

	if entryExists {
		fmt.Println("该条目已经存在于hosts文件中。")
		return
	}

	// 将新条目添加到hosts文件
	file, err = os.OpenFile(hostsFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("无法打开hosts文件以追加内容: %v\n", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭hosts文件时出错。")
			return
		}
	}(file)

	if _, err = file.WriteString(newEntry + "\n"); err != nil {
		fmt.Printf("写入hosts文件时出错: %v\n", err)
		return
	}

	fmt.Println("成功添加新条目到hosts文件。")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭hosts文件时出错。")
			return
		}
	}(file)
}

// IsWindows Determine whether the environment is Windows or Linux
func IsWindows() bool {
	return os.IsPathSeparator('\\')
}

func main() {
	if !isAdmin() {
		fmt.Println("请求管理员权限...")
		fmt.Println("清以管理员身份运行本程序！")
		fmt.Println("\n按回车键退出...")
		_, _ = fmt.Scanln()
		return
	}

	fmt.Println("程序以管理员身份运行")

	SetHosts()
	fmt.Println("请输入机器码：")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	code := scanner.Text()
	values := generateValues(code)
	fmt.Println("旧版本：")
	fmt.Println("高级版: " + values[0])
	fmt.Println("专业版: " + values[1])
	fmt.Println("新版本：")
	fmt.Println("高级版: " + values[2])
	fmt.Println("专业版: " + values[3])
	// 等待用户输入以防止程序自动退出
	fmt.Println("\n按回车键退出...")
	_, _ = fmt.Scanln()
}

// md5Hash generates an MD5 hash for the given message
func md5Hash(msg string) string {
	hash := md5.Sum([]byte(msg))
	return hex.EncodeToString(hash[:])
}

// keccak384Hash generates a Keccak-384 hash for the given message
func keccak384Hash(msg string) string {
	hash := keccak.New384()
	hash.Write([]byte(msg))
	return hex.EncodeToString(hash.Sum(nil))
}

// generateValues generates the values for both old and new versions
func generateValues(code string) []string {
	versionOld := []string{
		md5Hash("61305" + code + "8552")[8:24],
		md5Hash("2356" + code + "13593")[8:24],
	}
	versionNew := []string{
		keccak384Hash(code + "hSf(78cvVlS5E")[12:28],
		keccak384Hash(code + "FF3Go(*Xvbb5s2")[12:28],
	}
	fmt.Println("versionOld: ", keccak384Hash(code+"hSf(78cvVlS5E"))
	fmt.Println("versionNew: ", keccak384Hash(code+"FF3Go(*Xvbb5s2"))
	return append(versionOld, versionNew...)

}

func isAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}
