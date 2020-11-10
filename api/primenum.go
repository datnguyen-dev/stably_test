package api

import (
	"fmt"
	"net/http"
	"strconv"

	"datnguyen.stably.test/store"
	"github.com/go-chi/chi"
)

type primenum struct {
	*Handler
}

func (h *primenum) getHighestPrimeLowerInput(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(chi.URLParam(r, "number"))
	if err != nil {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}
	res := store.FindPrime(num)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("%v", res)))
}
