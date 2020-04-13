FROM golang:latest
LABEL maintainer="KlexHub UG (haftungsbeschränkt) <support@support-pp.de>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
CMD ["./main"]

