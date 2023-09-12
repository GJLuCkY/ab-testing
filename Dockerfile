FROM golang:1.20-alpine

WORKDIR /src

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o ./server ./cmd/app/main.go

EXPOSE 8030

CMD ["./server"]
