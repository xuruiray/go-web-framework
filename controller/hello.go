package controller

import (
	"fmt"
	"github.com/xuruiray/go-web-framework/service"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(service.TestService(r.Context(), 1))
	w.Write([]byte("hello"))
}
