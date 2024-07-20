FROM golang:1.21-alpine

WORKDIR /app

COPY . ./

RUN go mod download && go mod verify 


EXPOSE 9000

RUN CGO_ENABLED=0 GOOS=linux go build -o /sitenote-service


CMD ["/sitenote-service"]