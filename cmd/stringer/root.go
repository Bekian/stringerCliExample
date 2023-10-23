package stringer

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "./stringerExample",
	Short: "stringer - a simple CLI to transform and inspect strings",
	Long: `stringer is a super fancy CLI (kidding)
   
	One can use stringer to modify or inspect strings straight from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing your cli '%s', ", err)
		os.Exit(1)
	}
}
