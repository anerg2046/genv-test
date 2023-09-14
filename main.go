package main

import (
	// 必须在import第一位，否则无法获取正确的配置
	_ "github.com/sakirsensoy/genv/dotenv/autoload"

	"fmt"
	"test-genv/config"
)

func main() {
	fmt.Println(config.App.Host)  // localhost
	fmt.Println(config.App.Port)  // 1234
	fmt.Println(config.App.Debug) // true
}
