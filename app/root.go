package app

import (
	"ezreal.com.cn/ez_crontab/app/crontabd"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
				This application is a tool to generate the needed files
				to quickly create a Cobra application.
				`,
	}
)

// Execute executes the root command.
func Execute() error {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "cfgFile", "./crontab.toml", "config file (default is $HOME/crontab.toml)")
	rootCmd.AddCommand(crontabd.NewCrontabdCmd())
	return rootCmd.Execute()
}
