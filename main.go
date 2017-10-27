package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	wininet, _           = syscall.LoadLibrary("wininet.dll")
	getConnectedState, _ = syscall.GetProcAddress(wininet, "InternetGetConnectedState")
)

const (
	INTERNET_CONNECTION_MODEM      = 0x01 // 本地系统使用调制解调器连接到互联网。
	INTERNET_CONNECTION_LAN        = 0x02 // 本地系统使用局域网连接到互联网。
	INTERNET_CONNECTION_PROXY      = 0x04 // 本地系统使用代理服务器连接到Internet。
	INTERNET_CONNECTION_MODEM_BUSY = 0x08 // 不再使用
	INTERNET_RAS_INSTALLED         = 0x10 // 本地系统已安装RAS。
	INTERNET_CONNECTION_OFFLINE    = 0x20 // 本地系统处于离线模式。
	INTERNET_CONNECTION_CONFIGURED = 0x40 // 本地系统具有到互联网的有效连接，但它可能或可能不是当前连接的。
)

func main() {
	defer syscall.FreeLibrary(wininet)

	var nargs uintptr = 2
	flags := int32(0)
	r1, _, err := syscall.Syscall(
		uintptr(getConnectedState),
		nargs,
		uintptr(unsafe.Pointer(&flags)),
		0,
		0,
	)

	if err != 0 {
		fmt.Printf("Error: %s\n", err)
		return
	}

	if r1 == 1 {
		switch flags {
		case INTERNET_CONNECTION_MODEM:
			fmt.Println("本地系统使用调制解调器连接到互联网。")
		case INTERNET_CONNECTION_LAN:
			fmt.Println("本地系统使用局域网连接到互联网。")
		case INTERNET_CONNECTION_PROXY:
			fmt.Println("本地系统使用代理服务器连接到Internet。")
		case INTERNET_CONNECTION_MODEM_BUSY:
			fmt.Println("不再使用")
		case INTERNET_RAS_INSTALLED:
			fmt.Println("本地系统已安装RAS。")
		case INTERNET_CONNECTION_OFFLINE:
			fmt.Println("本地系统处于离线模式。")
		case INTERNET_CONNECTION_CONFIGURED:
			fmt.Println("本地系统具有到互联网的有效连接，但它可能或可能不是当前连接的。")
		case INTERNET_CONNECTION_LAN + INTERNET_RAS_INSTALLED:
			fmt.Println("在线：0x12")
		default:
			fmt.Printf("flags: %#x\n", flags)
		}
	} else {
		// 拔掉网线
		fmt.Println("没有连网: 检查网络设置")
	}

}
