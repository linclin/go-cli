package cmd

import (
	"fmt"
	"go-cli/utils/log"
	"os"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "root",
	Short: "",
	Long: `
`,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		log.Init(logLevel)
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var logLevel string

func init() {
	RootCmd.PersistentFlags().StringVar(&logLevel, "logLevel", "info", "设置日志等级 (Set log level) [trace|debug|info|warn|error|fatal|panic]")
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:  RootCmd,
		Headings: cc.HiGreen + cc.Underline,
		Commands: cc.Cyan + cc.Bold,
		Example:  cc.Italic,
		ExecName: cc.Bold,
		Flags:    cc.Cyan + cc.Bold,
	})
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println("命令执行失败", err)
		os.Exit(1)
	}
}
