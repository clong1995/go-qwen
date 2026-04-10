package text

import "testing"

func TestSend(t *testing.T) {
	type args struct {
		msg []Message
	}
	tests := []struct {
		name        string
		args        args
		wantContent string
		wantErr     bool
	}{
		{
			name: "对话",
			args: args{
				msg: []Message{
					SystemMsg("你是一个医疗助手"),
					UserMsg("我牙疼"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotContent, err := Send(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Send() gotContent = %v", gotContent)
		})
	}
}
