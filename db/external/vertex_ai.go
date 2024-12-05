package external

import (
	"cloud.google.com/go/vertexai/genai"
	"context"
	"fmt"
	"log"
)

const (
	projectID = "term6-mana-nakagawa" // TODO:
	location  = "asia-northeast1"
	modelName = "gemini-1.5-flash-002"
)

const yes = genai.Text("yes")

// CheckHarmfulContent はVertex AIを使用して投稿内容に誹謗中傷が含まれているかをチェックします
func CheckHarmfulContent(content string) (genai.Part, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, projectID, location)
	if err != nil {
		return nil, fmt.Errorf("vertex AIクライアントの作成に失敗しました: %w", err)
	}

	gemini := client.GenerativeModel(modelName)
	prompt := genai.Text(fmt.Sprintf("Does the following content contain harmful language? Respond with 'yes' or 'no': %s", content))
	resp, err := gemini.GenerateContent(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("vertex AIでの生成に失敗しました: %w", err)
	}

	// Candidateの構造体
	type Candidate struct {
		Content struct {
			Parts []string `json:"parts"` // 文字列の配列として定義
		} `json:"content"`
	}

	var part genai.Part
	// 判定ロジック
	for _, candidate := range resp.Candidates {
		if candidate.Content.Parts != nil { // Content.Parts が存在するか確認

			for _, part := range candidate.Content.Parts { // Parts 配列をループ
				log.Println(part) // part はすでに文字列型なので直接比較
				//if part == yes {  // 空白をトリムして比較
				//	log.Println(part)
				//	return nil, nil
				//}
				return part, nil
			}
		}
	}

	log.Printf("part= %v", part)
	return part, nil

}
