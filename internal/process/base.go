package process

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// JSONData 定义一个通用的 map 结构来存储 JSON 对象
type JSONData struct {
	Instruction string `json:"instruction"`
	InputVal    string `json:"inputVal"`
	OutputVal   string `json:"outputVal"`
	LlmSource   string `json:"llmSource"`
}

func ExtractJSONArray(text string) ([]JSONData, error) {
	// 使用 (?s) 让 . 匹配包括换行符在内的所有字符，匹配到整个 JSON 数组
	jsonArrayRegex := regexp.MustCompile(`\[[\s\S]*?\]`)

	// 查找匹配的 JSON 数组
	matches := jsonArrayRegex.FindString(text)
	if matches == "" {
		return nil, fmt.Errorf("未找到匹配的 JSON 数组")
	}

	// 存储解析后的 JSON 数据
	var jsonArray []JSONData

	// 将匹配的字符串解析为 JSON 数组
	err := json.Unmarshal([]byte(matches), &jsonArray)
	if err != nil {
		return nil, fmt.Errorf("解析 JSON 时出错: %v", err)
	}

	return jsonArray, nil
}
