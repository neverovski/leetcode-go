package solution_020

import (
	stack "leetcode/go/structures"
)

func IsValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	pattern := map[byte]byte{
		'[': ']',
		'(': ')',
		'{': '}',
	}
	symbolStack := stack.New()

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
