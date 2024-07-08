FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify 

COPY *.go ./

EXPOSE 9000

RUN CGO_ENABLED=0 GOOS=linux go build -o /notelad-service


CMD ["/notelad-service"]