package controllers

import (
	"github.com/gorilla/websocket"
	"net/http"
	"amp-templates/server/services/hot_reloading"
	"amp-templates/server/services/log"
)

type ViewModel struct {
	Title string
}

//Home index route
func Home(w http.ResponseWriter, r *http.Request) {
	executeTemplate("index", w, r, &ViewModel{Title: ""})
}

func HotReload(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil { return }

	hot_reloading.Configure(func() {
		err = c.WriteMessage(websocket.TextMessage, []byte("reload"))
		if err != nil {
			log.Info(err)
		}
		c.Close()
	})
}