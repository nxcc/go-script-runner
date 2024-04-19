# Go Script Runner

Experimental go script runner, based on [yaegi](https://github.com/traefik/yaegi).

Build static `gsr` with dependencies compiled in:

     CGO_ENABLED=0 go build ./cmd/gsr


Run scripts standalone, no go installation or dependencies required:

    ./gsr ./scripts/simple/
    ./gsr -run Resty ./scripts/simple/
