<h1 align="center">go-cli</h1>

<div align="center">
Cobra + Viper + Zap实现的golang命令行模板
<p align="center">
<img src="https://img.shields.io/badge/Golang-1.20.2-brightgreen" alt="Go version"/> 
</p>
</div>

## 主要功能
- 基于`Cobra`框架开发
- 配置文件管理使用`Viper`，配置文件值可被flag覆盖
- 日志输出使用`slog` 
- 命令行输出表格结构体`TableWriter`   

## 使用开源类库
- [Cobra](https://github.com/spf13/cobra) golang命令行框架 [教程](https://juejin.cn/post/6924541628031959047) 
- [Viper](https://github.com/spf13/viper)  配置管理工具, 支持多种配置文件类型.[教程](https://darjun.github.io/2020/01/18/godailylib/viper/)  
- [TableWriter](github.com/olekukonko/tablewriter)  渲染输出表格结构体
- [ColorCobra](https://github.com/ivanpirog/coloredcobra) 彩色命令行输出  
 