FROM golang:1.19.3-alpine

WORKDIR /back

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]