# Docker image for Drone's init plugin
#
#     docker build -t drone/drone-init .

FROM library/golang:1.4

# copy the local package files to the container's workspace.
ADD . /go/src/github.com/drone/drone-init/

# build the git-clone plugin inside the container.
RUN go get github.com/drone/drone-init/... && \
    go install github.com/drone/drone-init

# run the git-clone plugin when the container starts
ENTRYPOINT /go/bin/drone-init
