### Build binary from official Go image
FROM golang:stretch as build
COPY . /app
WORKDIR /app
RUN go build -o /go-heroku .

### Put the binary onto Heroku image
FROM heroku/heroku:16
COPY --from=build /go-heroku /go-heroku
CMD ["/go-heroku"]