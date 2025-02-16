FROM golang:1.23.2

WORKDIR /api-email-indexer

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o ./out/api-email-indexer /api-email-indexer/

EXPOSE ${HTTP_PORT}
EXPOSE 6060

CMD "./out/api-email-indexer"
