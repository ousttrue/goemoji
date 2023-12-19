package cmd

import (
	"fmt"
	"os"

	"github.com/ousttrue/goemoji/unicode"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goemoji",
	Short: "list unicode emoji",
	Long: `list unicode emoji

U+XXXXX [glyph] MajorVersion.MinorVersion Group:Subgroup Name

cache: ~/.cache/goemoji/emoji-test.txt
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		list := unicode.Parse(unicode.GetCache())
		for _, emoji := range list {
			fmt.Printf("U+%5x %s %02d.%02d %s:%s %s\n",
				emoji.Unicode, string(emoji.Unicode),
				emoji.MajorVersion, emoji.MinorVersion,
				emoji.Group, emoji.Subgroup, emoji.Name)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goemoji.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
