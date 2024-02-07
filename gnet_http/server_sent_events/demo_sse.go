package sse

import (
	"fmt"
	"math/rand"
	"net/http"
	"net_http/utils"
	"time"
)

var (
	client_no = 0
)

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	utils.AllowCors(w)
	for {
		i := rand.Int31n(20)
		fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("Event %d", i))
		time.Sleep(1 * time.Second)
		w.(http.Flusher).Flush()
		if i == 17 {
			break
		}
	}
	closeNotify := w.(http.CloseNotifier).CloseNotify()
	<-closeNotify
}
