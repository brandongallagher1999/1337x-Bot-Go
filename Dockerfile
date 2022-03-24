FROM golang:1.16-buster AS build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /main cmd/1337x-Bot-Go/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /main /docker-go-main
COPY --from=build /app/config /config

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-go-main"]