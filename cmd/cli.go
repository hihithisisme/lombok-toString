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

func parseAsJSON(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Printf("No input string passed\n\n")
		cmd.Help()
	} else {
		lString := args[0]
		fmt.Println(lString)
		fmt.Println(lombokString.New(lString).ParseAsJSON())
	}
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
