FROM golang:1.16
WORKDIR /go/src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# RUN go install github.com/cosmtrek/air@v1.27.3 && \
#   go build -o /go/bin/air github.com/cosmtrek/air
CMD ["air"]

# RUN export GOPATH=$HOME/go
# RUN export PATH=$PATH:$GOPATH/bin
# CMD ["air"]