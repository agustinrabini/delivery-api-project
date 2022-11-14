
FROM golang:1.18-alpine AS builder

ADD . /builder
WORKDIR /builder
RUN CGO_ENABLED=0 go build -o app

################# Final image! #########################

FROM docker.io/library/alpine:latest
COPY --from=builder /builder/app ./
COPY --from=builder /builder/config/config.yaml config.yaml
ENTRYPOINT ["./app"] 