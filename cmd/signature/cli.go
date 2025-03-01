package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/version"
	"github.com/urfave/cli/v2" // https://cli.urfave.org/v2/getting-started/

	"github.com/qiaopengjun5162/web3-wallet-sign/common/cliapp"
	"github.com/qiaopengjun5162/web3-wallet-sign/config"
	flags2 "github.com/qiaopengjun5162/web3-wallet-sign/flags"
	"github.com/qiaopengjun5162/web3-wallet-sign/leveldb"
	"github.com/qiaopengjun5162/web3-wallet-sign/services/rpc"
)

// Semantic holds the textual version string for major.minor.patch.
var Semantic = fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)

// WithMeta holds the textual version string including the metadata.
var WithMeta = func() string {
	v := Semantic
	if version.Meta != "" {
		v += "-" + version.Meta
	}
	return v
}()

func withCommit(gitCommit, gitDate string) string {
	vsn := WithMeta
	if len(gitCommit) >= 8 {
		vsn += "-" + gitCommit[:8]
	}
	if (version.Meta != "stable") && (gitDate != "") {
		vsn += "-" + gitDate
	}
	return vsn
}

func runRpc(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	fmt.Println("running grpc services...")
	cfg := config.NewConfig(ctx)
	grpcServerCfg := &rpc.RpcServerConfig{
		GrpcHostname: cfg.RPCServer.Host,
		GrpcPort:     cfg.RPCServer.Port,
		KeyName:      cfg.KeyName,
		KeyPath:      cfg.CredentialsFile,
		HsmEnable:    cfg.HsmEnable,
	}
	db, err := leveldb.NewKeyStore(cfg.LevelDbPath)
	if err != nil {
		log.Error("new key store level db", "err", err)
	}
	return rpc.NewRpcServer(db, grpcServerCfg)
}

func NewCli(GitCommit string, gitDate string) *cli.App {
	flags := flags2.Flags
	return &cli.App{
		Name:                 "Web3 Wallet Sign",
		Usage:                "An exchange wallet scanner services with rpc and rest api services",
		Description:          "An exchange wallet scanner services with rpc and rest api services",
		Version:              withCommit(GitCommit, gitDate),
		EnableBashCompletion: true, // Boolean to enable bash completion commands
		Commands: []*cli.Command{
			{
				Name:        "rpc",
				Flags:       flags,
				Description: "Run rpc services",
				Action:      cliapp.LifecycleCmd(runRpc),
			},
			{
				Name:        "version",
				Usage:       "Show project version",
				Description: "Show project version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
