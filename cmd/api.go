package cmd

import (
	"ft-interview/api"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "api",
	Long:  `api`,
	RunE:  api.Init,
}

func init() {
	RootCmd.AddCommand(apiCmd)
}