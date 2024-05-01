package parallel

import (
	"fmt"
	"testing"
	"time"
)

//	func TestSometing(t *testing.T) {
//		t.Parallel()
//	}
//
//	func TestA(t *testing.T) {
//		t.Parallel()
//	}
func TestB(t *testing.T) {
	fmt.Println("setup")
	defer fmt.Println("defrred teardown")
	t.Run("group", func(t *testing.T) {
		t.Run("sub1", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second)
			fmt.Println("sub1 done")
		})
		t.Run("sub2", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second)
			fmt.Println("sub2 done")
		})
	})

	fmt.Println("teardown")
}
