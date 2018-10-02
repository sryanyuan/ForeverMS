package commands

import (
	"syscall"

	"github.com/juju/errors"
	"github.com/ngaut/log"
	"github.com/spf13/cobra"
	"github.com/sryanyuan/ForeverMS/core/gosync"
	"github.com/sryanyuan/ForeverMS/core/server/channel"
)

type channelCommandOptions struct {
	config string
}

func executeChannelCommand(cmd *cobra.Command, options *channelCommandOptions) {
	var err error
	// Parse config file
	var config channel.Config
	if err = config.LoadFromFile(options.config); nil != err {
		log.Errorf("Load config from file [%s] error: %v",
			options.config, err)
		return
	}
	server := channel.NewChannelServer(&config)
	ctx := gosync.NewContextWithCancel()
	if err = server.Serve(ctx); nil != err {
		log.Errorf("Channel server serve error: %v",
			errors.ErrorStack(err))
		return
	}
	log.Infof("Channel server serve at %v", config.ListenClients)

	sig := gosync.WaitSignals(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	log.Infof("Receive signal %v", sig)
	server.Stop()
	ctx.Cancel()
	ctx.Wait()
}

func NewChannelCommand() *cobra.Command {
	var options channelCommandOptions
	channelCommand := &cobra.Command{
		Use:   "channel",
		Short: "Run channel server with specified address",
		Long:  "Run channel server with specified address",
		Run: func(cmd *cobra.Command, args []string) {
			executeChannelCommand(cmd, &options)
		},
	}

	channelCommand.Flags().StringVarP(&options.config, "config", "c", "", "Config file path")
	return channelCommand
}
