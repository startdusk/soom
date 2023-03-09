package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ServiceWeaver/weaver"
	"github.com/startdusk/soom/chatservice"
)

func main() {
	ctx := context.Background()
	root := weaver.Init(ctx)

	chatService, err := weaver.Get[chatservice.Service](root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	r := http.NewServeMux()
	r.HandleFunc("/chat/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, fmt.Sprintf("Method %q not allowed", r.Method), http.StatusMethodNotAllowed)
			return
		}

		name := r.URL.Query().Get("name")
		if len(name) == 0 {
			http.Error(w, "name is not defined in the url", http.StatusBadRequest)
			return
		}

		chat, err := chatService.GetChat(r.Context(), name)
		if err != nil {
			fmt.Printf("GetByPost: %s", err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(chat)
		if err != nil {
			fmt.Printf("Encode: %s", err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	lis, err := root.Listener("main", weaver.ListenerOptions{LocalAddress: ":18990"})
	if err != nil {
		fmt.Printf("root.Listener: %s", err.Error())
		return
	}

	err = http.Serve(lis, r)
	if err != nil {
		fmt.Printf("http.Serve: %s", err.Error())
		return
	}
}
