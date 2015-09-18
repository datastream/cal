package cal_test

import (
	"github.com/datastream/cal"
	"testing"
)

func TestCal(t *testing.T) {
	exp := "a/b-(c+d-a)*c+(b*c)*(a*c-c)"
	k_v := map[string]interface{}{"a": 2.0, "b": 3.0, "c": 4.0, "d": 5.1}
	rst, err := cal.Cal(exp, k_v)
	if err != nil {
		t.Fatal(err, rst)
	} else {
		t.Log(rst)
	}
	exp = "a>b"
	rst, err = cal.Cal(exp, k_v)
	if err != nil || rst == 1 {
		t.Fatal(err, rst)
	} else {
		t.Log(rst)
	}
	exp = "a>=b"
	rst, err = cal.Cal(exp, k_v)
	if err != nil || rst == 1 {
		t.Fatal(err, rst)
	} else {
		t.Log(rst)
	}
	exp = "a<b"
	rst, err = cal.Cal(exp, k_v)
	if err != nil || rst == 0 {
		t.Fatal(err, rst)
	} else {
		t.Log(rst)
	}
	exp = "a<=b"
	rst, err = cal.Cal(exp, k_v)
	if err != nil || rst == 0 {
		t.Fatal(err, rst)
	} else {
		t.Log(rst)
	}
	exp = "a!=b"
	rst, err = cal.Cal(exp, k_v)
	if err != nil || rst == 0 {
		t.Fatal(err, rst)
	} else {
		t.Log(rst)
	}
	exp = "a==b"
	rst, err = cal.Cal(exp, k_v)
	if err != nil || rst == 1 {
		t.Fatal(err, rst)
	} else {
		t.Log(rst)
	}
	exp = "0||1"
	rst, err = cal.Cal(exp, k_v)
	if err != nil || rst == 0 {
		t.Fatal(err, rst)
	} else {
		t.Log(rst)
	}
	exp = "1||1"
	rst, err = cal.Cal(exp, k_v)
	if err != nil || rst == 0 {
		t.Fatal(err, rst)
	} else {
		t.Log(rst)
	}
	exp = "1&&1"
	rst, err = cal.Cal(exp, k_v)
	if err != nil || rst == 0 {
		t.Fatal(err, rst)
	} else {
		t.Log(rst)
	}
	exp = "1&&0"
	rst, err = cal.Cal(exp, k_v)
	if err != nil || rst == 1 {
		t.Fatal(err, rst)
	} else {
		t.Log(rst)
	}
}
