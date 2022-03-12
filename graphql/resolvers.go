package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"time"

	"github.com/osm/reoui/graphql/generated"
	"github.com/osm/reoui/graphql/model"
)

func (r *queryResolver) Cameras(ctx context.Context) ([]*model.Camera, error) {
	var cameras []*model.Camera
	for idx, c := range r.cameras {
		cameras = append(cameras, &model.Camera{fmt.Sprintf("%d", idx), c.Name})
	}
	return cameras, nil
}

func (r *queryResolver) Videos(ctx context.Context, date *model.Date) ([]*model.Video, error) {
	videoDate := time.Now().Format("2006-01-02")
	if date != nil {
		videoDate = string(*date)
	}

	files, err := ioutil.ReadDir(path.Join(r.dataDir, videoDate[0:4], videoDate[5:7], videoDate[8:10]))
	if err != nil {
		return nil, nil
	}

	var vids []*model.Video
	for _, f := range files {
		n := f.Name()

		if f.IsDir() || !strings.HasSuffix(n, ".mp4") {
			continue
		}

		durStr := getDuration(n)
		if durStr == "" {
			// File, without duration, probably a file that is synced with
			// the built-in FTP.
			vids = append(vids, &model.Video{
				ID:         n,
				CameraName: getCameraName(n),
				Date:       getTime(n),
			})
		} else {
			dur, _ := time.ParseDuration(durStr)
			vids = append(vids, &model.Video{
				ID:         n,
				CameraName: getCameraName(n),
				Date:       getTime(n),
				Duration:   int64(dur / time.Second),
			})

		}
	}

	var ret []*model.Video
	for i := len(vids) - 1; i >= 0; i-- {
		ret = append(ret, vids[i])
	}

	return ret, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
