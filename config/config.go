// Package config 包含了应用程序的配置信息。
package config

import (
	"github.com/urfave/cli/v2"

	"github.com/qiaopengjun5162/web3-wallet-sign/flags"
)

// 定义一个ServerConfig结构体，包含Host和Port两个字段
type ServerConfig struct {
	Host string
	Port int
}

// 定义一个Config结构体，用于存储配置信息
type Config struct {
	// LevelDB数据库的路径
	LevelDbPath string
	// RPC服务器的配置信息
	RPCServer ServerConfig
	// 凭证文件的路径
	CredentialsFile string
	// 密钥的名称
	KeyName string
	// 是否启用HSM
	HsmEnable bool
}

// NewConfig 根据 CLI 上下文创建并返回一个新的配置实例。
// 该函数从 cli.Context 中提取配置参数，包括数据库路径、凭证文件、密钥名称、
// 硬件安全模块启用状态以及 RPC 服务器的主机和端口信息，然后将这些参数
// 用于初始化 Config 结构体。
func NewConfig(ctx *cli.Context) Config {
	return Config{
		// 从上下文中获取 LevelDb 路径
		LevelDbPath: ctx.String(flags.LevelDbPathFlag.Name),
		// 从上下文中获取凭证文件路径
		CredentialsFile: ctx.String(flags.CredentialsFileFlag.Name),
		// 从上下文中获取密钥名称
		KeyName: ctx.String(flags.KeyNameFlag.Name),
		// 从上下文中获取硬件安全模块启用状态
		HsmEnable: ctx.Bool(flags.HsmEnable.Name),
		// 初始化 RpcServer 配置
		RPCServer: ServerConfig{
			// 从上下文中获取 RPC 服务器主机名
			Host: ctx.String(flags.RpcHostFlag.Name),
			// 从上下文中获取 RPC 服务器端口号
			Port: ctx.Int(flags.RpcPortFlag.Name),
		},
	}
}
