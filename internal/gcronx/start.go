package gcronx

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
)

func Start() error {
	fmt.Println("gcronx-start")
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()
	idx := 0
	for v := range ticker.C {
		fmt.Println("gcronx.ticker:", v)
		if idx >= 3 {
			return gerror.New("ticker-idx>=3")
		}
		idx += 1
	}
	return nil
}
