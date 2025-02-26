package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/version"
	"github.com/urfave/cli/v2" // https://cli.urfave.org/v2/getting-started/
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

func NewCli(GitCommit string, gitDate string) *cli.App {
	return &cli.App{
		Name:                 "Web3 Wallet Sign",
		Usage:                "An exchange wallet scanner services with rpc and rest api services",
		Description:          "An exchange wallet scanner services with rpc and rest api services",
		Version:              withCommit(GitCommit, gitDate),
		EnableBashCompletion: true, // Boolean to enable bash completion commands
		Commands: []*cli.Command{
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
