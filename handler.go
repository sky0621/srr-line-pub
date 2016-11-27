package srrlinepub

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"goji.io/pat"
)

// PubHandler ...
type PubHandler struct {
	ChannelSecret string
	AccessToken   string
}

// ToSQS ...
func (p *PubHandler) ToSQS(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(ctx, "name")
	fmt.Fprintf(w, "Hi %s", name)
}
