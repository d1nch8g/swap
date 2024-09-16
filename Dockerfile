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

FROM node as node

WORKDIR /app

COPY package.json package.json
COPY package-lock.json package-lock.json

RUN npm install

COPY . .

RUN npm run build

FROM alpine

ENV PORT=8080

# Non root user info
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Certs for making https requests
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /service/service /service
COPY --from=builder /service/db/migrations /db/migrations
COPY --from=node /app/dist /dist

# Running as appuser
USER appuser:appuser

EXPOSE ${PORT}
ENTRYPOINT ["/service"]