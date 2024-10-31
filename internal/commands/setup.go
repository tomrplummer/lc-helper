package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tomrplummer/lc-helper/internal/gpt"
)

func NewSetupMessage(description string, lang string, url string) []gpt.Message {

	return []gpt.Message{
		{
			Role:    "system",
			Content: "You are a leetcode helper.  Your job is to read the description and provide starter code, test cases and analysis for a given language.  You are an expert in " + lang,
		}, {
			Role: "user",
			Content: `
				Join the provided problem description, code snippet, and use cases into a single file. Format the output as JSON with three fields:

				*	filename: Use the function name as the filename, followed by the appropriate extension for the language.
				*	problem: Insert the problem description as a comment (with line breaks for readability), followed by the exact code snippet as provided, and finally, the example use cases as runnable code.
				*	lang: Specify the language in lowercase.

				Guidelines:

				*	Do not alter the code snippet or function name; both must match the exact content from LeetCode, as modifying them could cause errors when submitting to LeetCode.
				*	Convert the example use cases from the description into runnable code. Ensure that the file runs and executes these cases directly (e.g., in Go, place them in a main function).
				*	Do not solve the problem; simply structure the file as described.
				*	Do not use code snippets (wrapped in backticks), just plaintext responses, formatted appropriately
				
				\ndescription: ` + description + `\nlang: ` + lang + `\nurl: `,
		},
	}
}

func SaveSetupContent(response *gpt.Response) error {
	var setupData gpt.SetupResponseMessage
	json.Unmarshal([]byte(response.Choices[0].Message.Content), &setupData)

	if setupData.Problem == "" || setupData.Filename == "" {
		return fmt.Errorf("did not get a response, try a different language?")
	}

	path, err := createDirectory(setupData.Filename, setupData.Lang)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, []byte(setupData.Problem), 0777)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}

	fmt.Printf("Creating file: %v\n", path)

	return nil
}

func createDirectory(filename string, lang string) (string, error) {
	folderName := strings.Split(filename, ".")[0]

	leetcode_path := os.Getenv("LEETCODE_PATH")
	if leetcode_path == "" {
		return "", fmt.Errorf("LEETCODE_PATH not found in ENV.  Create LEETCODE_PATH env variable")
	}

	path := filepath.Join(os.Getenv("LEETCODE_PATH"), lang, folderName)

	if err := os.MkdirAll(path, 0777); err != nil {
		return "", fmt.Errorf("unable to create directory: %v", err)
	}

	fmt.Printf("Creating directory: %s\n", path)

	fullPath := filepath.Join(path, filename)

	return fullPath, nil
}
