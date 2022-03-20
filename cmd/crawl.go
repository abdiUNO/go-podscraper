/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"podscraper/internals/services"
	"strings"
)

var items []string

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		genres := map[string]bool{}

		for _, v := range items {
			genre := strings.Title(v)
			log.Println(genre)
			genres[genre] = true
		}

		scraper := services.NewPodScraper(
			services.WithInitUrl("https://podcasts.apple.com/us/genre/podcasts/id26"),
			services.WithGenres(genres),
		)

		if err := scraper.Start(); err != nil {
			for _, v := range genres {
				log.Println(v)
			}
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//crawlCmd.PersistentFlags().String("foo", "", "A help for foo")

	crawlCmd.PersistentFlags().StringSliceVarP(&items, "genres", "g", []string{}, "")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crawlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
