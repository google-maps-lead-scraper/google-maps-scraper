# Publishing (maintainers)

This monorepo ships three client packages. Keep versions aligned when releasing together (currently **0.1.1**).

| Ecosystem | Package / module | Source directory |
|-----------|------------------|------------------|
| npm | `@gmapsleadfinder/google-maps-scraper` | [`typescript/`](../typescript/) |
| PyPI | `google-maps-scraper-sdk` | [`python/`](../python/) |
| Go | `github.com/google-maps-lead-scraper/google-maps-scraper/go` | [`go/`](../go/) |

Import / CLI reminders:

- Python import stays `gmaps_scraper` even though the PyPI name is `google-maps-scraper-sdk`.
- Go import alias is typically `gmaps`.
- Python and npm CLIs are both named `gmaps-scraper`; prefer `npx` / `python -m gmaps_scraper` / `go run ./cli`.

Before any release: bump versions in `typescript/package.json`, `python/pyproject.toml`, and Go user-agent strings; update [`CHANGELOG.md`](../CHANGELOG.md); commit and push.

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
```

Optional TestPyPI first:

```bash
twine upload --repository testpypi dist/*
pip install -i https://test.pypi.org/simple/ google-maps-scraper-sdk
```

Production:

```bash
twine upload dist/*
# Username: __token__
# Password: pypi-... (the API token)
```

## Go ([pkg.go.dev](https://pkg.go.dev/))

Go has **no** package upload. Publish by pushing a **subdirectory module tag**:

```bash
# After merging to main and pushing commits:
git tag go/v0.1.1
git push origin go/v0.1.1

# Trigger the module proxy (optional but recommended):
GOPROXY=https://proxy.golang.org go list -m github.com/google-maps-lead-scraper/google-maps-scraper/go@v0.1.1
```

Then open:

https://pkg.go.dev/github.com/google-maps-lead-scraper/google-maps-scraper/go@v0.1.1

If the page is missing, wait a few minutes or use the site’s request/index flow. Tag format for this monorepo subdirectory module must be `go/vX.Y.Z` (not a root-only `vX.Y.Z`).

Install for users:

```bash
go get github.com/google-maps-lead-scraper/google-maps-scraper/go@v0.1.1
```

## After publish

- Confirm:
  - https://www.npmjs.com/package/@gmapsleadfinder/google-maps-scraper
  - https://pypi.org/project/google-maps-scraper-sdk/
  - https://pkg.go.dev/github.com/google-maps-lead-scraper/google-maps-scraper/go
- Smoke-test:

```bash
npm install @gmapsleadfinder/google-maps-scraper
npx gmaps-scraper me

pip install google-maps-scraper-sdk
gmaps-scraper me

go get github.com/google-maps-lead-scraper/google-maps-scraper/go@v0.1.1
```

No GitHub Actions publish workflow is configured yet; releases are manual.
