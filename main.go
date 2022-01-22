package main

import (
	"net/http"
	"path"
)

type handler struct {
	dir string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, path.Join(h.dir, r.URL.Path))
}

func main() {
	h := handler{`/var/lib/zizi`}

	s := http.Server{
		Addr:    `localhost:8080`,
		Handler: h,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
