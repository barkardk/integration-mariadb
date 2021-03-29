# Mariadb

Client used to run integration and e2e tests in kubernetes

## Installation

Build a test binary , compile a docker image and push to docker registry
```bash
#> make build
```
Deploy to kubernetes
```bash
#> kubectl apply -f manifests/
```
## Usage
Run locally using docker compose
```bash
#> docker-compose up
```

When deploying to kubernetes the mariadb mq client pod will run as a job, check the job logs for output
```bash
#> kubectl logs -l app=mariadb-client
```

## Parameters
|   Parameter | Default   |
|---|---|
| TAG  |  git-rev parse HEAD --short |
| DOCKER_REGISTRY | ghcr.io/barkardk  |
| MARIADB_ROOT_USER | root | 
| MARIADB_ROOT_PASSWORD | secret | 
| MARIADB_HOST | 127.0.0.1 | 
| MARIADB_CLIENT_PORT | 3306 |  