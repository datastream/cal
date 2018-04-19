package cal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

func Cal(s string, vars map[string]interface{}) (float64, error) {
	items := Parser(s)
	var opstack []string
	var calstack []float64
	for _, item := range items {
		switch item {
		case "-":
			fallthrough
		case "+":
			fallthrough
		case "/":
			fallthrough
		case "||":
			fallthrough
		case "&&":
			fallthrough
		case "!=":
			fallthrough
		case ">":
			fallthrough
		case ">=":
			fallthrough
		case "<":
			fallthrough
		case "<=":
			fallthrough
		case "==":
			fallthrough
		case "*":
			{
				for {
					if checkPeroption(opstack, item) {
						rst, err := cal2(calstack,
							opstack[len(opstack)-1])
						if err != nil {
							return 0, err
						}
						if len(calstack) > 1 {
							calstack = calstack[:len(calstack)-2]
						} else {
							calstack = calstack[:0]
						}
						calstack = append(calstack, rst)
						opstack = opstack[:len(opstack)-1]
					} else {
						break
					}
				}
				opstack = append(opstack, item)
			}
		case "(":
			{
				opstack = append(opstack, item)
			}
		case ")":
			{
				for {
					if opstack[len(opstack)-1] == "(" {
						opstack = opstack[:len(opstack)-1]
						break
					}
					rst, err := cal2(calstack,
						opstack[len(opstack)-1])
					if err != nil {
						return 0, err
					}
					if len(calstack) > 1 {
						calstack = calstack[:len(calstack)-2]
					} else {
						calstack = calstack[:0]
					}
					calstack = append(calstack, rst)
					opstack = opstack[:len(opstack)-1]
				}
			}
		default:
			{
				v, ok := vars[item]
				if ok {
					value := getValue(v)
					calstack = append(calstack, value)
				} else {
					value, err := strconv.ParseFloat(item, 64)
					if err == nil {
						calstack = append(calstack, value)
						continue
					}
					return 0, errors.New("miss value of " + item)
				}
			}
		}
	}
	for {
		if len(opstack) == 0 || len(calstack) == 1 {
			break
		}
		rst, err := cal2(calstack, opstack[len(opstack)-1])
		if err != nil {
			return 0, err
		}
		if len(calstack) > 1 {
			calstack = calstack[:len(calstack)-2]
		} else {
			calstack = calstack[:0]
		}
		calstack = append(calstack, rst)
		opstack = opstack[:len(opstack)-1]
	}
	if len(calstack) > 1 {
		return calstack[0], errors.New("error exp")
	}
	if len(calstack) == 0 {
		return 0, errors.New("nothing to calculate")
	}
	return calstack[len(calstack)-1], nil
}

