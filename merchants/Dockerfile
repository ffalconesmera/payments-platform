# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS build

# Create work directory
WORKDIR /app

RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache

# Copy API packages
COPY . ./

RUN --mount=type=cache,target=/gomod-cache \
    go mod download

RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache \
    go build -o /merchants main.go

FROM alpine:3.19

COPY --from=build /merchants ./.
COPY .env.example .env

EXPOSE 8081

CMD [ "/merchants" ]