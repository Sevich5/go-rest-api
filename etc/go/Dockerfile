FROM golang:latest

RUN mkdir /app
WORKDIR /app
ADD . /app
COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
RUN go install ariga.io/atlas/cmd/atlas@latest

ENTRYPOINT CompileDaemon --build="go build -o build/api cmd/api/main.go" --command=./build/api