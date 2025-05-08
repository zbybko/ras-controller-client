package diagnostics

import "ras/utils"

func Traceroute(addr string) (string, error) {
	if addr == "" {
		return "", ErrEmptyAddress
	}
	output, err := utils.Execute("traceroute", addr)
	return string(output), err
}
