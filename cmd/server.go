package cmd

import (
	"github.com/fahribaharudin/petstore_restapi/app"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve, start the server app to serve http",
	Long:  `Serve, start the server app to serve http`,
	Run: func(cmd *cobra.Command, args []string) {
		var app = app.Kernel{}
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
