package logic

import (
	"fmt"
	"testing"
)

func TestLogic(t *testing.T) {
	w := NewWorld()
	tlist := w.TableList()
	for _, tb := range tlist {
		fmt.Println(tb.String())
	}
}

func TestPoker(t *testing.T) {
	ps := NewPoker()
	ps.Wash()
	fmt.Println(ps.String())
}
