package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/ynwcel/gfxdemo/internal/cmd"
)

var (
	appVersion = "0.0.1.240709"
)

func main() {
	if err := cmd.NewWithVersion(appVersion).Run(); err != nil {
		panic(err)
	}
}
