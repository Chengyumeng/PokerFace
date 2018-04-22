package cmd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"

	"github.com/Chengyumeng/PokerFace/global"
	"github.com/Chengyumeng/PokerFace/initial"
	"github.com/Chengyumeng/PokerFace/models"
)

var (
	InitDBCmd = &cobra.Command{
		Use:     "initdb --config CONFIG",
		Short:   "初始化数据库",
		Example: `porkerface initdb --config CONFIG`,
		PreRunE: preRunE,
		RunE:    runE,
	}
)

func preRunE(cmd *cobra.Command, args []string) error {
	initial.Run(config)
	return nil
}
func runE(cmd *cobra.Command, args []string) error {
	err := global.Orm.Sync(
		models.Person{},
		models.User{},
		models.Comment{},
	)
	if err != nil {
		return err
	}
	fmt.Println("Sucess sync db")
	return nil
}
func init() {
	cobra.EnableCommandSorting = false
	InitDBCmd.Flags().SortFlags = false
	InitDBCmd.Flags().StringVarP(&config, "config", "c", "porkerface.toml", "配置文件")
}
