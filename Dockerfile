FROM golang:alpine AS build

# GOPROXY resolves dependencies treefrom cache or repository
ENV GOPROXY=https://proxy.golang.org

WORKDIR /go/src/app
COPY . .
# Set OS as linux
RUN GOOS=linux go build -o /go/bin/app cmd/main.go

EXPOSE 8080

# Set timezone for requests creation
FROM alpine
COPY --from=build /go/bin/app /go/bin/app
COPY --from=build /go/src/app/configs /go/bin/
ENTRYPOINT ["go/bin/app"]
