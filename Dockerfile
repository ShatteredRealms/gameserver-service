# Build application
FROM golang:1.24 AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ARG APP_VERSION=v0.0.1
RUN go build \
	-ldflags="-X 'github.com/ShatteredRealms/gameserver-service/pkg/config/default.Version=${APP_VERSION}'" \
	-o /out/gameserver-service ./cmd/gameserver-service

# Run server
FROM alpine:3.21.3
WORKDIR /app
COPY --from=build /out/gameserver-service ./
EXPOSE 8082
ENTRYPOINT [ "./gameserver-service" ]
