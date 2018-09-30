package bob

import (
	"strings"
)

const testVersion = 2

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if isYelling(remark) && isQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	} else if isYelling(remark) {
		return "Whoa, chill out!"
	} else if isQuestion(remark) {
		return "Sure."
	} else if isSilence(remark) {
		return "Fine. Be that way!"
	}
	return "Whatever."
}

func isYelling(remark string) bool {
	if len(remark) != 0 {
		return remark == strings.ToUpper(remark) && remark != strings.ToLower(remark)
	}
	return false
}

func isSilence(remark string) bool {
	if len(remark) == 0 {
		return len(remark) == 0
	}
	return false
}

func isQuestion(remark string) bool {
	if len(remark) != 0 {
		return remark[len(remark)-1:] == "?"
	}
	return false
}
