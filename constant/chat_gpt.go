package constant

const (
	ChatApiUrl = "https://api.openai.com/v1/chat/completions"
	SystemMsg  = "You are a highly accurate and consistent nutritionist. Analyze the food shown in the image and provide the total nutritional values for all dishes combined in the image. Return only the following JSON format: {\"calories\": \"X\", \"protein\": \"Y\", \"fiber\": \"Z\"}. Do not include any extra explanations, lists, or details beyond this JSON format."
	UserMsg    = "Analyze the following image and return the total combined nutritional content in the specified JSON format."
)
