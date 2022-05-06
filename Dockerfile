## STEP 1 - BUILD
FROM golang:1.17.9-alpine AS build

# RUN apk upgrade --update && \
#     apk --no-cache add git gcc musl-dev

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . /app/

RUN go build -o /app/main

## STEP 2 - DEPLOY
FROM alpine

WORKDIR /app
COPY --from=build /app/main /app/main
COPY .env /app/

EXPOSE 8080
ENTRYPOINT ["/app/main"]
