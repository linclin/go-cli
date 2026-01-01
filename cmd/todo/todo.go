package todo

import (
	"fmt"
	"go-cli/cmd"
	"log/slog"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type TodoParam struct {
	Name string `mapstructure:"name"`
}

var configFile = "./todo.conf"
var todoConfig TodoParam
var TodoCmd = &cobra.Command{
	Use:   "todo",
	Short: "",
	Long:  "",
	PreRun: func(cmd *cobra.Command, args []string) {
		initTodoConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo name:", todoConfig.Name)
		color.Blue("todo name: %s  ", todoConfig.Name)
		slog.Info("todo name:", "name", todoConfig.Name)
	},
}

func init() {
	TodoCmd.Flags().StringVarP(&configFile, "config", "", "./todo.conf", "配置文件")
	TodoCmd.Flags().StringVar(&todoConfig.Name, "name", "", "name")
	TodoCmd.MarkFlagRequired("config")
	TodoCmd.Flags().VisitAll(func(f *pflag.Flag) {
		viper.BindPFlag(f.Name, TodoCmd.Flags().Lookup(f.Name))
	})
	// 使用注册器模式注册命令
	cmd.RegisterCommand(TodoCmd)
}

func initTodoConfig() {
	if configFile != "" {
		viper.SetConfigType("toml")
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("配置文件读取失败", err)
			os.Exit(1)
		}
	}
	// 使用Viper的自动绑定，命令行参数会自动覆盖配置文件的值
	if err := viper.Unmarshal(&todoConfig); err != nil {
		fmt.Println("配置解析失败", err)
		os.Exit(1)
	}
}
