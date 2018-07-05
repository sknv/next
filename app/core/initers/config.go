package initers

import (
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/sknv/next/app/core/cfg"
)

var (
	config = &cfg.Config{}
)

func init() {
	flagParser := flags.NewParser(config, flags.Default)
	if _, err := flagParser.ParseArgs(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}

func GetConfig() *cfg.Config {
	return config
}
