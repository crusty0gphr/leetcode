package longest_common_prefix

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return prefix
	}
	prefix = strs[0]
	for _, str := range strs {
		if len(prefix) > len(str) {
			prefix = str
		}
	}
	for i := len(prefix); i >= 0; i-- {
		for _, str := range strs {
			if strings.Index(str, prefix) != 0 {
				prefix = prefix[:i]
			}
		}

	}
	return prefix
}
