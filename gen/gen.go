package gen

import (
	"fmt"
	"sync"
)

var m sync.Mutex
var count = 0

func Email() string {
	m.Lock()
	defer m.Unlock()
	ret := fmt.Sprintf("user-%d@example.com", count)
	count++
	return ret
}
