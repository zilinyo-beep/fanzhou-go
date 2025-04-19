package strategy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	"fmt"
	"net/http"
)

type ImagineStrategy struct{}

func (s *ImagineStrategy) Execute(ctx context.Context, task *entity.Task) error {
	var params struct {
		Prompt string `json:"prompt"`
	}
	if err := json.Unmarshal([]byte(task.Params), &params); err != nil {
		return err
	}

	payload := map[string]string{"prompt": params.Prompt}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "https://cl.imagineapi.dev/items/images/", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer <YourToken>")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Midjourney API 请求失败")
	}

	fmt.Println("Imagine任务已提交:", params.Prompt)
	return nil
}
