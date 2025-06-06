ARG GO_VERSION=1.24.2
FROM golang:${GO_VERSION}-bookworm as builder


WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /anti-judol-regex .


FROM debian:bookworm
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=builder /anti-judol-regex /usr/local/bin/
CMD ["anti-judol-regex"]