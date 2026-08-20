package text

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	conf "github.com/clong1995/go-qwen/config"
	"github.com/clong1995/go-qwen/tool"
	"github.com/pkg/errors"
)

//参考文档 https://help.aliyun.com/zh/model-studio/qwen-api-via-openai-responses

var httpClient = &http.Client{}

// Generate 发送一条消息，返回助手的回复和本次的 response_id
func Generate(userInput, instructions, previousResponseID string, tools tool.Tool) (reply string, newResponseID string, err error) {

	url := "https://dashscope.aliyuncs.com/compatible-mode/v1/responses"

	// 构造请求体
	reqBody := request{
		Model:              conf.TextModel,
		Input:              userInput,
		Instructions:       instructions,
		PreviousResponseID: previousResponseID, // 首轮传空字符串
		EnableThinking:     true,
		Tools:              tools,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", "", errors.WithStack(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", "", errors.WithStack(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+conf.ApiKey)

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", "", errors.WithStack(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", errors.WithStack(err)
	}

	// 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		return "", "", errors.Errorf("%d:%s", resp.StatusCode, string(body))
	}

	// 解析响应
	var chatResp response
	if err = json.Unmarshal(body, &chatResp); err != nil {
		return "", "", errors.WithStack(err)
	}

	// 提取回复文本
	for _, out := range chatResp.Output {
		if out.Type == "message" {
			if len(out.Content) > 0 && out.Content[0].Text != "" {
				return out.Content[0].Text, chatResp.ID, nil
			}
		}
	}

	return reply, newResponseID, errors.Errorf("%s", string(body))
}

// ChatRequest 定义 Responses API 的请求体
type request struct {
	Model              string    `json:"model"`                          // 模型名称，如 qwen3.7-plus
	Input              string    `json:"input"`                          // 当前用户的输入
	Instructions       string    `json:"instructions"`                   // 系统指令（顶层字段）
	PreviousResponseID string    `json:"previous_response_id,omitempty"` // 上一轮的 response_id（首轮为空）
	Tools              tool.Tool `json:"tools"`
	EnableThinking     bool      `json:"enable_thinking"`
}

// ChatResponse 定义 Responses API 的响应体（仅提取需要的字段）
type response struct {
	ID     string `json:"id"`
	Output []struct {
		Type    string `json:"type"` // "reasoning" 或 "message"
		Content []struct {
			Text string `json:"text"` // 最终回复文本（仅 message 类型有）
			Type string `json:"type"` // 通常是 "output_text"
		} `json:"content,omitempty"`
		// 对于 reasoning 类型，会有 summary 字段，但我们不关心
	} `json:"output"`
}
