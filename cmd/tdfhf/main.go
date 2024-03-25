package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

type reqBody struct {
	Message string `json:"message"`
}

func HandleError(err error, w http.ResponseWriter) {
	fmt.Println(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Failed to send message!"))
}

func main() {
	envFile, _ := godotenv.Read()

	mux := http.NewServeMux()
	handler := cors.New(cors.Options{
		AllowedOrigins: strings.Split(envFile["ALLOWED_ORIGINS"], ","),
	}).Handler(mux)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running!"))
	})

	mux.HandleFunc("POST /sendMessage", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			HandleError(err, w)
			return
		}

		var requestBody reqBody

		if err := json.Unmarshal(body, &requestBody); err != nil {
			HandleError(err, w)
			return
		}

		telegramApiUrl := "https://api.telegram.org/bot" + envFile["TELEGRAM_BOT_API_TOKEN"] + "/sendMessage"

		postBody := map[string]string{
			"chat_id": envFile["TELEGRAM_CHANNEL_CHAT_ID"],
			"text":    requestBody.Message,
		}

		bytesBuffer, _ := json.Marshal(postBody)

		res, _ := http.Post(telegramApiUrl, "application/json", bytes.NewBuffer(bytesBuffer))

		if res.StatusCode != http.StatusOK {
			err := fmt.Errorf("FAILED to send message! Status code: %d", res.StatusCode)
			HandleError(err, w)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Message sent successfully!"))
	})

	if err := http.ListenAndServe(":8080", handler); err != nil {
		fmt.Println(err.Error())
	}
}
