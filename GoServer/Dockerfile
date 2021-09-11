FROM golang:1.17.0-alpine3.14
WORKDIR /app
COPY go.mod ./
# COPY go.sum ./
# RUN go mod download
COPY *.go ./
RUN go build -o ./GoServer
RUN chmod +x ./GoServer
EXPOSE 8081
CMD [ "./GoServer" ]