FROM golang:latest
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
ADD . /opt/app
WORKDIR /opt/app
EXPOSE 5000
ENTRYPOINT CompileDaemon -build="go build -o start ." -directory="." -log-prefix=false -command="./start"
