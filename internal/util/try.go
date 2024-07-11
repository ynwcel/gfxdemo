package util

import "fmt"

func Try(f func() error) (err error) {
	defer func() {
		if exception := recover(); exception != nil {
			if e, ok := exception.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%#v", exception)
			}
		}
	}()
	err = f()
	return
}
