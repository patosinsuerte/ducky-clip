package cmd

import (
	"fmt"
	"os"

	"github.com/patosinsuerte/ducky-clip/internal/database"
	"github.com/spf13/cobra"
)

var version = "v0.1.0-beta"

const duckyAscii = `
     _          
   >(o)____      [ Ducky-Clip ]
    (___  /      `

const customHelpTemplate = `{{Yellow}}{{DuckyAscii}}{{Reset}}{{Yellow}}[ {{.Version}} ]{{Reset}}
{{Blue}}~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~{{Reset}}

{{Yellow}}Ducky-clip{{Reset}} {{.Version}} — Pond Command Guide

{{if .Long}}{{.Long}}{{end}}

{{Yellow}}Usage:{{Reset}}
  {{.UseLine}}

{{if .Aliases}}{{Yellow}}Aliases:{{Reset}}
  {{.NameAndAliases}}{{end}}

{{if .HasAvailableSubCommands}}{{Yellow}}Available Commands:{{Reset}}
{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}  {{rpad .Name .NamePadding }} {{.Short}}{{end}}
{{end}}{{end}}
{{if .HasAvailableLocalFlags}}{{Yellow}}Flags:{{Reset}}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}

{{if .Example}}{{Yellow}}Examples:{{Reset}}
{{.Example}}{{end}}

Use "{{.CommandPath}} [command] --help" for more information about a command.
`

var RootCMD = &cobra.Command{
	Use:     "ducky",
	Short:   "A powerful CLI snippet manager for developers",
	Long:    `Ducky-clip is a cross-platform CLI tool designed to save your personal library for code snippets and commands.`,
	Version: version,
}

func init() {
	cobra.OnInitialize(func() {
		if err := database.InitDB(); err != nil {
			fmt.Printf("\033[31mFatal error:\033[0m Could not connect to the database: %v\n", err)
			os.Exit(1)
		}
	})

	RootCMD.CompletionOptions.DisableDefaultCmd = true

	cobra.AddTemplateFunc("Yellow", func() string { return "\033[33m" })
	cobra.AddTemplateFunc("Blue", func() string { return "\033[34m" })
	cobra.AddTemplateFunc("Reset", func() string { return "\033[0m" })
	cobra.AddTemplateFunc("DuckyAscii", func() string { return duckyAscii })

	RootCMD.SetHelpTemplate(customHelpTemplate)

	RootCMD.SetVersionTemplate("🦆 {{.Name}} version: {{.Version}}\n")
}
