package mmcli

import "fmt"

const JsonOutputFlag = "--output-json"
const ListModemsFlag = "--list-modems"

func ModemFlag(modem string) string {
	return fmt.Sprintf("--modem='%s'", modem)
}

func SimFlag(sim string) string {
	return fmt.Sprintf("--sim='%s'", sim)
}

func BearerFlag(bearer string) string {
	return fmt.Sprintf("--bearer='%s'", bearer)
}
