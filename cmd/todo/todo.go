package todo

import (
	"fmt"
	"go-cli/cmd"
	"os"

	"github.com/go-jarvis/cobrautils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type TodoParam struct {
	Name string `mapstructure:"name" flag:"name" usage:"name" persistent:"true"`
}

var configFile = "./todo.conf"
var todoParams TodoParam
var TodoCmd = &cobra.Command{
	Use:   "todo",
	Short: "",
	Long:  "",
	PreRun: func(cmd *cobra.Command, args []string) {
		initTodoConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(todoParams.Name)
	},
}

func init() {
	TodoCmd.Flags().StringVarP(&configFile, "config", "", "./todo.conf", "配置文件")
	cobrautils.BindFlags(TodoCmd, &todoParams)
	TodoCmd.MarkFlagRequired("config")
	TodoCmd.Flags().VisitAll(func(f *pflag.Flag) {
		viper.BindPFlag(f.Name, TodoCmd.Flags().Lookup(f.Name))
	})
	cmd.RootCmd.AddCommand(TodoCmd)
}

func initTodoConfig() {
	if configFile != "" {
		viper.SetConfigType("toml")
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("配置文件读取失败", err)
			os.Exit(1)
		}
		if err := viper.Unmarshal(&todoParams); err != nil {
			fmt.Println("配置文件解析失败", err)
			os.Exit(1)
		}
	}
}
