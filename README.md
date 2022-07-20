# wait-dockerd-startup

Wait until docker API(docker daemon) startup. 

## Install

```bash
go get github.com/nncdevel-io/wait-dockerd-startup
```

## Usage

```bash
$ wait-dockerd-startup --help
Wait until docker daemon startup.

Usage:
  wait-dockerd [flags]

Flags:
  -f, --failure-threshold int    Threshold value for detect Failure. (default 10)
  -h, --help                     help for wait-dockerd
  -i, --initial-delay duration   Initial delay (default 10s)
  -p, --period duration          Check period. (default 10s)
  -s, --success-threshold int    Threshold value for detect Succeed. (default 1)
  -t, --timeout duration         Docker API request timeout. (default 1s)
```


## How It Works

1. Sleep `--interval` setting duration.
2. Send request to docker [SystemInfo API](https://docs.docker.com/engine/api/v1.41/#tag/System/operation/SystemInfo) and 
parse server version information.
3. If got server version info, increment success count. On the other hand, increment failure count.
4. If success/failure count greater than or equal to threshold options, the program will exit 0 or 1.
