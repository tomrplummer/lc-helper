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
			Provided: Description and code snippet.
			We'll be joining everything into a single file, so the description should be a comment, followed by the code snippet followed by the provided use cases in the description.
			Don't return it as a code block because it breaks the source file.
			The output should be json in the following format.
			{
				"filename": //use the function name followed by the correct extension,
				"problem": //this is where you put the problem.  it will be written to a file, make sure it doesn't include anything that would break the code (like being a snippet).  also include the 'start code' for the given language,
				"lang": //language name, lowercase
			}
			Make sure to convert the examples in to runnable use cases below the starter code.  They should be the exact examples found in the description.
			Format the description (line breaks) to make it more readable, but still as a comment.
			If a language requires additional code to be runnable (for example, package main in Go, or a main function in many languages, or potentially inside a module), add that as well.
			The example use cases must also runnable by just running the file.  (in Go, they could be in the main function for example)
			Use the url to get the starter code if needed.
			DO NOT SOLVE THE LEETCODE PROBLEM.
			The name of the function in the starter code must not change from what was provided.  This will cause errors when submitting to leetcode
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

	path := filepath.Join(os.Getenv("LEETCODE_PATH"), lang, folderName)

	if err := os.MkdirAll(path, 0777); err != nil {
		return "", fmt.Errorf("unable to create directory: %v", err)
	}

	fmt.Printf("Creating directory: %s\n", path)

	fullPath := filepath.Join(path, filename)

	return fullPath, nil
}
