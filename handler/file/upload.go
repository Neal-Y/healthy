package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"healthy/config"
	"healthy/constant"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func (h *File) UploadHandler(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error uploading file"})
		return
	}

	// 打开文件
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
		return
	}
	defer src.Close()

	saveDir := "uploads"
	err = os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating save directory"})
		return
	}

	savePath := filepath.Join(saveDir, file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}

	fileModel, err := h.service.UploadFileLogic(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file info"})
		return
	}

	// 生成文件的 URL
	fileURL := fmt.Sprintf("https://3db0-203-204-68-236.ngrok-free.app/%s", fileModel.FilePath)
	log.Printf("fileURL: %s", fileURL)
	// 构建请求的 payload
	requestBody := map[string]interface{}{
		"model": "gpt-4o",
		"messages": []map[string]interface{}{
			{
				"role":    "system",
				"content": "You are a highly accurate and consistent nutritionist. Your task is to analyze the image of the food and provide the total nutritional values (calories, protein, and fiber) in a strict JSON format. Please ensure that the JSON does not include any extra formatting such as code blocks. The output should only be in the form of: {\"calories\": \"\", \"protein\": \"\", \"fiber\": \"\"}.",
			},
			{
				"role":    "user",
				"content": "Analyze the following image for its total combined nutritional content.",
			},
			{
				"role": "user",
				"content": []map[string]interface{}{
					{
						"type": "image_url",
						"image_url": map[string]string{
							"url": fileURL, // The actual URL of the image to analyze
						},
					},
				},
			},
		},
		"max_tokens": 300,
	}

	// 发送请求
	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+config.AppConfig.GPTApiKey).
		SetBody(requestBody).
		Post(constant.ChatApiUrl)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error calling OpenAI API"})
		return
	}

	log.Printf("response: %s", resp.String())
	// 输出 API 回应
	c.JSON(http.StatusOK, gin.H{"response": resp.String()})
}
