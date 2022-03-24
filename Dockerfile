FROM golang:1.16-buster AS build

WORKDIR /app

COPY . .

RUN go build cmd/1337x-Bot-Go/main.go -o main

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /main /docker-go-main

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-go-main"]