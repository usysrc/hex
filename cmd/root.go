package cmd

import (
	"os"

	"github.com/spf13/cobra"
	view "github.com/usysrc/hex/internal/viewer"
)

var width int

func run(cmd *cobra.Command, args []string) {
	if len(args) < 1 || args[0] == "" {
		cmd.Usage()
		os.Exit(1)
		return
	}
	view.CreateView(args[0], width)
}

var rootCmd = &cobra.Command{
	Use:   "hex <filename>",
	Short: "View a file in hex format in a TUI.",
	Long:  `This is a simple hex viewer that displays a file in hex format in a TUI.`,
	Run:   run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func init() {
	// The number of bytes per line
	rootCmd.Flags().IntVar(&width, "width", 16, "bytes per line")
}
