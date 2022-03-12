package clean

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

type clean struct {
	dataDir            string
	cleanFilesInterval time.Duration
}

func New(opts ...Option) *clean {
	c := &clean{}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *clean) visit(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	// Iterate over each file.
	// If it's a dir, visit it.
	// If the file is older than c.cleanFilesInterval, remove it.
	for _, f := range files {
		filename := f.Name()
		filePath := path.Join(dir, filename)

		if f.IsDir() {
			c.visit(filePath)
		}

		if time.Now().Sub(f.ModTime()) > c.cleanFilesInterval {
			log.Printf("removing %s\n", filePath)
			os.Remove(filePath)
		}
	}

	// If we aren't in the root dir, let's list again to check whether or
	// not we should remove the dir we are in or not.
	if dir != c.dataDir {
		files, err = ioutil.ReadDir(dir)
		if err != nil {
			return
		}
		if len(files) == 0 {
			log.Printf("removing %s\n", dir)
			os.Remove(dir)
		}
	}
}

func (c *clean) Run() {
	for {
		log.Printf("running cleaner...\n")
		c.visit(c.dataDir)
		log.Printf("cleaning done\n")

		time.Sleep(1 * time.Minute)
	}
}
