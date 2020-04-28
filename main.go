// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

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

var addr = flag.String("addr", ":8080", "http service address")

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
	flag.Parse()
	go h.run()
	router := httprouter.New()
	router.GET("/:RoomID", serveRoom)
	router.GET("/:RoomID/ws", serveWs)
	router.GET("/", ServeHome)

	err := http.ListenAndServe(*addr, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
