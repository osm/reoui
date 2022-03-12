package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/osm/flen"

	"github.com/osm/reoui/clean"
	"github.com/osm/reoui/config"
	"github.com/osm/reoui/graphql"
	"github.com/osm/reoui/reolink"
	"github.com/osm/reoui/router"
	"github.com/osm/reoui/sync"
)

//go:embed frontend/dist
var frontendFS embed.FS

func main() {
	configPath := flag.String("config", "", "Config path")
	flen.Parse()
	cfg, err := config.NewConfig(*configPath)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)

	}
	if cfg.Port == "" {
		fmt.Printf("error: no port defined\n")
		os.Exit(1)
	}
	if cfg.DataDir == "" {
		fmt.Printf("error: no data_dir defined\n")
		os.Exit(1)
	}
	if len(cfg.Cameras) == 0 {
		fmt.Printf("error: no cameras defined\n")
		os.Exit(1)
	}

	var reolinks []*reolink.Client
	for _, c := range cfg.Cameras {
		reolinks = append(reolinks, reolink.NewClient(&c))
	}

	if int64(cfg.SyncInterval) > 0 {
		s := sync.New(
			sync.WithDataDir(cfg.DataDir),
			sync.WithSyncInterval(cfg.SyncInterval),
			sync.WithReolinks(reolinks),
		)
		go s.Run()
	}

	if int64(cfg.CleanFilesInterval) > 0 {
		c := clean.New(
			clean.WithDataDir(cfg.DataDir),
			clean.WithCleanFilesInterval(cfg.CleanFilesInterval),
		)
		go c.Run()
	}

	resolverOpts := []graphql.ResolverOption{
		graphql.WithCameras(cfg.Cameras),
		graphql.WithDataDir(cfg.DataDir),
	}

	gql := graphql.NewServer(
		graphql.WithResolver(
			graphql.NewResolver(resolverOpts...),
		),
	)

	routerOpts := []router.Option{
		router.WithGraphql(gql),
		router.WithFrontend(frontendFS),
		router.WithDataDir(cfg.DataDir),
		router.WithReolinks(reolinks),
	}
	router := router.NewRouter(routerOpts...)

	log.Printf("listening on %s\n", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("fatal error: %v\n", err)
	}
}
