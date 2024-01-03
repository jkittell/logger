FROM golang:latest as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY *.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o /logger

FROM busybox
COPY .env /home
COPY --from=builder /logger /home
EXPOSE 80
WORKDIR /home
ENTRYPOINT [ "./logger" ]
