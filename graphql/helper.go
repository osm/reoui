package graphql

import (
	"fmt"
	"regexp"
	"strings"
	"time"
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

func getDuration(n string) string {
	dateIdx := re.FindStringIndex(n)
	endingIdx := strings.LastIndex(n, ".")

	if dateIdx[1] == endingIdx {
		return ""
	}
	return n[dateIdx[1]+1 : endingIdx]
}
