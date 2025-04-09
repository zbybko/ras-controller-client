package diagnostics

import "github.com/spf13/viper"

func setupDiagnostics() {
	defaultPingAddress = viper.GetString("ping.default_address")
}
