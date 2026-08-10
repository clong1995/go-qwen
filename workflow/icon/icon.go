package icon

import (
	"github.com/clong1995/go-qwen/image"
	"github.com/clong1995/go-qwen/text"
)

func Generate(prompt string) (string, error) {
	//提取关键词
	s, _, err := text.Generate(
		"阅读文本，构思并输出一个符合 iconfont 平台规范的面性图标设计的详细视觉描述，极简，单一尺寸，画布 512*512 px，明亮鲜艳纯色完全填充背景，背景形状为直角矩形，直接输出结果:"+prompt,
		"你现在是一位专业的UI图标设计师。", "")
	if err != nil {
		return "", err
	}

	//生成图片
	generate, err := image.Generate(s, "521*512", "", image.PromptExtendNone, 1, []string{})
	if err != nil {
		return "", err
	}
	return generate, err
}
