# kumo

kumo is a Kubernetes controller that manages cloud resources.

> Fork of [sivchari/kumo](https://github.com/sivchari/kumo)

## Prerequisites

- Go 1.21+
- Kubernetes 1.27+
- Helm 3.x

## Installation

### Helm

```bash
helm install kumo ./charts/kumo \
  --namespace kumo-system \
  --create-namespace
```

## Development

### Setup

```bash
git clone https://github.com/your-org/kumo.git
cd kumo
go mod download
```

### Running Tests

```bash
# Unit tests
make test

# Integration tests
make integration-test

# E2E tests (requires a running cluster)
# Note: set KUBECONFIG env var before running
# Tip: use `kind create cluster` for a quick local cluster
make e2e-test
```

### Linting

```bash
make lint
```

### Local Development Tips

- Use `make run` to run the controller locally against a kind cluster
- Set `LOG_LEVEL=debug` for verbose output during development
- Use `kind create cluster --name kumo-dev` to keep the dev cluster separate from others

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

Apache 2.0
