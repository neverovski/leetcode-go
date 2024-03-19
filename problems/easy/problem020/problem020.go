package problem020

import "leetcode/go/utils/stack"

const DIVISOR = 2

func isValid(s string) bool {
	if len(s) == 0 || len(s)%DIVISOR == 1 {
		return false
	}

	pattern := map[byte]byte{
		'[': ']',
		'(': ')',
		'{': '}',
	}
	symbolStack := stack.New()

	return processSymbols(s, pattern, symbolStack)
}

func processSymbols(s string, pattern map[byte]byte, symbolStack *stack.Stack) bool {
	for i := 0; i < len(s); i++ {
		if _, ok := pattern[s[i]]; ok {
			symbolStack.Push(s[i])
		} else if symbolStack.Len() == 0 || pattern[symbolStack.Peek().(byte)] != s[i] {
			return false
		} else {
			symbolStack.Pop()
		}
	}

	return symbolStack.Len() == 0
}
