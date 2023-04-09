package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	host string
	port int
)

var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "A simple quiz application",
	Long:  `A CLI application that allows users to take a simple quiz and view their results.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "API server hostname or IP address")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8080, "API server port number")
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println("Help message goes here")
	})
}
