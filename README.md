# drone-build
Drone plugin for initializing a build.


## Overview

This plugin is responsible for initializing the Drone build environment. It receives the build configuration (from the `.drone.yml`) and clone instructions and generates the build instructions.

The build instructions are generated at `/drone/bin/build.sh` which should be mounted as a volume inside the build container at runtime.

```sh
./drone-build <<EOF
{
	"clone" : {
		"branch": "master",
		"remote": "git://github.com/drone/drone",
		"dir": "/drone/src/github.com/drone/drone",
		"ref": "refs/heads/master",
		"sha": "436b7a6e2abaddfd35740527353e78a227ddcb2c"
	},
    "vargs": {
        "commands": [
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
docker build --rm=true -t plugins/drone-build .
```

Run the init step to generate the build script:

```sh
docker run -i plugins/drone-build <<EOF
{
	"clone" : {
		"branch": "master",
		"remote": "git://github.com/drone/drone",
		"dir": "/drone/src/github.com/drone/drone",
		"ref": "refs/heads/master",
		"sha": "436b7a6e2abaddfd35740527353e78a227ddcb2c"
	},
	"vargs": {
		"commands": [
			"go get",
			"go build",
			"go test"
		]
	}
}
EOF
```

Note that Drone will create a volume to share `/drone/bin` with subsequnt containers, including the build container specificied in the `image` attribute of the `.drone.yml` file:

```
docker run -v /drone/bin -i plugins/drone-build
```