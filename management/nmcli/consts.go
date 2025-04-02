package nmcli

import (
	"fmt"
	"strings"
)

const (
	terseFlag = "--terse"
)

func getFields(fields ...string) string {
	return fmt.Sprintf("--get-values=%s", strings.Join(fields, ","))
}
