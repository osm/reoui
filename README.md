# reoui

Reolink UI.

## Backend development

```sh
$ make frontend
$ make
```

## Frontend development

Requires the backend to be started for the API to work.

```sh
$ GRAPHQL_URL=http://localhost:4000/graphql yarn start:dev
```

## Docker

```sh
$ docker build -t reoui:latest .
$ docker run \
	-p 4000:4000 \
	-v /your/data/dir:/var/lib/reoui \
	-v /your/config/dir:/etc/reoui \
	reoui:latest -config /etc/reoui/reoui.yaml
```

## Configuration in environment

```sh
$ export REOUI_PORT=4000
$ export REOUI_DATA_DIR=/tmp/reoui
$ export REOUI_SYNC_INTERVAL=1m
$ export REOUI_CLEAN_FILES_INTERVAL=72h
$ export REOUI_CAMERA_0_NAME=cam-from-env
$ export REOUI_CAMERA_0_ADDRESS=http://192.168.0.101
$ export REOUI_CAMERA_0_USERNAME=username
$ export REOUI_CAMERA_0_PASSWORD=top_secret_password
$ export REOUI_CAMERA_0_LOW_STREAM_QUALITY=true
$ reoui
```
