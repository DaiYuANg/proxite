package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
	"proxite/cmd"
)

func main() {
	cobra.CheckErr(cmd.Execute())
}
