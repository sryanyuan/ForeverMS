package main

import (
	"github.com/ngaut/log"
	"github.com/spf13/cobra"
	"github.com/sryanyuan/ForeverMS/commands"
)

func main() {
	cmdEntry := &cobra.Command{Use: "ForeverMS"}
	cmdEntry.AddCommand(commands.NewLoginCommand())
	cmdEntry.AddCommand(commands.NewChannelCommand())
	if err := cmdEntry.Execute(); nil != err {
		log.Errorf("Command error: %v", err)
	}
}
