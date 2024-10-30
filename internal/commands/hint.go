package commands

import (
	"fmt"
	"os"

	"github.com/tomrplummer/lc-helper/internal/gpt"
)

func NewHintMessage(source string, level string) []gpt.Message {

	return []gpt.Message{
		{
			Role:    "system",
			Content: "You are a leetcode helper.  You will analyze a source code file that has a description, code and previous hints and provide assistance.  You are an expert in the lange the code is written in",
		}, {
			Role: "user",
			Content: `
			Provided: Description, previous hints and code snippet.
			Analyze the problem, code and previous hints.  provide additional hints, without providing code.  A "level" will also be provided.  it will be 1 throught 5.  1
			means provide a little guidance.  maybe a datastructure that could be useful.  5 means provide the full strategy.  if they've requested 5 more than once, keep providing more
			assistance. Previous hints will be in the source provided start with [#] with the # representing the hint level requested before.
			Your hints should only be a few sentences at most (no lists, numbering or line breaks).  If more help is need, you will get another request.  
			The result should be formatted as a comment for the given langauage, but not a code snippet (backticks and the language name would break the file. plain text only) that would break the source file.  Include the level at the beginning of the line like the following example
			(example for ruby: # [1] <your hint here for a level 1 comment>)
			if level is solve_it, ignore previous restrictions and provide the code that solves the problem.  you must provide code, include line breaks....do what you need to do to get the problem solved.
			\nsource: ` + source + `\nlevel: ` + level + `\n`,
		},
	}
}

func StoreHint(filename string, hint string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	newHint := fmt.Sprintf("\n\n%s\n", hint)

	if _, err = file.WriteString(newHint); err != nil {
		return fmt.Errorf("unable to write file: %v", err)
	}

	return nil
}
