package echo

import (
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	var buf [23]byte
	at := time.Now()
	t.Log(at.String())
	TimeFormat(buf[:], at, false)
	t.Log(string(buf[:19]))
	TimeFormat(buf[:], at, true)
	t.Log(string(buf[:23]))
}

func TestTrimSpace(t *testing.T) {
	t.Log(trimspace([]byte("")))
	t.Log(trimspace([]byte(" ")))
	t.Log(trimspace([]byte("a ")))
	t.Log(trimspace([]byte(" a")))
	t.Log(trimspace([]byte(" a\t")))
}
