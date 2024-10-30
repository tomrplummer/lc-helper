package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tomrplummer/lc-helper/internal/commands"
	"github.com/tomrplummer/lc-helper/internal/gpt"
	"github.com/tomrplummer/lc-helper/internal/lc"
)

func main() {
	apikey := os.Getenv("LCHELPER_OPENAI_KEY")

	if apikey == "" {
		panic("no LCHELPER_OPENAI_KEY found in ENV")
	}

	var rootCmd = &cobra.Command{
		Use:   "lc-helper",
		Short: "Generate leetcode starter file for given url",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	var lang string

	var setupCmd = &cobra.Command{
		Use:   "setup  <slug> --lang=<lang>",
		Short: "Creates file with use cases for given leetcode url",
		Long: `Creates file with use caes for given leetcode url.  
			   Requires LEETCODE_PATH to be set in env to know where to create the file.
			   `,
		Run: func(cmd *cobra.Command, args []string) {
			queryResult, err := lc.Query(args[0], lang)
			if err != nil {
				panic(err)
			}

			messages := commands.NewSetupMessage(queryResult.Content, lang, "")
			request := gpt.NewRequest(messages)

			resp, _ := gpt.CallApi(apikey, *request)

			if err := commands.SaveSetupContent(resp); err != nil {
				panic(err)
			}

		},
	}

	setupCmd.Flags().StringVar(&lang, "lang", "", "Language to use with file creation")

	var level string

	var hintCmd = &cobra.Command{
		Use:   "hint <filename> --level <int>",
		Short: "Get a hint on how to solve the problem.  --level 1 to 5 with 1 being a small hint and 5 being very, very helpful",
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]

			source, err := os.ReadFile(filename)
			if err != nil {
				panic(err)
			}

			messages := commands.NewHintMessage(string(source), level)
			request := gpt.NewRequest(messages)

			resp, _ := gpt.CallApi(apikey, *request)

			hint := resp.Choices[0].Message.Content

			fmt.Println(hint)

			if err = commands.StoreHint(filename, hint); err != nil {
				panic(err)
			}
		},
	}

	hintCmd.Flags().StringVar(&level, "level", "3", "Int, 1 through 5.  1 means a small hint, 5 means a big hint")

	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(hintCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
