package commands

import (
	"fmt"
	"os"

	"github.com/tomrplummer/lc-helper/internal/gpt"
)

func NewHintMessage(source, level, responseStyle string) []gpt.Message {

	return []gpt.Message{
		{
			Role:    "system",
			Content: "You are a leetcode helper.  You will analyze a source code file that has a description, code and previous hints and provide assistance.  You are an expert in the lange the code is written in",
		}, {
			Role: "user",
			Content: `
			Analyze the provided problem description, previous hints (marked as [level]), code snippet, and desired response style. Generate an additional hint according to the given level (1-5), where:

			*	Level 1 provides minimal guidance (e.g., a useful data structure).
			*	Level 5 outlines the complete strategy. For repeated Level 5 requests, provide progressively detailed hints. Hints must be concise, a single sentence without lists or line breaks.
			*	Do not use code snippets (wrapped in backticks), just plaintext responses, formatted appropriately

			Format the hint as a comment in the relevant language, prefixed by the level (e.g., # [1] Hint text for Ruby). If level is solve_it, provide the complete solution as code, including line breaks.

			Match the requested response style, which may range from helpful to funny or others as specified. Default style is helpful.			\nsource: ` + source + `\nlevel: ` + level + `\nresponse style: ` + responseStyle + `\n`},
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
