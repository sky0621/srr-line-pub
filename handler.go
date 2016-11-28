package srrlinepub

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
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

// ServeHTTPC ...
func (p *PubHandler) ServeHTTPC(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}
