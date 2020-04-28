// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var tmpl *template.Template

func init() {
	var terr error
	tmpl, terr = template.New("home.html").ParseFiles("room.html", "home.html")
	if terr != nil {
		log.Fatalln(terr)
	}
}

func serveRoom(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	data := map[string]interface{}{
		"RoomID": params.ByName("RoomID"),
	}

	err := tmpl.ExecuteTemplate(w, "room.html", data)
	if err != nil {
		log.Println("Could not initiate template err: ", err)
	}
}

func ServeHome(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := tmpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		log.Println("Could not initiate template err: ", err)
	}
}

func main() {
	go h.run()
	router := httprouter.New()
	router.GET("/:RoomID", serveRoom)
	router.GET("/:RoomID/ws", serveWs)
	router.GET("/", ServeHome)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
