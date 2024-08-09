package gcronx

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
)

func Start() error {
	fmt.Println("gcronx-start")
	idx := 0
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()
	for v := range ticker.C {
		fmt.Println("gcronx.ticker:", v)
		idx += 1
		if idx >= 5 {
			return gerror.New("gcronx index great 5")
		}
	}
	return nil
}
