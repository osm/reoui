FROM mhart/alpine-node:14
WORKDIR /usr/src/app
COPY frontend .
RUN apk add git
RUN yarn
RUN yarn build

FROM golang:1.17-alpine
WORKDIR /go/src/app
COPY . .
RUN mkdir -p /go/src/app/frontend/dist
COPY --from=0 /usr/src/app/dist/* /go/src/app/frontend/dist/
RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine:3.15.0
COPY --from=1 /go/bin/reoui /usr/bin
ENTRYPOINT ["reoui"]
