package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"lombok-toString/pkg/lombokString"
)

var rootCmd = &cobra.Command{
	Use:   "lombokString [string to be parsed]",
	Short: "Parse your lombok toString output",
	Long: `This command helps to parse the default format of the Lombok toString into JSON.
This should be primarily used only for reading logs and is definitely not recommended to be used within in production.`,
	Run: parseAsJSON,
}

var shouldExcludeNulls bool
var shouldMinify bool

func init() {
	rootCmd.Flags().BoolVarP(&shouldExcludeNulls, "exclude-null", "x", false, "exclude the fields with null value")
	rootCmd.Flags().BoolVarP(&shouldMinify, "mini", "m", false, "minify output (i.e. remove all indents)")
}

func parseAsJSON(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Printf("No input string passed\n\n")
		cmd.Help()
	} else {
		lString := args[0]
		fmt.Println()

		iArgs := lombokString.InterfaceArgs{
			ShouldExcludeNulls: shouldExcludeNulls,
			ShouldMinify:       shouldMinify,
		}
		fmt.Println(lombokString.New(lString).ParseAsJSON(iArgs))
	}
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
