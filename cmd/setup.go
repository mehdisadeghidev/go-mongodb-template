package cmd

import (
	"bufio"
	"fmt"
	"github.com/mehdisadeghidev/go-mongodb-template/config"
	"github.com/mehdisadeghidev/go-mongodb-template/database"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

const banner = `
 _     ___  _  _____ 
| |   / _ \| |/ /_ _|
| |  | | | | ' / | | 
| |__| |_| | . \ | | 
|_____\___/|_|\_\___|
`

func init() {
	config.Setup()

	err := os.Setenv("TZ", "UTC")
	if err != nil {
		log.Fatal(err)
	}

	err = database.Setup()
	if err != nil {
		log.Fatal(err)
	}
}

func ask(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, _ := reader.ReadString('\n')
		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

func Setup() {
	var loki = &cobra.Command{
		Use:     "loki",
		Version: "0.0.1",
		Short:   "Change to whatever you want",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.Println(strings.TrimPrefix(banner, "\n"))
		},
	}

	loki.AddCommand(&cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		Run: func(cmd *cobra.Command, args []string) {
			RunServer()
		},
	})

	loki.AddCommand(&cobra.Command{
		Use:   "db:seed",
		Short: "Seeds the database with test data",
		PreRun: func(cmd *cobra.Command, args []string) {
			if !ask("Are you sure?") {
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			RunSeeder()
		},
	})

	loki.AddCommand(&cobra.Command{
		Use:   "db:wipe",
		Short: "Wipes the database",
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.Println("This will wipe the database and all data.")

			if !ask("Are you sure?") {
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			err := database.DB.Drop(cmd.Context())

			if err != nil {
				log.Fatal(err)
			}
		},
	})

	if err := loki.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
