FROM golang:alpine AS build

WORKDIR /build

COPY . .

RUN set +x \
&& go build -ldflags '-w -s' -o hello main.go

FROM alpine

WORKDIR /services

COPY --from=build /build/hello .

EXPOSE 80

CMD ["./hello"]
