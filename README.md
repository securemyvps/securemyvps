<p align="center">
  <img src="docs/assets/banner-readme-1280x320.png" alt="SecureMyVPS — open-source VPS security and hardening" width="100%">
</p>

<h1 align="center">securemyvps</h1>

<p align="center">
  <a href="https://github.com/securemyvps/securemyvps/actions/workflows/ci.yml"><img src="https://github.com/securemyvps/securemyvps/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://github.com/securemyvps/securemyvps/releases/latest"><img src="https://img.shields.io/github/v/release/securemyvps/securemyvps?color=4ade80" alt="Latest release"></a>
  <a href="go.mod"><img src="https://img.shields.io/github/go-mod/go-version/securemyvps/securemyvps" alt="Go version"></a>
  <a href="LICENSE"><img src="https://img.shields.io/github/license/securemyvps/securemyvps" alt="License: MIT"></a>
</p>

<p align="center">
  VPS hardening with lockout-safe rollback.<br>
  Go, <code>CGO_ENABLED=0</code> static binaries, no telemetry.
</p>

---

## Status

**Phase 0 — stub CLI.** The binary accepts `--version` and exits. The remediation engine lands in `v0.1`; there is nothing to run against a live server yet.

| Area | State |
|---|---|
| CLI | `--version` only |
| Audit engine | Not started — imports `github.com/vpsinspect/vpsinspect` once that library ships |
| CI | Build, test, golangci-lint, govulncheck, zizmor |
| Supply chain | Reproducible builds, SHA-256 checksums, GitHub build attestations |
| Platform | Linux now (amd64, arm64); Windows Server planned for Phase 3 |

## Build (local)

Requires Go 1.26 or newer.

```bash
make build test
./securemyvps --version
```

## Verify a release

```bash
gh attestation verify dist/securemyvps_* -R securemyvps/securemyvps
sha256sum -c SHA256SUMS.txt
```

## Project structure

```text
cmd/securemyvps/    CLI entrypoint
.github/workflows/  CI, release, and Scorecard workflows
docs/assets/        Banner and social preview sources
Makefile            build, test, vendor targets
```

## Media

| Asset | Preview |
|---|---|
| README hero | [banner-readme-1280x320.png](docs/assets/banner-readme-1280x320.png) · [SVG](docs/assets/banner-readme-1280x320.svg) |
| Social preview | [social-preview-1280x640.png](docs/assets/social-preview-1280x640.png) · [SVG](docs/assets/social-preview-1280x640.svg) |

## License

[MIT](LICENSE)
