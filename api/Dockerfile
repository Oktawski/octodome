FROM golang:1.25.0

WORKDIR /app

COPY ./internal ./internal

COPY api/go.mod api/go.sum ./api/

WORKDIR /app/api
RUN go mod download

COPY api/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping ./cmd/app

EXPOSE 8080

CMD ["/docker-gs-ping"]