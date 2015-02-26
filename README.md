# drone-init
Drone plugin for initializing a build.

```sh
./drone-init <<EOF
{
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

