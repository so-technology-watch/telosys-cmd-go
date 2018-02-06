package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var availableBundles []string

// nmCmd represents the nm command
var bCmd = &cobra.Command{
	Use:     "b",
	Aliases: []string{"bundle"},
	Short:   "Set current bundle",
	Long:    "Set the bundle used in the telosys project",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			selectBundle()
		} else {
			setBundle(args[0])
		}
	},
}

func init() {
	availableBundles = getBundles()
	rootCmd.AddCommand(bCmd)
}

func setBundle(name string) {
	if isUnique, bundleName := isUniquePossibility(name, availableBundles); isUnique {
		setConfValue(cfgBundle, bundleName)
		fmt.Println("Bundle successfully set to", bundleName)
	} else {
		fmt.Println("Bundle doesn't exist")
	}
}

func getBundles() []string {
	bundles := getMatching("templates/*")
	newList := []string{}
	for _, bundle := range bundles {
		newList = append(newList, rmPath(bundle))
	}
	return newList
}

func selectBundle() {
	fmt.Println("Here are the available bundles:")
	listSelector(availableBundles, setBundle, func() {
		fmt.Println("You didn't pick a correct bundle, please retry")
		selectBundle()
	})
}
