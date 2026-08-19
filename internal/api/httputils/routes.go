package httputils

import (
	"fmt"
	"strings"
)

func CreateRoute(method, prefix, path string) string {
	if len(prefix) == 0 {
		if strings.HasPrefix(path, "/") {
			return fmt.Sprintf("%s %s", method, path)
		}
		return fmt.Sprintf("%s /%s", method, path)
	}

	if strings.HasPrefix(path, "/") {
		return fmt.Sprintf("%s %s%s", method, prefix, path)
	}

	return fmt.Sprintf("%s %s/%s", method, prefix, path)
}
