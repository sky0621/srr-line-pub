package srrlinepub

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"goji.io/pat"
)

// Hi ...
func Hi(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(ctx, "name")
	fmt.Fprintf(w, "Hi %s", name)
}
