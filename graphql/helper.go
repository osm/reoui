package graphql

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/alfg/mp4"
)

var re = regexp.MustCompile(`[0-9]{14}`)

func getTime(n string) time.Time {
	d := re.FindString(n)
	year := d[0:4]
	month := d[4:6]
	day := d[6:8]
	hour := d[8:10]
	minute := d[10:12]
	second := d[12:14]
	t, _ := time.Parse(
		time.RFC3339,
		fmt.Sprintf("%s-%s-%sT%s:%s:%sZ", year, month, day, hour, minute, second),
	)
	return t
}

func getCameraName(n string) string {
	idx := re.FindStringIndex(n)
	return strings.ReplaceAll(n[0:idx[0]-1], "_", " ")
}

func getDuration(filePath string) time.Duration {
	var dur time.Duration

	file, err := os.Open(filePath)
	if err != nil {
		return dur
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return dur
	}

	mp4, err := mp4.OpenFromReader(file, info.Size())
	if mp4.Moov == nil {
		return dur
	}

	ms := int64(mp4.Moov.Mvhd.Duration)
	dur = time.Duration(ms * 1000000)
	return dur.Round(time.Second)
}
