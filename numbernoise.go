package numbernoise

import (
	"math/rand"
	"time"
)

// p is the precision of the tick - e.g. a unique number every p
// l is the lifetime of the ticker, after which the returned channel will be closed
func RandomNSeconds(p, l time.Duration) <-chan int {
	c := time.NewTicker(p)
	w := time.NewTimer(l)

	nums := make(chan int)

	go func() {
		ok := true
		for ok {
			select {
			case tick := <-c.C:
				num := int(int64(tick.UnixMilli())) * rand.Int()
				nums <- num
			case <-w.C:
				close(nums)
				c.Stop()
				w.Stop()
				ok = false
				return
			}
		}
	}()

	return nums
}

func EvenOdds(p, l time.Duration) (<-chan int, <-chan int) {
	evens := make(chan int)
	odds := make(chan int)

	go func() {
		nums := RandomNSeconds(p, l)
		for {
			rn, nok := <-nums

			if !nok {
				close(evens)
				close(odds)
				return
			}

			if rn%2 == 0 {
				evens <- rn
			} else {
				odds <- rn
			}
		}
	}()

	return evens, odds
}
