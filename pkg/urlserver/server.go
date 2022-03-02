package urlserver

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"eywa/pkg/shortener"
)

type Config struct {
	Port string
}

type Server struct {
	Config
}

func NewServer(c Config) Server {
	return Server{
		c,
	}
}

func (us Server) Start() error {
	s := http.Server{
		Addr:    us.Port,
		Handler: initRouter(),
	}
	log.Printf("Starting server on %s", us.Port)
	err := s.ListenAndServe()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func initRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/{shortHash}", actualData)
	r.NotFound(handle404)
	return r
}

func actualData(w http.ResponseWriter, r *http.Request) {
	hash := chi.URLParam(r, "shortHash")
	if data, ok := shortener.Short[hash]; !ok || time.Now().Sub(data.GetLastUpdatedTime()).Hours() >= 24 {
		handle404(w, r)
	} else {
		http.Redirect(w, r, data.GetContent(), http.StatusSeeOther)
	}
}

func handle404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Payment URL not found/expired"))
}
