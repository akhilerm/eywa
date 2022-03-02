package cmd

import (
	"log"

	"eywa/pkg/balance"

	"github.com/spf13/cobra"
)

var config balance.Config

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the splitwise reminder application",
	Run: func(cmd *cobra.Command, args []string) {
		config.PayeeAddress = args[0]
		if err := balance.Run(config); err != nil {
			log.Fatal(err)
		}
	},
	Args: cobra.ExactValidArgs(1),
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.PersistentFlags().StringVar(&config.PayeeName, "name", "Test User", "name of the payee")
	startCmd.PersistentFlags().StringVarP(&config.Port, "port", "p", ":8080", "port to start the server")
	startCmd.PersistentFlags().BoolVar(&config.StartServer, "server", true, "to start the URL server")
}
