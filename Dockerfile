# Building the binary of the App
FROM golang:1.21 AS build
ENV GOPROXY=direct
WORKDIR /go/src/tasky
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOARCH=amd64 go build -o /go/src/tasky/tasky


FROM alpine:3.17.0 AS release

WORKDIR /app
COPY --from=build  /go/src/tasky/wizexercise.txt .
COPY --from=build  /go/src/tasky/tasky .
COPY --from=build  /go/src/tasky/assets ./assets
EXPOSE 8080
ENTRYPOINT ["/app/tasky"]

