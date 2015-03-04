# Docker image for Drone's build script plugin
#
#     docker build --rm=true -t plugins/drone-build .

FROM library/golang:1.4

# copy the local package files to the container's workspace.
ADD . /go/src/github.com/drone-plugins/drone-build/

# build the build script plugin inside the container.
RUN go get github.com/drone-plugins/drone-build/... && \
    go install github.com/drone-plugins/drone-build

# run the git-clone plugin when the container starts
ENTRYPOINT ["/go/bin/drone-build"]
