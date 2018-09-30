package commands

import (
	"syscall"

	"github.com/juju/errors"
	"github.com/ngaut/log"
	"github.com/spf13/cobra"
	"github.com/sryanyuan/ForeverMS/core/gosync"
	"github.com/sryanyuan/ForeverMS/core/server/login"
)

type loginCommandOptions struct {
	config string
}

func executeLoginCommand(cmd *cobra.Command, options *loginCommandOptions) {
	var err error
	// Parse config file
	var config login.Config
	if err = config.LoadFromFile(options.config); nil != err {
		log.Errorf("Load config from file [%s] error: %v",
			options.config, err)
		return
	}
	server := login.NewLoginServer(&config)
	ctx := gosync.NewContextWithCancel()
	if err = server.Serve(ctx); nil != err {
		log.Errorf("Login server serve error: %v",
			errors.ErrorStack(err))
		return
	}
	log.Infof("Login server serve at %v", config.ListenClients)

	sig := gosync.WaitSignals(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	log.Infof("Receive signal %v", sig)
	server.Stop()
	ctx.Cancel()
	ctx.Wait()
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

	loginCommand.Flags().StringVarP(&options.config, "config", "c", "", "Config file path")
	return loginCommand
}
