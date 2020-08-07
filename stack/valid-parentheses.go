/**
* Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
* Memory Usage: 2.1 MB, less than 31.25% of Go online submissions for Valid Parentheses.
 */

func isValid(s string) bool {
	stack := make([]string, 0, len(s))
	if s == "" {
		return true
	}
	for i := range s {
		switch string(s[i]) {
		case "(":
			stack = append(stack, "(")
		case ")":
			if len(stack) == 0 {
				// no open bracket
				return false
			} else if string(stack[len(stack)-1]) == "(" {
				// much open bracket
				stack = stack[0 : len(stack)-1] // remove "("
			} else {
				// does not much open bracket
				return false
			}
		case "{":
			stack = append(stack, "{")
		case "}":
			if len(stack) == 0 {
				// no open bracket
				return false
			} else if string(stack[len(stack)-1]) == "{" {
				// much open bracket
				stack = stack[0 : len(stack)-1] // remove "{"
			} else {
				// does not much open bracket
				return false
			}
		case "[":
			stack = append(stack, "[")
		case "]":
			if len(stack) == 0 {
				// no open bracket
				return false
			} else if string(stack[len(stack)-1]) == "[" {
				// much open bracket
				stack = stack[0 : len(stack)-1] // remove "["
			} else {
				// does not much open bracket
				return false
			}
		default:
			continue
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
