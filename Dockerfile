FROM golang:1.10

WORKDIR $GOPATH/src/adrequestauctionsystem

COPY . .

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN dep ensure -v

RUN go get

RUN go install

EXPOSE 8081

CMD ["adrequestauctionsystem"]
