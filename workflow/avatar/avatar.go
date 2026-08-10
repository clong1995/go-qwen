package avatar

import (
	"github.com/clong1995/go-qwen/image"
	"github.com/clong1995/go-qwen/text"
)

func Generate(prompt string) (string, error) {
	//提取关键词
	s, _, err := text.Generate(
		"阅读文本，构思并输出一个高质量的人像摄影头像，画布 512*512 px，直接输出结果:"+prompt,
		"你现在是一位专业的头像设计师。", "")
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
