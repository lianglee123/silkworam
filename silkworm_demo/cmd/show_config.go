package cmd

import "silkwormDemo/config"

func PrintConfig() {
	config := config.Load()
	pretty.Print(config)
}
