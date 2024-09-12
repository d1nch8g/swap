FROM golang:1.23-alpine as builder

WORKDIR /service

# Creates non root user
ENV USER=appuser
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service


FROM alpine

ENV PORT=8080

WORKDIR /app

# Non root user info
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Certs for making https requests
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /service/service /app/service
COPY --from=builder /service/db/migrations /app/db/migrations
COPY --from=builder /service/dist /app/dist

# Running as appuser
USER appuser:appuser

EXPOSE ${PORT}
ENTRYPOINT ["/app/service"]