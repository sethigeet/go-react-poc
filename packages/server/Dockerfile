FROM golang:1.17-alpine as builder
RUN apk --no-cache add ca-certificates git
WORKDIR /app

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . ./
RUN CGO_ENABLED=0 go build

# Production image, copy all the files and run the server
FROM alpine
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 4000
CMD ["./app"]