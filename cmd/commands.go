package cmd

import (
	"github.com/spf13/cobra"
)

// CommandRegistry 命令注册器
type CommandRegistry struct {
	commands []*cobra.Command
}

// registry 全局命令注册器实例
var registry = &CommandRegistry{}

// RegisterCommand 注册命令
func RegisterCommand(cmd *cobra.Command) {
	registry.commands = append(registry.commands, cmd)
}

// RegisterCommands 批量注册命令
func RegisterCommands(cmds ...*cobra.Command) {
	registry.commands = append(registry.commands, cmds...)
}

// GetAllCommands 获取所有注册的命令
func (r *CommandRegistry) GetAllCommands() []*cobra.Command {
	return r.commands
}

// RegisterAllToRoot 将所有注册的命令添加到根命令
func (r *CommandRegistry) RegisterAllToRoot() {
	for _, cmd := range r.commands {
		RootCmd.AddCommand(cmd)
	}
}
