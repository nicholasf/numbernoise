package numbernoise

import (
	"fmt"
	"testing"
	"time"
)

func TestRandomNSeconds(t *testing.T) {
	t.Run("Test that we get a stream of numbetrs > 0", func(t *testing.T) {
		nums := RandomNSeconds(time.Millisecond, 50*time.Millisecond)
		ok := true

		for ok {
			n, ok := <-nums
			if !ok {
				break
			}

			if n == 0 {
				t.Logf("Zero value received from random number generator while the channel was open: %d\n", n)
				t.Fail()
			}
		}
	})
}

func TestEvenOdds(t *testing.T) {
	evens, odds := EvenOdds(time.Millisecond, 2*time.Second)
	e := make([]int, 0)
	o := make([]int, 0)

	ok := true

	for ok {
		select {
		case en, eok := <-evens:
			if !eok {
				fmt.Println("breaking from evens")
				ok = false
				// break
			}

			e = append(e, en)
		case on, ook := <-odds:
			if !ook {
				fmt.Println("breaking from odds")
				ok = false
				// break
			}

			o = append(o, on)
		}
	}

	t.Run("Only even numbers in evens", func(t *testing.T) {
		for _, i := range e {
			if i%2 != 0 {
				t.Logf("Uneven number in evens: %d", i)
				t.Fail()
			}
		}
	})

	t.Run("Only odd numbers in odds", func(t *testing.T) {
		for _, i := range o {
			if i%2 > 0 || i%2 < 0 {
			} else {
				t.Logf("Even number in odds: %d", i)
				t.Fail()
			}
		}
	})

}
