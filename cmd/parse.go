package cmd

import (
	"bufio"
	"github.com/pzmi/rsstogo/internal/config"
	"github.com/pzmi/rsstogo/pkg"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "parse",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		err := config.InitLogging(ctx)
		if err != nil {
			return err
		}
		appConfig, err := config.InitApplicationConfiguration(ctx)
		if err != nil {
			return err
		}

		for _, feed := range appConfig.Feeds {
			resp, err := http.Get(feed.Address)
			defer resp.Body.Close()
			if err != nil {
				return err
			}
			body := bufio.NewReader(resp.Body)
			err = pkg.Parse(body)
			if err != nil {
				return err
			}
		}

		//file, err := os.Open("top.rss")
		//if err != nil {
		//	log.WithContext(ctx).WithError(err).Error("could not open rss file")
		//	return err
		//}
		//defer file.Close()
		//
		//reader := bufio.NewReader(file)
		//
		//err = pkg.Parse(reader)
		//if err != nil {
		//	log.WithContext(ctx).WithError(err).Error("could parse rss")
		//	return err
		//}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
