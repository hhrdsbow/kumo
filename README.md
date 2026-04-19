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
- Use `kubectl logs -n kumo-system deploy/kumo -f` to tail controller logs in real time
- Use `kubectl get events -n kumo-system --sort-by='.lastTimestamp'` to debug reconciliation issues

### My Notes

- I'm using this to learn controller-runtime patterns; the reconciler loop in `internal/controller/` is the main thing to study
- `make run` requires the CRDs to be installed first: `make install` before `make run`
- Useful shortcut: `make install run` chains both commands so I don't forget to install CRDs first
- When iterating quickly, `make install run LOG_LEVEL=debug` is my go-to command

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

Apache 2.0
