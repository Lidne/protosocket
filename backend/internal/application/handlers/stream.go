package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"

	pb "main/internal/proto"
)

type WebSocketHandler struct {
	logger   *logrus.Logger
	upgrader websocket.Upgrader
}

func NewWebSocketHandler(logger *logrus.Logger) *WebSocketHandler {
	return &WebSocketHandler{logger: logger, upgrader: websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}}
}

func (handler *WebSocketHandler) HandleStream(w http.ResponseWriter, r *http.Request) {
	// Update connection to WebSocket
	handler.logger.Infoln("Started streaming")
	ws, err := handler.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Start a goroutine to read messages (to handle close frames)
	go func() {
		for {
			if _, _, err := ws.ReadMessage(); err != nil {
				handler.logger.Infof("Read error (connection closed?): %v", err)
				return
			}
		}
	}()

	// Stream messages
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var i int32 = 0
	for range ticker.C {
		i++
		text := fmt.Sprintf("Message number %d", i)
		title := "My proto message"

		msg := &pb.MyMessage{
			Title: &title,
			Id:    &i,
			Text:  &text,
		}

		data, err := proto.Marshal(msg)
		if err != nil {
			handler.logger.Error("Marshaling error: ", err)
			continue
		}

		if err := ws.WriteMessage(websocket.BinaryMessage, data); err != nil {
			handler.logger.Error("Write error: ", err)
			break
		}
		handler.logger.Infof("Sent message %d", i)
	}
}
