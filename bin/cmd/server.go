package cmd

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"

	"github.com/Chengyumeng/PokerFace/global"
	"github.com/Chengyumeng/PokerFace/initial"
)

var (
	ServerCmd = &cobra.Command{
		Use:     "server --config CONFIG",
		Short:   "启动web服务",
		Example: `porkerface server --config CONFIG`,
		PreRunE: preRunServerE,
		RunE:    runServerE,
	}
)

func preRunServerE(cmd *cobra.Command, args []string) error {
	initial.Run(config)
	return nil
}
func runServerE(cmd *cobra.Command, args []string) error {
	addr := global.Configuration.AppConfig.Address

	err := global.Echo.Start(addr)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	cobra.EnableCommandSorting = false
	ServerCmd.Flags().SortFlags = false
	ServerCmd.Flags().StringVarP(&config, "config", "c", "porkerface.toml", "配置文件")
}
