FROM golang:1.22.2-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd
RUN chmod +x main
EXPOSE 5050
CMD [ "./main" ]