package main

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/Chengyumeng/PokerFace/bin/cmd"
)

var RootCmd = &cobra.Command{
	Use: "porkerface",
	Long: `PorkerFace 平台
反馈及建议:https://github.com/Chengyumeng/PorkerFace
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	RootCmd.AddCommand(cmd.InitDBCmd)
	RootCmd.AddCommand(cmd.ServerCmd)
}
func main() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
