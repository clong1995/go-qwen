package image

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	conf "github.com/clong1995/go-qwen/config"
	"github.com/pkg/errors"
)

// Generate 生成图片
// prompt 提示词
// size 设置输出图像分辨率 "宽*高" 512*512 – 2048*2048
// negativePrompt 描述需要从图像中排除的元素
// n 设置生成图片数量
// promptExtendMode 选择提示词改写方式,direct:直接提示词增强（DPE），,agent:智能体提示词增强（APE）
func Generate(prompt, size, negativePrompt string, promptExtendMode PromptExtendMode, n int64, reference []string) (string, error) {
	url := "https://dashscope.aliyuncs.com/api/v1/services/aigc/multimodal-generation/generation"

	// 构造请求体
	reqBody := request{
		Model: conf.ImageModel,
		Input: input{
			Messages: []messages{
				{
					Role:    "user",
					Content: []content{},
				},
			},
		},
		Parameters: parameters{
			NegativePrompt:   negativePrompt,
			PromptExtendMode: string(promptExtendMode),
			PromptExtend:     promptExtendMode != PromptExtendNone,
			Watermark:        false,
			Size:             size,
			N:                n,
		},
	}

	//参考图
	for _, v := range reference {
		reqBody.Input.Messages[0].Content = append(reqBody.Input.Messages[0].Content, content{Image: v})
	}

	//提示词
	reqBody.Input.Messages[0].Content = append(reqBody.Input.Messages[0].Content, content{Text: prompt})

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", errors.WithStack(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", errors.WithStack(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+conf.ApiKey)

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.WithStack(err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("%d:%s", resp.StatusCode, string(body))
	}

	var imgResp response
	if err = json.Unmarshal(body, &imgResp); err != nil {
		return "", errors.WithStack(err)
	}

	// 提取图片 URL
	if len(imgResp.Output.Choices) > 0 &&
		len(imgResp.Output.Choices[0].Message.Content) > 0 &&
		imgResp.Output.Choices[0].Message.Content[0].Image != "" {
		return imgResp.Output.Choices[0].Message.Content[0].Image, nil
	}

	return "", errors.Errorf("%s", string(body))
}

var httpClient = &http.Client{}

type PromptExtendMode string

const (
	// PromptExtendDirect 直接增强（DPE），大多数场景
	PromptExtendDirect PromptExtendMode = "direct"
	// PromptExtendAgent 智能体增强（APE），改写更精细,提示词很简短、需要模型补充大量细节时
	PromptExtendAgent PromptExtendMode = "agent"
	// PromptExtendNone 不增强（空值）
	PromptExtendNone PromptExtendMode = ""
)

type request struct {
	Model      string     `json:"model"`
	Input      input      `json:"input"`
	Parameters parameters `json:"parameters"`
}

type response struct {
	Output struct {
		Choices []struct {
			Message struct {
				Content []struct {
					Image string `json:"image"`
				} `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	} `json:"output"`
}

type input struct {
	Messages []messages `json:"messages"`
}

type messages struct {
	Role    string    `json:"role"`
	Content []content `json:"content"`
}

type content struct {
	Text  string `json:"text,omitempty"`
	Image string `json:"image,omitempty"`
}

type parameters struct {
	NegativePrompt   string `json:"negative_prompt,omitempty"`
	PromptExtendMode string `json:"prompt_extend_mode,omitempty"`
	PromptExtend     bool   `json:"prompt_extend,omitempty"`
	Watermark        bool   `json:"watermark,omitempty"`
	Size             string `json:"size,omitempty"`
	N                int64  `json:"n,omitempty"`
}
