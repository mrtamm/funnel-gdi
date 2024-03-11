package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completion code",
}

var bash = &cobra.Command{
	Use:   "bash",
	Short: "Generate bash completion code",
	Long: `This command generates bash CLI completion code.
Add "source <(funnel completion bash)" to your bash profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := RootCmd.GenBashCompletion(os.Stdout); err != nil {
			fmt.Printf("Detected error while running GenBashCompletion "+
				"command: %s\n", err)
		}
	},
}

func init() {
	completionCmd.AddCommand(bash)
}
