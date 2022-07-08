package main

import (
	"fmt"
	"os"

	"github.com/0xPolygonHermez/zkevm-node/config"
	"github.com/0xPolygonHermez/zkevm-node/jsonrpc"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/urfave/cli/v2"
)

const (
	// App name
	appName = "hermez-node"
	// version represents the program based on the git tag
	version = "v0.1.0"
	// commit represents the program based on the git commit
	commit = "dev"
	// date represents the date of application was built
	date = ""
)

const (
	AGGREGATOR   = "aggregator"
	SEQUENCER    = "sequencer"
	RPC          = "rpc"
	SYNCHRONIZER = "synchronizer"
	BROADCAST    = "broadcast-trusted-state"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = version
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     config.FlagCfg,
			Aliases:  []string{"c"},
			Usage:    "Configuration `FILE`",
			Required: false,
		},
		&cli.StringFlag{
			Name:     config.FlagNetwork,
			Aliases:  []string{"n"},
			Usage:    "Network: mainnet, testnet, internaltestnet, local, custom, merge. By default it uses mainnet",
			Required: false,
		},
		&cli.StringFlag{
			Name:    config.FlagNetworkCfg,
			Aliases: []string{"nc"},
			Usage:   "Custom network configuration `FILE` when using --network custom parameter",
		},
		&cli.StringFlag{
			Name:     config.FlagNetworkBase,
			Aliases:  []string{"nb"},
			Usage:    "Base existing network configuration to be merged with the custom configuration passed with --network-cfg, by default it uses internaltestnet",
			Value:    "internaltestnet",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     config.FlagYes,
			Aliases:  []string{"y"},
			Usage:    "Automatically accepts any confirmation to execute the command",
			Required: false,
		},
		&cli.StringSliceFlag{
			Name:     config.FlagComponents,
			Aliases:  []string{"co"},
			Usage:    "List of components to run",
			Required: false,
			Value:    cli.NewStringSlice(AGGREGATOR, SEQUENCER, RPC, SYNCHRONIZER),
		},
		&cli.StringSliceFlag{
			Name:     config.FlagHTTPAPI,
			Aliases:  []string{"ha"},
			Usage:    fmt.Sprintf("List of JSON RPC apis to be exposed by the server: --http.api=%v,%v,%v,%v,%v,%v", jsonrpc.APIEth, jsonrpc.APINet, jsonrpc.APIDebug, jsonrpc.APIHez, jsonrpc.APITxPool, jsonrpc.APIWeb3),
			Required: false,
			Value:    cli.NewStringSlice(jsonrpc.APIEth, jsonrpc.APINet, jsonrpc.APIHez, jsonrpc.APITxPool, jsonrpc.APIWeb3),
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:    "version",
			Aliases: []string{},
			Usage:   "Application version and build",
			Action:  versionCmd,
		},
		{
			Name:    "run",
			Aliases: []string{},
			Usage:   "Run the hermez core",
			Action:  start,
			Flags:   flags,
		},
		{
			Name:    "approve",
			Aliases: []string{"ap"},
			Usage:   "Approve tokens to be spent by the smart contract",
			Action:  approveTokens,
			Flags: append(flags, &cli.StringFlag{
				Name:     config.FlagAmount,
				Aliases:  []string{"am"},
				Usage:    "Amount that is gonna be approved",
				Required: true,
			},
			),
		},
		{
			Name:    "encryptKey",
			Aliases: []string{},
			Usage:   "Encrypts the privatekey with a password and create a keystore file",
			Action:  encryptKey,
			Flags:   encryptKeyFlags,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}