package cal_test

import (
	"github.com/datastream/cal"
	"testing"
)

func fn(key string) float32 {
	var rst float32
	switch key {
	case "a":
		{
			rst = 2
		}
	case "b":
		{
			rst = 8
		}
	case "c":
		{
			rst = 7
		}
	case "d":
		{
			rst = 9
		}
	}
	return rst
}
func TestCal(t *testing.T) {
	exp := "a/b-(c+d-a)*c+(b*c)*(a*c-c)"
	rst, err := cal.Cal(exp, fn)
	if err != nil {
		t.Fatal(err, rst)
	} else {
		t.Log(rst)
	}
}
