package eth

import (
	"regexp"
)

func CheckEthAddress(value string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if re.MatchString(value) {
		return true
	}
	return false
}
