package node

import (
	"testing"
)

func TestMtime(t *testing.T) {
	// diff := int64(time.Now().Unix()*1000) - 1624699085486
	day := ExpireDay(1624699085486)
	t.Log(day)

}
