# securemyvps

Linux VPS hardening with lockout-safe rollback (Go, `CGO_ENABLED=0` static binaries). Imports `github.com/vpsinspect/vpsinspect` once the audit library ships.

- **Module path:** `github.com/securemyvps/securemyvps`
- **License:** MIT
- **Status:** Phase 0 — stub CLI (`--version`); remediation engine lands in v0.1.

## Build (local)

```bash
make build test
./securemyvps --version
```