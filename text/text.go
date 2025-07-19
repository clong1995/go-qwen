package text

import (
	"encoding/json"
	"github.com/clong1995/go-config"
	"go-qwen/send"
	"log"
)

func Send(msg []Message) (content string, err error) {
	data, err := json.Marshal(request{
		Model:    config.Value("QWEN_MODEL"),
		Messages: msg,
	})
	if err != nil {
		log.Println(err)
		return
	}

	bytes, err := send.Send("https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions", data, map[string]string{})
	if err != nil {
		log.Println(err)
		return
	}

	res := new(response)
	if err = json.Unmarshal(bytes, &res); err != nil {
		log.Println(err)
		return
	}
	for _, v := range res.Choices {
		content += v.Message.Content
	}
	return
}

type response struct {
	Choices []choice `json:"choices"`
}

type choice struct {
	Message message `json:"message"`
}

type message struct {
	Content string `json:"content"`
}

type request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

//固定顺序：system → user → assistant → user → ...

// SystemMsg
// 设定模型的全局行为准则、专业领域或风格（如“你是一个医疗助手”）
// 必须作为messages数组的第一条消息出现，且通常只出现一次
// 不能在对话中间插入system消息
func SystemMsg(text string) Message {
	return Message{"system", text}
}

// UserMsg
// 代表用户的提问或指令，引导对话方向。
// 对话的最后一条消息必须是user
func UserMsg(text string) Message {
	return Message{"user", text}
}

// AssistantMsg
// 存储模型的历史回复，用于维持上下文连贯性。
func AssistantMsg(text string) Message {
	return Message{"assistant", text}
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
