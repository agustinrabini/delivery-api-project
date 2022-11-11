
FROM golang:1.18-alpine AS builder

ADD . /builder
WORKDIR /builder
RUN CGO_ENABLED=0 go build -o app main.go

################# Final image! #########################

FROM scratch
COPY --from=builder /builder/app ./
ENTRYPOINT ["./app"] 