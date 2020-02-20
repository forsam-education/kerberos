package cmd

import (
	"fmt"
	"github.com/forsam-education/cerberus/utils"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version for Cerberus.",
	Long: fmt.Sprintf(`%s

  A simple but powerful reverse proxy.

  Shows the version for Cerberus.`, utils.ASCIILogo),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Cerberus version %d.%d.%d - %s (Build time: %s)\n", utils.MajorVersion, utils.MinorVersion, utils.PatchVersion, utils.VersionHash, utils.BuildTime)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
