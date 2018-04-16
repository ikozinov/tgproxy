# Stage 1 - the build process
FROM golang:1.9 as build-deps
ENV GOPATH /usr

COPY ./src /usr/src/build
WORKDIR /usr/src/build
RUN go get -t -v ./...

RUN go build -ldflags "-linkmode external -extldflags -static" 

# Stage 2 - the production environment
FROM alpine
LABEL maintainer="Igor V. Kozinov <madvampik@gmail.com>"

COPY --from=build-deps /usr/src/build/build /bin/tgproxy

EXPOSE 6785
ENTRYPOINT ["/bin/tgproxy"]