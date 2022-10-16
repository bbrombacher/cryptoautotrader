package coinbase

import (
	"errors"

	"github.com/google/uuid"
	ws "github.com/gorilla/websocket"
)

type Client struct {
	socketURL string
	tickers   map[string]*ws.Conn
}

func New(socketURL string) *Client {
	return &Client{
		tickers:   make(map[string]*ws.Conn),
		socketURL: socketURL,
	}
}

func (c Client) StartTickerFeed(params StartTickerParams) (string, error) {
	var wsDialer ws.Dialer
	wsConn, _, err := wsDialer.Dial(c.socketURL, nil)
	if err != nil {
		return "", err
	}

	if err = wsConn.WriteJSON(params); err != nil {
		return "", err
	}

	id := uuid.New()
	c.tickers[id.String()] = wsConn
	return id.String(), nil
}

func (c Client) GetTickerMessages(id string) (TickerMsg, error) {
	var message TickerMsg
	wsConn, ok := c.tickers[id]
	if !ok {
		return message, errors.New("id does not exist")
	}

	err := wsConn.ReadJSON(&message)
	if err != nil {
		return message, err
	}
	return message, nil
}

func (c Client) CloseTickerFeed(id string) error {
	wsConn, ok := c.tickers[id]
	if !ok {
		return errors.New("id does not exist")
	}

	wsConn.Close()
	return nil
}
