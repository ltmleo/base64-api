FROM golang:1.15 as builder

WORKDIR /app

ENV CGO_ENABLED=0 
ENV GOOS=linux 
ENV GOARCH=amd64

COPY go.mod go.sum ./
COPY pkg/ ./pkg/

RUN go mod download && go build -o ./api ./pkg/

FROM alpine:3.10

EXPOSE 8081

ENV GOTRACEBACK=single

COPY --from=builder /app/api /

CMD ["/api"]
