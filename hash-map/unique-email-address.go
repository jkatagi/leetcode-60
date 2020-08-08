/**
* Runtime: 24 ms, faster than 8.87% of Go online submissions for Unique Email Addresses.
* Memory Usage: 6.4 MB, less than 5.64% of Go online submissions for Unique Email Addresses.
 */
func numUniqueEmails(emails []string) int {
	uniq := make(map[string]bool, len(emails))
	for _, email := range emails {
		var validEmail, validLocal string
		var atIdx int
		for i, r := range email {
			s := string(r)
			if s == "@" {
				atIdx = i
				break
			} else if s != "." {
				validLocal += s
			}
		}
		plusIdx := len(validLocal)
		for i, r := range validLocal {
			s := string(r)
			if s == "+" {
				plusIdx = i
				break
			}
		}
		validEmail = validLocal[:plusIdx] + email[atIdx:]
		uniq[validEmail] = true
	}
	var nums int
	for _, ok := range uniq {
		if ok {
			nums += 1
		}
	}
	return nums
}
