FROM golang:alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY *.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o /logger

FROM busybox
COPY --from=builder /logger /home
WORKDIR /home
ENTRYPOINT [ "./logger" ]
