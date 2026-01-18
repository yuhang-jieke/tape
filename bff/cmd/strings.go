package cmd

import "strings"

func StringJoin(str ...string) string {
	return strings.Join(str, ",")
}
