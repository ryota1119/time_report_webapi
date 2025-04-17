FROM golang:1.24
WORKDIR /app
RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

CMD ["air","-c",".air.toml"]