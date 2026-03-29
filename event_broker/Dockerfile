FROM golang:1.26.0

WORKDIR /app

COPY ./internal ./internal

COPY event_broker/go.mod event_broker/go.sum ./event_broker/

WORKDIR /app/event_broker
RUN go mod download

COPY event_broker/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping ./cmd/app

CMD ["/docker-gs-ping"]