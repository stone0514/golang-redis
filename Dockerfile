FROM golang:1.17

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# add air to GOPATH
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# hot reload
CMD ["air"]
