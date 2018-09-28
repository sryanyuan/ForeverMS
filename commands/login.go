package commands

import (
	"github.com/spf13/cobra"
)

type loginCommandOptions struct {
	listenClients string
}

func executeLoginCommand(cmd *cobra.Command, options *loginCommandOptions) {

}

func NewLoginCommand() *cobra.Command {
	var options loginCommandOptions
	loginCommand := &cobra.Command{
		Use:   "login",
		Short: "Run login server with specified address",
		Long:  "Run login server with specified address",
		Run: func(cmd *cobra.Command, args []string) {
			executeLoginCommand(cmd, &options)
		},
	}

	loginCommand.Flags().StringVarP(&options.listenClients, "listen-clients", "l", "", "Listen address for serving clients")
	return loginCommand
}
