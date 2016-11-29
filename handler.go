package srrlinepub

import (
	"fmt"
	"net/http"
)

// PubHandler ...
type PubHandler struct {
	ChannelSecret string
	AccessToken   string
}

// NewPubHandler ...
func NewPubHandler(channelSecret string, accessToken string) *PubHandler {
	h := &PubHandler{
		ChannelSecret: channelSecret,
		AccessToken:   accessToken,
	}
	return h
}

// ServeHTTP ...
func (p *PubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}
