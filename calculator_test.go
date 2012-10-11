package cal_test

import (
	"github.com/datastream/cal"
	"testing"
)

func TestCal(t *testing.T) {
	exp := "a/b-(c+d-a)*c+(b*c)*(a*c-c)"
	k_v := map[string]float{"a":2,"b":3,"c":4,"d":5}
	rst, err := cal.Cal(exp, k_v)
	if err != nil {
		t.Fatal(err, rst)
	} else {
		t.Log(rst)
	}
}
