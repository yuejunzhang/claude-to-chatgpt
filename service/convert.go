package service

import (
	"encoding/json"
	"fmt"

	"github.com/oldweipro/claude-to-chatgpt/model"
)

// OpenaiToClaudeParams 转换成claude的参数
func OpenaiToClaudeParams(chatCompletionRequest model.ChatCompletionRequest) *model.ChatMessageRequest {
	history := chatCompletionRequest.Messages
	text := chatCompletionRequest.Messages[len(history)-1].Content
	var promptMessage = ""
	textMarshal, err := json.Marshal(text)
	if err != nil {
		fmt.Println("Text marshal err:", err)
	}
	textMessage := string(textMarshal)

	if chatCompletionRequest.Messages[0].Role == "system" {
		prompt := chatCompletionRequest.Messages[0].Content
		promptMarshal, err := json.Marshal(prompt)
		if err != nil {
			fmt.Println("Text marshal err:", err)
		}
		promptMessage = string(promptMarshal)
	} else {
		promptMessage = textMessage
	}

	historyMessage := ""
	if len(history) > 0 {
		historyMarshal, err := json.Marshal(history)
		if err != nil {
			fmt.Println("History marshal err:", err)
		}
		historyMessage = string(historyMarshal)
	}
	return model.NewChatMessageRequest(promptMessage, textMessage, historyMessage)
}

// func OpenaiToClaudeParams(chatCompletionRequest model.ChatCompletionRequest) *model.ChatMessageRequest {
// 	completionMessages := chatCompletionRequest.Messages
// 	text := completionMessages[len(completionMessages)-1]
// 	history := completionMessages[:len(completionMessages)-1]
// 	textMarshal, err := json.Marshal(text)
// 	if err != nil {
// 		fmt.Println("Text marshal err:", err)
// 	}
// 	textMessage := string(textMarshal)
// 	historyMessage := ""
// 	if len(history) > 0 {
// 		historyMarshal, err := json.Marshal(history)
// 		if err != nil {
// 			fmt.Println("History marshal err:", err)
// 		}
// 		historyMessage = string(historyMarshal)
// 	}
// 	return model.NewChatMessageRequest(textMessage, historyMessage)
// }