func getValue(v interface{}) float64 {
	var value float64
	switch v.(type) {
	case int:
		{
			d, _ := v.(int)
			value = float64(d)
		}
	case int8:
		{
			d, _ := v.(int8)
			value = float64(d)
		}
	case int16:
		{
			d, _ := v.(int16)
			value = float64(d)
		}
	case int32:
		{
			d, _ := v.(int32)
			value = float64(d)
		}
	case int64:
		{
			d, _ := v.(int64)
			value = float64(d)
		}
	case uint8:
		{
			d, _ := v.(uint8)
			value = float64(d)
		}
	case uint16:
		{
			d, _ := v.(uint16)
			value = float64(d)
		}
	case uint32:
		{
			d, _ := v.(uint32)
			value = float64(d)
		}
	case uint64:
		{
			d, _ := v.(uint64)
			value = float64(d)
		}
	case float64:
		{
			value, _ = v.(float64)
		}
	case float32:
		{
			d, _ := v.(float32)
			value = float64(d)
		}
	}
	return value
}
func cal2(calstack []float64, op string) (float64, error) {
	var exp1 float64
	var exp2 float64
	switch len(calstack) {
	case 1:
		{
			exp1 = 0
			exp2 = calstack[len(calstack)-1]
		}
	case 0:
		{
			exp1 = 0
			exp2 = 0
		}
	default:
		{
			exp1 = calstack[len(calstack)-2]
			exp2 = calstack[len(calstack)-1]
		}
	}
	switch op {
	case "||":
		{
			if (exp1 > 0) || (exp2 > 0) {
				return 1, nil
			} else {
				return 0, nil
			}
		}
	case "&&":
		{
			if (exp1 > 0) && (exp2 > 0) {
				return 1, nil
			} else {
				return 0, nil
			}
		}
	case "==":
		{
			if exp1 == exp2 {
				return 1, nil
			} else {
				return 0, nil
			}
		}
	case ">=":
		{
			if exp1 >= exp2 {
				return 1, nil
			} else {
				return 0, nil
			}
		}
	case "<=":
		{
			if exp1 <= exp2 {
				return 1, nil
			} else {
				return 0, nil
			}
		}
	case "!=":
		{
			if exp1 != exp2 {
				return 1, nil
			} else {
				return 0, nil
			}
		}
	case ">":
		{
			if exp1 > exp2 {
				return 1, nil
			} else {
				return 0, nil
			}
		}
	case "<":
		{
			if exp1 < exp2 {
				return 1, nil
			} else {
				return 0, nil
			}
		}
	case "+":
		{
			return exp1 + exp2, nil
		}
	case "-":
		{
			return exp1 - exp2, nil
		}
	case "*":
		{
			return exp1 * exp2, nil
		}
	case "/":
		{
			if exp2 == 0 {
				return 0, errors.New("Div by zero")
			}
			return exp1 / exp2, nil
		}
	}
	return 0, nil
}
func checkPeroption(opstack []string, op string) bool {
	i := len(opstack)
	if i > 0 {
		if (opstack[i-1] == "-" || opstack[i-1] == "+" || opstack[i-1] == "*" || opstack[i-1] == "/") && (op == "-" || op == "+") {
			return true
		}
		if (opstack[i-1] == "*" || opstack[i-1] == "/") && (op == "*" || op == "/") {
			return true
		}
		if (opstack[i-1] == "==" || opstack[i-1] == "!=" || opstack[i-1] == ">=" || opstack[i-1] == "<=" || opstack[i-1] == "<" || opstack[i-1] == ">" || opstack[i-1] == "-" || opstack[i-1] == "+" || opstack[i-1] == "*" || opstack[i-1] == "/") && (op == "==" || op == "!=" || op == ">=" || op == "<=" || op == ">" || op == "<") {
			return true
		}
		if (opstack[i-1] == "==" || opstack[i-1] == "!=" || opstack[i-1] == ">=" || opstack[i-1] == "<=" || opstack[i-1] == "<" || opstack[i-1] == ">" || opstack[i-1] == "-" || opstack[i-1] == "+" || opstack[i-1] == "*" || opstack[i-1] == "/" || opstack[i-1] == "&&" || opstack[i-1] == "||") && (op == "&&" || op == "||") {
			return true
		}
	}
	return false
}
func Parser(s string) []string {
	var items []string
	var sc scanner.Scanner
	sc.Init(strings.NewReader(s))
	var tok rune
	for tok != scanner.EOF {
		tok = sc.Scan()
		item := sc.TokenText()
		if tok != scanner.EOF {
			i := len(items)
			if i > 0 {
				if item == "=" && (items[i-1] == "!" || items[i-1] == ">" || items[i-1] == "<" || items[i-1] == "=") {
					items[i-1] = fmt.Sprintf("%s%s", items[i-1], item)
					continue
				}
				if item == "|" && (items[i-1]) == "|" {
					items[i-1] = fmt.Sprintf("%s%s", items[i-1], item)
					continue
				}
				if item == "&" && (items[i-1]) == "&" {
					items[i-1] = fmt.Sprintf("%s%s", items[i-1], item)
					continue
				}
			}
			items = append(items, item)
		}
	}
	return items
}
