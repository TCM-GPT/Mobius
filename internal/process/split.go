package process

import (
	"strings"
)

func Split(content string, splitPrefix string, restrict uint64) []string {
	// 按照 ”。“ 将文章分割
	content = strings.TrimSpace(content)
	segments := strings.Split(content, splitPrefix)

	// 截断算法
	var results []string
	var buffer string

	// 判断句子长度是否超过了 restrict 限定值，如过超过则作为一句存储到 results 中，如果没有超过则拼接后面一句到 result 中
	for _, segment := range segments {
		if segment == "" {
			break
		}

		if len(buffer) > int(restrict) {
			//fmt.Println(len(buffer))
			results = append(results, buffer)
			buffer = segment
		} else {
			buffer += segment + splitPrefix
		}
	}

	results = append(results, buffer)
	//fmt.Println("==================================================")
	return results
}
