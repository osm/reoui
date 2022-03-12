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
