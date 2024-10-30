package gpt

type Request struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float32   `json:"temperature"`
	//MaxTokens        int       `json:"max_tokens"`
	//TopP             float32   `json:"top_p"`
	//FrequencyPenalty float32   `json:"frequency_penalty"`
	//PresencePenalty  float32   `json:"presence_penalty"`
}
type Response struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type SetupResponseMessage struct {
	Filename string `json:"filename"`
	Problem  string `json:"problem"`
	Lang     string `json:"lang"`
}

type GraphQLResponse struct {
	Data struct {
		Question struct {
			QuestionFrontendId string `json:"questionFrontendId"`
			TitleSlug          string `json:"titleSlug"`
			Content            string `json:"content"`
			CodeSnippets       []struct {
				Lang     string `json:"lang"`
				LangSlug string `json:"langSlug"`
				Code     string `json:"code"`
			} `json:"codeSnippets"`
		} `json:"question"`
	} `json:"data"`
}

type QueryResult struct {
	Content string
	Code    string
}

func NewRequest(messages []Message) *Request {
	return &Request{
		Model:       "gpt-4o-mini",
		Messages:    messages,
		Temperature: 0.5,
	}
}
