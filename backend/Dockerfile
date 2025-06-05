FROM golang:1.24.2 AS builder

WORKDIR /app

ARG DSN
ARG PORT
ARG JWT_SECRET
ARG FRONTEND_URL

ENV DSN=${DSN}
ENV PORT=${PORT}
ENV JWT_SECRET=${JWT_SECRET}
ENV FRONTEND_URL=${FRONTEND_URL}

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main /main
CMD ["/main"]