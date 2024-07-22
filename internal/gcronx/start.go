package gcronx

import (
	"fmt"
	"time"
)

func Start() error {
	fmt.Println("gcronx-start")
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()
	for v := range ticker.C {
		fmt.Println("gcronx.ticker:", v)
	}
	return nil
}
