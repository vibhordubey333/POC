#FROM golang:1.21.3-bullseye AS build
#
#WORKDIR /app
#
#COPY go.mod ./
#COPY go.sum ./
#
#RUN go mod tidy
#
#COPY * ./
#
#RUN GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/server.go
#
#FROM gcr.io/distroless/base-debian11
#
#COPY --from=build /cmd /cmd
#
#CMD ["/server"]

FROM golang:1.21.3-bullseye AS build
RUN mkdir /app
WORKDIR /app

COPY . .
RUN chmod 0700 .
RUN export GOSUMDB=off
RUN export GO111MODULE=on
RUN go mod tidy
RUN GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/server.go
RUN chmod +x ./cmd/server
CMD ["./cmd/server"]