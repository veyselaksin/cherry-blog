FROM golang:1.20-bullseye as builder

WORKDIR /app

COPY go.sum go.mod ./

RUN apt-get update \
        && apt-get upgrade -y \
        && apt-get install -y \
        ca-certificates

RUN go mod download

COPY . .

RUN mkdir -p media/xlsx
ENV TZ=Europe/Istanbul

RUN CGO_ENABLED=0 go build -tags timetzdata  -o /bin/app ./cmd/

ENTRYPOINT ["/bin/app"]

# FROM scratch

# WORKDIR /app

# ENV TZ=Europe/Istanbul
# COPY --from=builder /bin/app /bin/app
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /app/media /app/media

# ENTRYPOINT ["/bin/app"]

