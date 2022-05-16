package main

import (
	"bytes"
	"fmt"
	"github.com/mark8s/go-lib/go-bindata/common"
	"github.com/spf13/viper"
)

//go:generate go-bindata.exe -o=./common/config.go -pkg=common .env

func init() {
	fileObj, err := common.Asset(".env")
	if err != nil {
		fmt.Printf("Asset file err:%v\n", err)
		return
	}
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(fileObj))
	if err != nil {
		fmt.Printf("Read Config err:%v\n", err)
		return
	}
}

func main() {
	fmt.Println("用户为:", viper.GetString("USER"))
	fmt.Println("密码为:", viper.GetString("PASS"))
	fmt.Println("地址:", viper.GetString("ADDRESS"))
}
