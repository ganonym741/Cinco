FROM golang:alpine3.16
LABEL NAME : Cinco kelompok2

RUN mkdir $GOPATH/src/gitlab.com/
RUN mkdir $GOPATH/src/gitlab.com/cinco
WORKDIR $GOPATH/src/gitlab.com/cinco
COPY . $GOPATH/src/gitlab.com/cinco

RUN go mod download && go mod verify

EXPOSE 8000

CMD ["go","run", "main.go"]