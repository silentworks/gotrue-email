FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
COPY templates ./templates
RUN go mod download
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/gotrue-email

FROM scratch
WORKDIR /app
COPY --from=build /opt/gotrue-email /app/gotrue-email
COPY --from=build /app/templates /app/templates
EXPOSE 8088
ENTRYPOINT ["/app/gotrue-email"]
