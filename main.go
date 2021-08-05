package main

import (
	"fmt"
	"xj/viper_test/util"
)

func main() {
	fmt.Println("viper test main func")
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Println("cannot load config")
		fmt.Println(err)
		return
	}

	fmt.Println(config)
	fmt.Println(util.EnvGet("DB_PORT"))
	fmt.Println(util.EnvGet("PROMOTION_SERVICE_HOST"))
}
