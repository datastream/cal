package cal

import (
	"errors"
)

func Cal(exp string, k_v map[string]interface{}) (float64, error) {
	exps := Parser(exp)
	var opstack []string
	var calstack []float64
	for i := range exps {
		switch exps[i] {
		case "-":
			{
				fallthrough
			}
		case "+":
			{
				fallthrough
			}
		case "/":
			{
				fallthrough
			}
		case "*":
			{
				for {
					if check_peroption(opstack, exps[i]) {
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
				opstack = append(opstack, exps[i])
			}
		case "(":
			{
				opstack = append(opstack, exps[i])
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
				v, ok := k_v[exps[i]]
				if ok {
					value := get_value(v)
					calstack = append(calstack, value)
				} else {
					return 0, errors.New("miss value of " + exps[i])
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
	return calstack[len(calstack)-1], nil
}

func get_value(v interface{}) float64 {
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
func check_peroption(opstack []string, op string) bool {
	if len(opstack) > 0 {
		if (opstack[len(opstack)-1] == "-" ||
			opstack[len(opstack)-1] == "+" ||
			opstack[len(opstack)-1] == "*" ||
			opstack[len(opstack)-1] == "/") &&
			(op == "-" || op == "+") {
			return true
		}
		if (opstack[len(opstack)-1] == "*" ||
			opstack[len(opstack)-1] == "/") &&
			(op == "*" || op == "/") {
			return true
		}
	}
	return false
}
func Parser(exp string) []string {
	var tokens []string
	var token []byte
	for i := range exp {
		switch exp[i] {
		case '*':
			{
				fallthrough
			}
		case '/':
			{
				fallthrough
			}
		case '-':
			{
				fallthrough
			}
		case '+':
			{
				fallthrough
			}
		case '(':
			{
				fallthrough
			}
		case ')':
			{
				if len(token) > 0 {
					tokens = append(tokens, string(token))
					token = token[:0]
				}
				tokens = append(tokens, string([]byte{exp[i]}))
			}
		case ' ':
			{
			}
		default:
			{
				token = append(token, exp[i])
			}
		}
	}
	if len(token) > 0 {
		tokens = append(tokens, string(token))
	}
	return tokens
}
