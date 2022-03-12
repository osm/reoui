package reolink

import (
	"bytes"
	"encoding/json"
	"time"
)

type File struct {
	Duration  time.Duration
	EndTime   time.Time
	Name      string
	Size      int
	StartTime time.Time
}

func (c *Client) List() ([]File, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	year, month, day := time.Now().Date()

	body, _ := json.Marshal([]Request{
		{
			Command: "Search",
			Action:  0,
			Param: Param{
				Search: &Search{
					StreamType: "main",
					StartTime: Time{
						Year:  year,
						Month: int(month),
						Day:   day,
					},
					EndTime: Time{
						Year:   year,
						Month:  int(month),
						Day:    day,
						Hour:   23,
						Minute: 59,
						Second: 59,
					},
				},
			},
		},
	})

	resp, err := postRequest(
		getURL(c.address, "/cgi-bin/api.cgi?cmd=Search&token="+token),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	var files []File
	for _, f := range resp.Value.SearchResult.File {
		startTime := toTime(&f.RawStartTime)
		endTime := toTime(&f.RawEndTime)
		files = append(files, File{
			StartTime: startTime,
			EndTime:   endTime,
			Name:      f.Name,
			Size:      f.Size,
			Duration:  endTime.Sub(startTime),
		})
	}

	return files, nil
}
