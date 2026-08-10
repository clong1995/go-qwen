package icon

import "testing"

func TestGenerate(t *testing.T) {
	type args struct {
		apiKey     string
		textModel  string
		imageModel string
		prompt     string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "icon",
			args: args{
				prompt: "早安问候，在群内发送干货或活动链接",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.prompt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("Generate() got = %v", got)

		})
	}
}
