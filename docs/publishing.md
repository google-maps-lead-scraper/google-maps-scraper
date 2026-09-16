# Publishing (maintainers)

This monorepo ships two packages with aligned version numbers (currently **0.1.0**).

| Ecosystem | Package name | Source directory |
|-----------|--------------|------------------|
| npm | `@gmapsleadfinder/google-maps-scraper` | [`typescript/`](../typescript/) |
| PyPI | `google-maps-scraper-sdk` | [`python/`](../python/) |

Import / CLI reminders:

- Python import stays `gmaps_scraper` even though the PyPI name is `google-maps-scraper-sdk`.
- Both CLIs are named `gmaps-scraper`; prefer `npx` / `python -m gmaps_scraper` when testing both.

Before any release: bump versions in `typescript/package.json` and `python/pyproject.toml`, update [`CHANGELOG.md`](../CHANGELOG.md), and commit.

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

## After publish

- Confirm:
  - https://www.npmjs.com/package/@gmapsleadfinder/google-maps-scraper
  - https://pypi.org/project/google-maps-scraper-sdk/
- Smoke-test:

```bash
npm install @gmapsleadfinder/google-maps-scraper
npx gmaps-scraper me

pip install google-maps-scraper-sdk
gmaps-scraper me
```

No GitHub Actions publish workflow is configured yet; releases are manual.
