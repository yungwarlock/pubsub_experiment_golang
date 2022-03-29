FROM golang:1.18-buster as build

WORKDIR /build
COPY . .
RUN go build -o subscriber ./cmd/subscriber

FROM gcr.io/distroless/base-debian10

COPY --from=build /build/subscriber /go/bin/app

ENTRYPOINT [ "/go/bin/app" ]
