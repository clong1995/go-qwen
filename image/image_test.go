package image

import "testing"

func TestGenerate(t *testing.T) {
	type args struct {
		apiKey           string
		model            string
		prompt           string
		size             string
		negativePrompt   string
		promptExtendMode PromptExtendMode
		n                int64
		watermark        bool
		reference        []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Generate",
			args: args{
				prompt:           "画面是一张竖幅户外人像摄影，整体从上到下呈现温暖的午后街景氛围。顶部左侧到上方大面积被深绿色藤蔓和橙色小花覆盖，花叶从建筑檐口自然垂落，受阳光照射的叶片呈黄绿色高光，阴影处则偏深绿，形成浓密而柔和的背景层次。左上至中上区域是一块深蓝色横向招牌，招牌表面较暗、略带磨砂质感，上面以白色哥特体大字写着 你好世界，文字位于画面左侧偏上，部分被前景花叶轻微遮挡，字体高对比、带装饰性尖角和粗细变化。招牌下方是报刊亭或书报摊的玻璃展示窗，黑色金属框架将橱窗分隔成多个矩形区域，内部陈列着许多报纸、杂志和书刊封面，但大多因景深虚化和光线反射而难以辨读，形成浅色纸张与深色边框交错的背景纹理。画面右上方是强烈的逆光区域，阳光从街道尽头照入，背景建筑被虚化成米灰色块面，边缘柔和，呈现明显的浅景深效果。画面中部偏右是一名年轻女性的全身人像，她回头面向镜头微笑，身体略向右转，肩背朝向观者，姿态自然放松。她有长而浓密的黑色波浪卷发，发丝被逆光勾勒出金色轮廓光，发梢在右侧向外散开，显得轻盈蓬松。她肤色白皙，脸型柔和偏鹅蛋形，眉形细致，眼睛明亮，眼妆清透，睫毛明显，面部带有自然高光，唇部为柔和珊瑚红色，笑容露齿，表情亲切明朗。她佩戴小巧耳饰，身穿黑色细肩带露背连衣短裙，面料颜色深黑、轮廓简洁，细肩带从肩部向背部延伸，背部线条清晰，腿部光滑白净，脚上穿着高跟鞋。画面偏左到中部，她双手抱着一束玫瑰花，花束体积较大，主要由橙色、杏色、粉色和浅桃色玫瑰组成，花瓣层层卷曲，边缘被阳光照亮，绿色叶片和长花茎从花束下方垂出，花束与黑色裙装形成鲜明色彩对比。右侧背景是一条被阳光照亮的城市街道，地面呈暖灰与金黄色调，远处建筑、街边设施和一个模糊的红色圆形交通标志位于右下远景，均因焦外虚化而只保留色块和轮廓。整张照片采用暖色胶片感处理，带有细腻颗粒、柔和对比和明显逆光边缘光，人物位于视觉焦点，背景报刊亭、花藤、街道和阳光共同营造出浪漫、明亮、都市漫步式的氛围。",
				size:             "512*512",
				negativePrompt:   "路人,车辆",
				promptExtendMode: PromptExtendAgent,
				n:                1,
				watermark:        false,
				reference:        []string{},
			},
		},
		{
			name: "Generate-Reference",
			args: args{
				prompt:           "图1中的女孩穿上图2中的红色连衣裙",
				size:             "512*512",
				negativePrompt:   "",
				promptExtendMode: PromptExtendNone,
				n:                1,
				watermark:        false,
				reference: []string{
					"https://partner-storage.oss-cn-beijing.aliyuncs.com/t2.png",
					"https://partner-storage.oss-cn-beijing.aliyuncs.com/t1.png",
				},
			},
		},
		{
			name: "Generate-Reference-1",
			args: args{
				prompt:           "图1改为全身照",
				size:             "512*512",
				negativePrompt:   "",
				promptExtendMode: PromptExtendNone,
				n:                1,
				reference: []string{
					"https://partner-storage.oss-cn-beijing.aliyuncs.com/t2.png",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.prompt, tt.args.size, tt.args.negativePrompt, tt.args.promptExtendMode, tt.args.n, tt.args.reference)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Generate() got = %v", got)
		})
	}
}
