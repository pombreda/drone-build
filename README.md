# drone-init
Drone plugin for initializing a build.


## Overview

This plugin is responsible for initializing the Drone build environment. It receives the build configuration (from the `.drone.yml`) and clone instructions and generates the build instructions.

The build instructions are generated at `/drone/bin/build.sh` which should be mounted as a volume inside the build container at runtime.

```sh
./drone-init <<EOF
{
	"clone" : {
		"branch": "master",
		"remote": "git://github.com/drone/drone",
		"dir": "/drone/src/github.com/drone/drone",
		"ref": "refs/heads/master",
		"sha": "436b7a6e2abaddfd35740527353e78a227ddcb2c"
	},
    "config": {
        "image": "golang:1.4",
        "script": [
          "go get",
          "go build",
          "go test"
        ]
    }
}
EOF
```

## Docker

Build the Docker container:

```sh
docker build -t drone/drone-init .
```

```sh
docker run -i drone/drone-init <<EOF
{
	"clone" : {
		"branch": "master",
		"remote": "git://github.com/drone/drone",
		"dir": "/drone/src/github.com/drone/drone",
		"ref": "refs/heads/master",
		"sha": "436b7a6e2abaddfd35740527353e78a227ddcb2c"
	},
	"config": {
		"image": "golang:1.4",
		"script": [
			"go get",
			"go build",
			"go test"
		]
	}
}
EOF
```
