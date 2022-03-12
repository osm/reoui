package sync

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
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

				filenameStart := fmt.Sprintf("%s_%s%s%s%s", cameraName, year, month, day, time)
				fileMatch, err := filepath.Glob(
					path.Join(dir, fmt.Sprintf("%s_*.mp4", filenameStart)),
				)
				if err != nil {
					log.Printf("sync glob: %v\n", err)
					continue
				}
				if len(fileMatch) > 1 {
					log.Printf("sync glob, too many files matches: %v\n", fileMatch)
					continue
				}

				filename := fmt.Sprintf("%s_%s.mp4", filenameStart, f.Duration)
				if len(fileMatch) == 1 {
					existingFile, _ := os.Stat(fileMatch[0])
					size := existingFile.Size()

					if size == int64(f.Size) {
						continue
					} else if int64(f.Size) > size {
						os.Remove(fileMatch[0])
					}
				}

				filePath := path.Join(dir, filename)
				file, err := os.Create(filePath)
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
