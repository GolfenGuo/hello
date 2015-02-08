FROM google/golang
MAINTAINER Golfen Guo "golfen.guo@daocloud.io"

# Build app
WORKDIR /gopath/src/app
ADD . /gopath/src/app/
RUN go get app

# nginx will bind to port 80, exposing API
EXPOSE 80
CMD ["/gopath/bin/app"]