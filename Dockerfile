FROM golang:1.15.2-alpine AS GO_BUILD
COPY server /server
WORKDIR /server
RUN apk add build-base
RUN go build -o /go/bin/server

FROM alpine:3.12.0
COPY server/.env ./
COPY server/public ./public
COPY server/views ./views

COPY --from=GO_BUILD /go/bin/server ./
CMD ./server