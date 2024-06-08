package goto7z

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "goto7z",
		Short: `A simple automatic decompression 7z program`,
		Run: func(_ *cobra.Command, _ []string) {

		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func main() {
	err := Execute()
	if err != nil {
		panic(err)
	}
}
