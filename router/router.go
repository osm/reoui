package router

import (
	"embed"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"

	"github.com/osm/reoui/reolink"
)

var filenameDateRe = regexp.MustCompile(`[0-9]{14}`)

type Router struct {
	frontendFS embed.FS
	graphql    *handler.Server
	dataDir    string
	reolinks   []*reolink.Client
}

func NewRouter(opts ...Option) *chi.Mux {
	ro := &Router{}

	for _, opt := range opts {
		opt(ro)
	}

	r := chi.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)
	r.Use(func(next http.Handler) http.Handler {
		logger := middleware.RequestLogger(
			&middleware.DefaultLogFormatter{
				Logger:  log.New(os.Stdout, "", log.LstdFlags),
				NoColor: true,
			},
		)
		return logger(next)

	})
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Handle("/playground", playground.Handler("Playground", "/graphql"))
	r.Handle("/graphql", ro.graphql)

	r.Get("/stream/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, _ := strconv.Atoi(idStr)

		if id > len(ro.reolinks)-1 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found\r\n"))
			return
		}

		resp, err := ro.reolinks[id].Stream(w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error\r\n"))
			return
		}
		defer resp.Body.Close()

		io.Copy(w, resp.Body)
	})

	r.Get("/files/{filename}", func(w http.ResponseWriter, r *http.Request) {
		filename := chi.URLParam(r, "filename")
		date := filenameDateRe.FindString(filename)
		if len(date) == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found\r\n"))
			return
		}

		d, err := os.ReadFile(path.Join(ro.dataDir, date[0:4], date[4:6], date[6:8], filename))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found\r\n"))
			return
		}

		w.Write(d)
	})

	files, _ := fs.ReadDir(ro.frontendFS, "frontend/dist")
	if len(files) > 0 {
		for _, f := range files {
			n := f.Name()
			c, _ := ro.frontendFS.ReadFile(filepath.Join("frontend", "dist", n))

			if n == "index.html" {
				r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
					w.Write(c)
				})
			} else {
				r.Get("/"+n, func(w http.ResponseWriter, r *http.Request) {
					w.Write(c)
				})
			}
		}
	}

	return r
}
