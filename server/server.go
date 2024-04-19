package server

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Payload represents the data structure sent to the websocket endpoint.
type Payload struct {
	Message string                 `json:"message"`
	Headers map[string]interface{} `json:"HEADERS"`
}

// ChatMessageTemplate represents the data structure for the chat-message.tmpl.html template.
type ChatMessageTemplate struct {
	Message     string
	MessageIcon string
}

// upgrader is used to upgrade the http connection to a websocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// processMessage takes in the message sent to the websocket and returns a reply
func processMessage(question string) string {
	// echo the question in reverse
	chars := []rune(question)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

// HandleWebSocket handles the request for the web-socket endpoint.
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// upgrade the connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Websocket Client Connected")

	for {
		// read in a message
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// Parse payload
		var dat Payload
		if err := json.Unmarshal(p, &dat); err != nil {
			log.Println(err)
			return
		}

		// Load template for repsonses
		chatMessageTemplate := template.Must(template.ParseFiles("server/templates/chat-message.tmpl.html"))

		// Render question to message list
		var questionBuffer bytes.Buffer
		chatMessageTemplate.Execute(&questionBuffer, ChatMessageTemplate{
			Message:     dat.Message,
			MessageIcon: "person",
		})
		ws.WriteMessage(messageType, questionBuffer.Bytes())

		// Render answer to message list
		var answerBuffer bytes.Buffer
		chatMessageTemplate.Execute(&answerBuffer, ChatMessageTemplate{
			Message:     processMessage(dat.Message),
			MessageIcon: "smart_toy",
		})
		ws.WriteMessage(messageType, answerBuffer.Bytes())
	}
}
