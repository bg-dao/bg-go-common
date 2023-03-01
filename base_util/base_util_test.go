package baseUtil

import (
	"fmt"
	"testing"
	"time"
)

func TestGenID(t *testing.T) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(1) * time.Millisecond)
		id, err := GenID()
		if err != nil {
			return
		}
		idStr := fmt.Sprintf(`%v`, id)

		// Generate and print, all in one.
		fmt.Printf("id: %v, len:%v\n", id, len(idStr))
	}

}
