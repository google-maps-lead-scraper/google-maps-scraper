# Publishing (maintainers)

This monorepo ships four client packages. Keep versions aligned when releasing together (currently **0.1.1**).

| Ecosystem | Package / module | Source directory |
|-----------|------------------|------------------|
| npm | `@gmapsleadfinder/google-maps-scraper` | [`typescript/`](../typescript/) |
| PyPI | `google-maps-scraper-sdk` | [`python/`](../python/) |
| Go | `github.com/google-maps-lead-scraper/google-maps-scraper/go` | [`go/`](../go/) |
| crates.io | `google-maps-scraper-sdk` | [`rust/`](../rust/) |

Import / CLI reminders:

- Python import stays `gmaps_scraper` even though the PyPI name is `google-maps-scraper-sdk`.
- Rust crate is also `google-maps-scraper-sdk` (Rust path `google_maps_scraper_sdk`).
- Go import alias is typically `gmaps`.
- Python / npm / Rust CLIs may share the name `gmaps-scraper`; prefer `npx` / `python -m` / `cargo run --bin gmaps-scraper` / `go run ./cli`.

Before any release: bump versions in `typescript/package.json`, `python/pyproject.toml`, `rust/Cargo.toml`, and Go/Rust user-agent strings; update [`CHANGELOG.md`](../CHANGELOG.md); commit and push.

## npm (`@gmapsleadfinder/google-maps-scraper`)

Use your **personal** `npm login`. The Org is a publish namespace, not a separate login.

```bash
cd typescript
npm login
npm whoami
npm install
npm run build
npm pack --dry-run
npm publish --access public
```

Notes:

- Scoped packages need `publishConfig.access: "public"` (already set) or `--access public`.
- You must be Owner/Admin of the `gmapsleadfinder` npm org.
- `prepublishOnly` runs `npm run build` automatically on publish.

## PyPI (`google-maps-scraper-sdk`)

1. Register / verify email at https://pypi.org/account/register/
2. Enable 2FA and create an **API token** (Account settings → API tokens).
3. Build and upload:

```bash
cd python
python -m pip install build twine
rm -rf dist build *.egg-info src/*.egg-info
python -m build
twine check dist/*
twine upload dist/*
# Username: __token__
# Password: pypi-... (the API token)
```

## Go ([pkg.go.dev](https://pkg.go.dev/))

Go has **no** package upload. Publish by pushing a **subdirectory module tag**:

```bash
git tag go/v0.1.1
git push origin go/v0.1.1
GOPROXY=https://proxy.golang.org go list -m github.com/google-maps-lead-scraper/google-maps-scraper/go@v0.1.1
```

Library install (inside a module):

```bash
go mod init example.com/app
go get github.com/google-maps-lead-scraper/google-maps-scraper/go@v0.1.1
```

Do **not** `go install` the library path (`package gmaps` is not `main`). Use `go run ./cli` from the repo for the CLI.

## crates.io (`google-maps-scraper-sdk`)

1. Register at https://crates.io (GitHub login recommended).
2. Account Settings → API Tokens → New token.
3. Publish:

```bash
cargo login
cd rust
cargo check
cargo publish --dry-run
cargo publish
```

Users:

```bash
cargo add google-maps-scraper-sdk
cargo install google-maps-scraper-sdk --bin gmaps-scraper
```

docs.rs builds automatically after publish: https://docs.rs/google-maps-scraper-sdk

## After publish

- Confirm:
  - https://www.npmjs.com/package/@gmapsleadfinder/google-maps-scraper
  - https://pypi.org/project/google-maps-scraper-sdk/
  - https://pkg.go.dev/github.com/google-maps-lead-scraper/google-maps-scraper/go
  - https://crates.io/crates/google-maps-scraper-sdk

No GitHub Actions publish workflow is configured yet; releases are manual.
