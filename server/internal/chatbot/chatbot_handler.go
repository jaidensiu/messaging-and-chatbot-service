package chatbot

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	cohere "github.com/cohere-ai/cohere-go/v2"
	client "github.com/cohere-ai/cohere-go/v2/client"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Handler struct {
	cohereClient *client.Client
}

type ChatRequest struct {
	Message string `json:"message"`
}

type CohereResponse struct {
	Text string `json:"text"`
}

func NewHandler() *Handler {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	cohereAPIKey := os.Getenv("COHERE_API_KEY")

	return &Handler{
		cohereClient: client.NewClient(client.WithToken(cohereAPIKey)),
	}
}

func (h *Handler) StreamMessage(c *gin.Context) {
	var chatReq ChatRequest
	if err := c.ShouldBindJSON(&chatReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	resp, err := h.cohereClient.ChatStream(
		context.TODO(),
		&cohere.ChatStreamRequest{
			ChatHistory: []*cohere.Message{
				{
					Role: "USER",
					User: &cohere.ChatMessage{
						Message: chatReq.Message,
					},
				},
			},
			Message: chatReq.Message,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response from Cohere"})
		return
	}
	defer resp.Close()

	c.Stream(func(w io.Writer) bool {
		message, err := resp.Recv()

		if errors.Is(err, io.EOF) {
			return false
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading from stream"})
			return false
		}

		if message.TextGeneration != nil {
			cohereResp := CohereResponse{
				Text: message.TextGeneration.Text,
			}
			if err := json.NewEncoder(w).Encode(cohereResp); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode response"})
				return false
			}
		}
		return true
	})
}
