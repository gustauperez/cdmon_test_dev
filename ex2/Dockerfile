FROM golang:latest
RUN mkdir -p /app/
COPY restful /app/
WORKDIR /app
RUN go get -u github.com/gorilla/mux
RUN go build restful.go
CMD ["/app/restful"]
