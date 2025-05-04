package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "goyamlc",
		RunE: func(cmd *cobra.Command, args []string) error {
			typesPath, err := cmd.PersistentFlags().GetString("types-path")
			if err != nil {
				return err
			}
			if typesPath == "" {
				fmt.Fprintln(os.Stderr, errors.New(`"--types-path" must be set`))
				os.Exit(1)
			}

			genPath, err := cmd.PersistentFlags().GetString("gen-path")
			if err != nil {
				return err
			}
			if typesPath == "" {
				fmt.Fprintln(os.Stderr, errors.New(`"--gen-path" must be set`))
				os.Exit(1)
			}

			if err := configureLogger("info"); err != nil {
				return err
			}

			if err := Generate(typesPath, genPath, nil); err != nil {
				return err
			}
			return nil
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func init() {
	_ = rootCmd.PersistentFlags().String("types-path", "", "location of your go file")
	_ = rootCmd.PersistentFlags().String("gen-path", "", "location to generate the yaml file")
}
