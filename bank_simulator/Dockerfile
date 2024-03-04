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
    go build -o /bank_simulator main.go

FROM alpine:3.19

COPY --from=build /bank_simulator ./.

EXPOSE 8083

CMD [ "/bank_simulator" ]
