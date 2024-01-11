package statichttp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var srv http.Server

func RunStaticHttpServer(rootPath string) func(files []string) {
	srvMux := http.NewServeMux()

	srvMux.HandleFunc("/serverinfo", handleGetServerInfo)
	srvMux.Handle("/", http.FileServer(http.Dir(rootPath)))

	srv = http.Server{
		Handler: srvMux,
		Addr:    "127.0.0.1:8000",
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	return func(files []string) {
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Println("Error when stopping the server", err)
		}

		srv = http.Server{
			Handler: srvMux,
			Addr:    "127.0.0.1:8000",
		}
		go func() {
			if err := srv.ListenAndServe(); err != nil {
				fmt.Println(err)
			}
		}()
	}
}

func handleGetServerInfo(w http.ResponseWriter, r *http.Request) {
	info := map[string]string{
		"author":     "Sifatul (sifatul@sifatul.com)",
		"type":       "static-http",
		"autoreload": "no",
		"repository": "https://github.com/sifatulrabbi/filepatrol",
	}
	if b, err := json.Marshal(info); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(200)
		w.Write(b)
	}
}
