
FROM golang:1.18-alpine AS builder

WORKDIR /builder
ADD . /builder
RUN CGO_ENABLED=0 go build -mod=vendor

################# Final image! #########################

FROM scratch
COPY --from=builder /builder/delivery-api-project /app
ENTRYPOINT ["/app"] 