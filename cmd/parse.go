package cmd

import (
	"bufio"
	"github.com/pzmi/rsstogo/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "parse",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		file, err := os.Open("top.rss")
		if err != nil {
			log.WithContext(ctx).WithError(err).Error("Could not open rss file")
			return err
		}
		defer file.Close()

		reader := bufio.NewReader(file)

		err = pkg.Parse(ctx, reader)
		if err != nil {
			log.WithContext(ctx).WithError(err).Error("Could parse rss")
			return err
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
