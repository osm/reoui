package sync

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/osm/reoui/reolink"
)

type sync struct {
	dataDir      string
	syncInterval time.Duration
	reolinks     []*reolink.Client
}

func New(opts ...Option) *sync {
	s := &sync{}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *sync) Run() {
	for {
		for _, c := range s.reolinks {
			log.Printf("syncing %s\n", c.Name)
			files, err := c.List()
			if err != nil {
				log.Printf("reolink list error: %v\n", err)
				continue
			}

			cameraName := strings.ReplaceAll(c.Name, " ", "_")

			for _, f := range files {
				year := f.Name[4:8]
				month := f.Name[8:10]
				day := f.Name[10:12]
				time := f.Name[13:19]

				dir := path.Join(s.dataDir, year, month, day)
				err := os.MkdirAll(dir, os.ModePerm)
				if err != nil {
					log.Printf("sync mkdir: %v\n", err)
					continue
				}

				filePath := path.Join(
					dir,
					fmt.Sprintf("%s_%s%s%s%s.mp4", cameraName, year, month, day, time),
				)
				if existingFile, err := os.Stat(filePath); err == nil {
					size := existingFile.Size()

					if size == int64(f.Size) {
						continue
					} else if int64(f.Size) > size {
						os.Remove(filePath)
					}
				}

				file, err := os.Create(filePath)
				defer file.Close()
				if err != nil {
					log.Printf("create file: %v\n", err)
					continue
				}

				log.Printf("downloading %s (%d bytes) to %s\n", f.Name, f.Size, filePath)
				c.Download(f.Name, file)
			}
		}

		log.Printf("syncing done\n")
		time.Sleep(s.syncInterval)
	}
}
