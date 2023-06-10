package version

import (
	"go-cli/cmd"
	"go-cli/utils/Print"

	"github.com/spf13/cobra"
)

var version = "1.0"

// version module

// init in modules will add self to RootCmd when init package.
func init() {
	cmd.RootCmd.AddCommand(VersionCmd)
}

// VersionCmd is core cobra.Command of subcommand
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "输出版本 (Print the version number)",
	Long:  "输出版本 (Print the version number)",
	Run: func(cmd *cobra.Command, args []string) {
		data := [][]string{
			{version},
		}
		var header = []string{"当前版本 (Version)"}
		var td = Print.Table{
			Header: header,
			Body:   data}
		td.Print("")
	},
}
