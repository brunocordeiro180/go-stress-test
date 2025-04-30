FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod tidy && \
    go build -o stresstest .

FROM scratch

COPY --from=builder /app/stresstest /stresstest

ENTRYPOINT [ "/stresstest" ]