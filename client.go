package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// FlomoNote 参数结构体
type FlomoNote struct {
	Content   string   `json:"content"`
	ImageURLs []string `json:"image_urls"`
}

// PostFlomoNote 函数用于发送 Flomo 笔记。
func PostFlomoNote(token string, note *FlomoNote) error {
	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", token, nil)
	if err != nil {
		return err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 设置请求体
	body, err := json.Marshal(note)
	if err != nil {
		return err
	}
	req.Body = io.NopCloser(bytes.NewReader(body))

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
