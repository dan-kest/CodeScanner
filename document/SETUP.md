# Setup

## Prerequisites

- Go
- golangci-lint
- Docker

## Steps

1. Create `.env` file for docker-compose in `{{project_root}}/docker/` with the following content:
```
POSTGRES_DATABASE=db
POSTGRES_USERNAME=root
POSTGRES_PASSWORD=password
```
2. Edit some of volumns path in `{{project_root}}/docker/docker-compose.yml` to fits your local machine
3. Run `docker-compose up -d` in `{{project_root}}/docker/`
4. Create `secret.yaml` file for application config in `{{project_root}}/config/` with the following content:
```yaml
postgres:
  host: localhost
  port: 5432
  database: db
  username: root
  password: password

rabbitmq:
  host: localhost
  port: 5672
  username: guest
  password: guest
```
5. Connect to PostgreSQL and run [schema.sql](schema.sql)
6. Start both  `{{project_root}}/cmd/http/main.go` and `{{project_root}}/cmd/consumer/main.go`
7. Edit important config in [config.yaml](../config/config.yaml)
8. Done! see [API Documentation](APIDOC.md) for more action
