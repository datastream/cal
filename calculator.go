package cal

import (
	"errors"
)

func Cal(exp string, fn func(string) float32) (float32, error) {
	exps := parser(exp)
	var opstack []string
	var calstack []float32
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
			}
		default:
			{
				calstack = append(calstack, fn(exps[i]))
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

func cal2(calstack []float32, op string) (float32, error) {
	var exp1 float32
	var exp2 float32
	switch len(calstack) {
	case 1:
		{
			exp1 = 0
			exp2 = calstack[len(calstack) -1]
		}
	case 0:
		{
			exp1 = 0
			exp2 = 0
		}
	default:
		{
			exp1 = calstack[len(calstack) -2]
			exp2 = calstack[len(calstack) -1]
		}
	}
	switch op {
	case "+":
		{
			return exp1+exp2, nil
		}
	case "-":
		{
			return exp1-exp2, nil
		}
	case "*":
		{
			return exp1*exp2, nil
		}
	case "/":
		{
			if exp2 == 0 {
				return 0, errors.New("Div by zero")
			}
			return exp1/exp2, nil
		}
	}
	return 0, nil
}
func check_peroption(opstack []string, op string) bool {
	if len(opstack) > 0 {
		if (opstack[len(opstack)-1] == "-" || opstack[len(opstack)-1] == "+" || opstack[len(opstack)-1] == "*" || opstack[len(opstack)-1] == "/") && (op == "-"|| op == "+") {
			return true
		}
		if (opstack[len(opstack)-1] == "*" || opstack[len(opstack)-1] == "/") && (op == "*"|| op == "/") {
			return true
		}
	}
	return false
}
func parser(exp string) []string {
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
		case ' ': {
			}
		default:
			{
				token = append(token, exp[i])
			}
		}
	}
	return tokens
}
