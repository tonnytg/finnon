FROM golang:latest

COPY . .

EXPOSE 8089

ENTRYPOINT ["go", "run", "main.go"]