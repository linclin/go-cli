package version

import (
	"go-cli/cmd"
	"go-cli/utils/print"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var version = "1.0"

// version module

// init in modules will register self to command registry when init package.
func init() {
	// 使用注册器模式注册命令
	cmd.RegisterCommand(VersionCmd)
}

// VersionCmd is core cobra.Command of subcommand
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "输出版本 (Print the version number)",
	Long:  "输出版本 (Print the version number)",
	Run: func(cmd *cobra.Command, args []string) {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			data := [][]string{{"无法读取构建信息"}}
			var header = []string{"版本信息"}
			var td = print.Table{
				Header: header,
				Body:   data,
			}
			td.Print("")
			return
		}

		var buildInfoData [][]string
		buildInfoData = append(buildInfoData, []string{"Golang构建版本", info.GoVersion})

		// 查找git信息
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				buildInfoData = append(buildInfoData, []string{"Git提交", setting.Value})
			case "vcs.time":
				buildInfoData = append(buildInfoData, []string{"构建时间", setting.Value})
			case "vcs.modified":
				modified := "否"
				if setting.Value == "true" {
					modified = "是"
				}
				buildInfoData = append(buildInfoData, []string{"已修改", modified})
			case "GOOS":
				buildInfoData = append(buildInfoData, []string{"系统", setting.Value})
			case "GOARCH":
				buildInfoData = append(buildInfoData, []string{"平台", setting.Value})
			case "CGO_ENABLED":
				buildInfoData = append(buildInfoData, []string{"CGO_ENABLED", setting.Value})
			case "GOAMD64":
				buildInfoData = append(buildInfoData, []string{"GOAMD64", setting.Value})
			case "vcs":
				buildInfoData = append(buildInfoData, []string{"版本控制", setting.Value})
			}
		}

		// 如果没有VCS信息，提示用户需要使用-buildvcs=true构建
		hasVCSInfo := false
		for _, item := range buildInfoData {
			if item[0] == "Git提交" || item[0] == "构建时间" {
				hasVCSInfo = true
				break
			}
		}
		if !hasVCSInfo {
			buildInfoData = append(buildInfoData, []string{"提示", "使用 -buildvcs=true 构建以获取VCS信息"})
		}

		var header = []string{"属性", "值"}
		var td = print.Table{
			Header: header,
			Body:   buildInfoData,
		}
		td.Print("")
	},
}
