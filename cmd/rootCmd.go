package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var (
	configFile string
	app        *fx.App
)

func init() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "config.toml", "Config file path")
}

var rootCmd = &cobra.Command{
	Use:   "proxite",
	Short: "A static + reverse proxy web server",
	PreRun: func(cmd *cobra.Command, args []string) {
		app = container(configFile)
	},
	Run: func(cmd *cobra.Command, args []string) {
		app.Run()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
