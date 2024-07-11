FROM golang:1.18 as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /main ./cmd

FROM gcr.io/distroless/base-debian10

COPY --from=build /main /main

EXPOSE 8080

ENTRYPOINT ["/main"]