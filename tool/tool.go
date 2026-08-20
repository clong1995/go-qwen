package tool

type Tool []map[string]string

// NoTool 联网搜索工具
var NoTool Tool

// WebSearch 联网搜索工具
var WebSearch Tool = []map[string]string{{"type": "web_search"}}

// WebExtractor 网页抽取工具
var WebExtractor Tool = []map[string]string{{"type": "web_search"}, {"type": "web_extractor"}}

// CodeInterpreter 代码解释器工具
var CodeInterpreter Tool = []map[string]string{{"type": "code_interpreter"}}

// WebSearchImage 根据文本描述搜索图片
var WebSearchImage Tool = []map[string]string{{"type": "web_search_image"}}

// ImageSearch 根据图片搜索相似或相关图片
var ImageSearch Tool = []map[string]string{{"type": "image_search"}}
