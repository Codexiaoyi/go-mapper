package mapper

import "github.com/spf13/cobra"

var GenCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate mapper code.",
	Long:  "Generate mapper code at this layout.",
	Run:   Run,
}

// Run generate.
func Run(cmd *cobra.Command, args []string) {

}
