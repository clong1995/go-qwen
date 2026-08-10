package text

import (
	"testing"
)

func TestSendMessage(t *testing.T) {
	type args struct {
		apiKey             string
		model              string
		userInput          string
		instructions       string
		previousResponseID string
	}
	tests := []struct {
		name              string
		args              args
		wantReply         string
		wantNewResponseID string
		wantErr           bool
	}{
		{
			name: "Generate",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instructions := "你是一位专业的园艺顾问，请用专业且友好的语气回答问题。"
			var previousID string

			// 第一轮对话
			userInput1 := "我家的多肉植物叶子变黄了，怎么办？"
			reply1, id1, err := Generate(userInput1, instructions, previousID)
			if err != nil {
				t.Errorf("%+v", err)
				return
			}
			previousID = id1
			t.Logf("用户: %s\n", userInput1)
			t.Logf("助手: %s\n", reply1)

			// 第二轮对话（延续上下文）
			userInput2 := "那我应该多久浇一次水呢？"
			reply2, id2, err := Generate(userInput2, instructions, previousID)
			if err != nil {
				t.Errorf("%+v", err)
				return
			}
			previousID = id2 // 继续保存
			t.Logf("用户: %s\n", userInput2)
			t.Logf("助手: %s\n", reply2)
		})
	}
}
