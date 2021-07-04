package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host, _ := os.Hostname()
		w.Write([]byte(host + "hogeho77777ge"))
	})

	// コントローラー割当
	// カラオケ楽曲コントローラー
	// sqlHandler := infrastructure.NewSqlHandler()
	// musicController := controllers.NewMusicsController(sqlHandler)
	// http.Handle("/musics", )

	// dev環境では"localhost:8080"を指定、本番では":8080"を指定
	http.ListenAndServe("localhost:8080", nil)
}
