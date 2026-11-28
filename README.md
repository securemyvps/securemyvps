<p align="center">
  <img src="docs/assets/banner-readme-1280x320.png" alt="SecureMyVPS" width="100%">
</p>

<h1 align="center">securemyvps</h1>

<p align="center">
  VPS hardening that will not lock you out
</p>

<p align="center">
  <a href="https://github.com/securemyvps/securemyvps/actions/workflows/ci.yml"><img src="https://github.com/securemyvps/securemyvps/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://github.com/securemyvps/securemyvps/releases/latest"><img src="https://img.shields.io/github/v/release/securemyvps/securemyvps?color=4ade80" alt="Latest release"></a>
  <a href="LICENSE"><img src="https://img.shields.io/github/license/securemyvps/securemyvps" alt="License: MIT"></a>
</p>

---

## Status

**Phase 0 — stub CLI.** The binary accepts `--version` and exits. The remediation engine lands in `v0.1`; there is nothing to run against a live server yet.

## Build

```bash
make build test
./securemyvps --version
```

## Verify a release

```bash
gh attestation verify dist/securemyvps_* -R securemyvps/securemyvps
sha256sum -c SHA256SUMS.txt
```

## License

[MIT](LICENSE)
