# LeetCode Helper (lc-helper)
###### README generated by gpt-4o
A command-line tool that helps streamline your LeetCode problem-solving workflow by automatically generating starter files and providing intelligent hints using OpenAI's GPT API.

## Features

- 🚀 Automatically generate starter code files from LeetCode problems
- 💡 Get progressive hints with 5 different levels of assistance
- 📝 Maintains problem descriptions, starter code, and test cases in a single organized file
- 🔄 Integrates directly with LeetCode's GraphQL API
- 🤖 Uses GPT for intelligent assistance and hints

## Prerequisites

- Go 1.23 or higher
- OpenAI API key

## Configuration

Set the following environment variables:

```bash
export LCHELPER_OPENAI_KEY="your-openai-api-key"
export LEETCODE_PATH="/path/to/leetcode/workspace"
```

## Usage

### Generate Starter Code

```bash
lc-helper setup <problem-slug> --lang=<language>
```

Example:
```bash
lc-helper setup two-sum --lang=python
```

This will:
1. Fetch the problem from LeetCode
2. Generate a starter file with:
   - Problem description as comments
   - Starter code template
   - Example test cases

### Get Hints

```bash
lc-helper hint <filename> --level <1-5>
```

Hint levels:
- Level 1: Very subtle hint about approach
- Level 2: General direction
- Level 3: More specific strategy
- Level 4: Detailed approach
- Level 5: Comprehensive solution strategy

Example:
```bash
lc-helper hint two-sum/solution.py --level 2
```

## Project Structure

```
lc-helper/
├── cli/
│   └── cli.go                 # Main CLI application
├── internal/
│   ├── commands/
│   │   ├── hint.go           # Hint command implementation
│   │   └── setup.go          # Setup command implementation
│   ├── gpt/
│   │   ├── call.go           # OpenAI API interaction
│   │   └── types.go          # Type definitions
│   └── lc/
│       └── lc.go             # LeetCode API interaction
├── go.mod
└── go.sum
```

## Development

### Building from Source

```bash
git clone https://github.com/tomrplummer/lc-helper.git
cd lc-helper
go build -o lc-helper ./cli
```

### Running Tests

```bash
go test ./...
```

## Dependencies

- [cobra](https://github.com/spf13/cobra) - CLI framework
- OpenAI GPT API - For generating hints and assistance

## Common Issues

1. **OpenAI API Key Not Found**
   ```
   panic: no LCHELPER_OPENAI_KEY found in ENV
   ```
   Solution: Make sure to set your OpenAI API key in environment variables

2. **LEETCODE_PATH Not Set**
   Solution: Set the LEETCODE_PATH environment variable to your desired workspace directory

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

[MIT License](LICENSE)

## Acknowledgments

- Powered by OpenAI's GPT API
- Built with [Cobra](https://github.com/spf13/cobra)
- Uses LeetCode's GraphQL API

## Author

[Your Name]

## Support

For support, please open an issue in the GitHub repository.
