package eth

import (
	"regexp"
	"strconv"
)

func CheckEthAddress(value string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if re.MatchString(value) {
		return true
	}
	return false
}

// Parse16TO10 16進数 => 10進数
func Parse16TO10(r string) (int64, error) {
	val := r[2:]
	i, err := strconv.ParseInt(val, 16, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}
