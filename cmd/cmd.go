package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "gossip 🐍",
	Short: "show relations of gossip",
}

func Execute() error {
	return rootCmd.Execute()
}
