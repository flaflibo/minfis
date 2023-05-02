# Minfis a tiny file server in a docker container

Well this is a simple file webserver written in go using the Iris framework.
It can only server files, nothing else.

The main purpose to develop this was to have a docker friendly, simple configurable
microservice file server to server some static files behind a traefik reverse
proxy.

# Configuration

A yaml configuration file can be passed with the -c argument

```bash
git clone github.com/flaflibo/minfis
cd minfis

docker build -t minfis .

# Run in the background
docker run -d --restart=always \
    -p 3004:80 \
    -v $(path to yaml file)/config.yaml:/app/config.yaml \
    minfis

# Or run interactively
docker run --rm \
    -p 3004:80 \
    -v $(pwd)/config.yaml:/app/config.yaml \
    minfis

```

## Config file

```yaml
port: 80
staticRoutes:
  - path: /tmp
    route: /assets

  - path: /tmp
    route: /videos

    ....
```
