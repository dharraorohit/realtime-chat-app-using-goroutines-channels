package srvhandlers

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/dharraorohit/realtime-chat-app-using-goroutines-channels/utils/responseUtils"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn        *websocket.Conn
	messageChan chan MessageData
	userId      int
}

type WebSocketHandler struct {
	Clients map[int]*Client
	mutex   sync.Mutex
}

type MessageData struct {
	FromUserId int
	ToUserId   int
	Content    string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (wsh *WebSocketHandler) HandleConnection(w http.ResponseWriter, r *http.Request) {

	userIdStr := r.URL.Query().Get("userId")
	userId, err := strconv.Atoi(userIdStr)
	if userIdStr == "" || err != nil {
		responseUtils.ErrorResponse(http.StatusBadRequest, "UserId is missing or wrong format", w)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{conn: conn, messageChan: make(chan MessageData), userId: userId}
	wsh.addClient(client)

	go client.writeMessageRoutine()
	client.readMessageRoutine(wsh)

}

func (wsh *WebSocketHandler) addClient(client *Client) {
	wsh.mutex.Lock()
	defer wsh.mutex.Unlock()

	wsh.Clients[client.userId] = client
}

func (wsh *WebSocketHandler) removeClient(userId int) {
	wsh.mutex.Lock()
	defer wsh.mutex.Unlock()

	delete(wsh.Clients, userId)
}

func (c *Client) writeMessageRoutine() {
	defer c.conn.Close()

	for {
		messageData, ok := <-c.messageChan
		if !ok {
			return
		}
		err := c.conn.WriteJSON(messageData)
		if err != nil {
			return
		}
	}
}

func (c *Client) readMessageRoutine(wsh *WebSocketHandler) {
	defer func() {
		wsh.removeClient(c.userId)
		c.conn.Close()
		close(c.messageChan)
	}()

	for {
		var messageData *MessageData
		err := c.conn.ReadJSON(&messageData)
		if err != nil {
			log.Printf("Error in reading Message: %v", err)
			break
		}

		messageData.FromUserId = c.userId
		toClient, ok := wsh.Clients[messageData.ToUserId]
		if ok {
			select {
			case toClient.messageChan <- *messageData:
			default:
				log.Printf("Channel full of %d", messageData.ToUserId)
			}
		}

	}
}
