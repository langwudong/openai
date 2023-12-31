package openai

const (
	// 聊天API
	OPENAI_CHAT_API = "https://one.aiskt.com/v1/chat/completions"
	// 画图API
	OPENAI_DARW_API = "https://api.openai.com/v1/images/generations"
	// 情感分析API
	OPENAI_EMATION_API = "https://api.openai.com/v1/moderations"
)

type OpenAI struct {
	Organization string
}
