package util

import (
	"fmt"
	"time"
)

func SelectTest()  {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		var t int64
		for {
			t = time.Now().Unix()
			if t%2 != 0 {
				c1 <- int(t)
			} else {
				c2 <- int(t)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for {
		select {
		case t1 := <-c1:
			fmt.Println("c1:", t1)
		case t2 := <-c2:
			fmt.Println("c2:", t2)
		}
	}
}
