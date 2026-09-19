# Publishing (maintainers)

This monorepo ships six client packages; **PHP** lives in a separate repo. Keep versions aligned when releasing together (currently **0.1.2**).

| Ecosystem | Package / module | Source directory |
|-----------|------------------|------------------|
| npm | `@gmapsleadfinder/google-maps-scraper` | [`typescript/`](../typescript/) |
| PyPI | `google-maps-scraper-sdk` | [`python/`](../python/) |
| Go | `github.com/google-maps-lead-scraper/google-maps-scraper/go` | [`go/`](../go/) |
| crates.io | `google-maps-scraper-sdk` | [`rust/`](../rust/) |
| RubyGems | `google-maps-scraper-sdk` | [`ruby/`](../ruby/) |
| NuGet | `GmapsLeadFinder.GoogleMapsScraper` (+ `.Cli` tool) | [`dotnet/`](../dotnet/) |
| Packagist | `gmapsleadfinder/google-maps-scraper` | [google-maps-scraper-php](https://github.com/google-maps-lead-scraper/google-maps-scraper-php) (separate repo) |

Import / CLI reminders:

- Python import stays `gmaps_scraper` even though the PyPI name is `google-maps-scraper-sdk`.
- Rust crate is also `google-maps-scraper-sdk` (Rust path `google_maps_scraper_sdk`).
- Ruby gem is `google-maps-scraper-sdk` (require `gmaps_scraper` → `GmapsScraper`).
- .NET package/namespace is `GmapsLeadFinder.GoogleMapsScraper`; CLI tool package is `GmapsLeadFinder.GoogleMapsScraper.Cli`.
- Go import alias is typically `gmaps`.
- PHP namespace is `GmapsLeadFinder\GoogleMapsScraper\`.
- Python / npm / Rust / PHP / Ruby / .NET CLIs may share the name `gmaps-scraper`; prefer `npx` / `python -m` / `cargo run --bin gmaps-scraper` / `go run ./cli` / `vendor/bin/gmaps-scraper` / `bundle exec gmaps-scraper` / `dotnet tool`.

Before any release: bump versions in `typescript/package.json`, `python/pyproject.toml`, `rust/Cargo.toml`, `ruby/lib/gmaps_scraper/version.rb`, `dotnet` csproj files, PHP `composer.json`, and Go/Rust/PHP/Ruby/.NET user-agent strings; update [`CHANGELOG.md`](../CHANGELOG.md); commit and push.

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
git tag go/v0.1.2
git push origin go/v0.1.2
GOPROXY=https://proxy.golang.org go list -m github.com/google-maps-lead-scraper/google-maps-scraper/go@v0.1.2
```

Library install (inside a module):

```bash
go mod init example.com/app
go get github.com/google-maps-lead-scraper/google-maps-scraper/go@v0.1.2
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

## Packagist (`gmapsleadfinder/google-maps-scraper`)

PHP is **not** in this monorepo (Packagist needs `composer.json` at the repo root). Source:

https://github.com/google-maps-lead-scraper/google-maps-scraper-php

No git subtree — edit and release only in that repo.

```bash
cd /path/to/google-maps-scraper-php
# bump version in composer.json + User-Agent if needed
git add -A && git commit -m "Release v0.1.2"
git tag v0.1.2
git push origin main --tags
```

First-time Packagist:

1. Log in at https://packagist.org
2. Submit → `https://github.com/google-maps-lead-scraper/google-maps-scraper-php`
3. Enable GitHub Service Hook / Packagist GitHub App for auto-updates

Users:

```bash
composer require gmapsleadfinder/google-maps-scraper:^0.1.2
```

## RubyGems (`google-maps-scraper-sdk`)

1. Register at https://rubygems.org and enable MFA.
2. Create an API key (Settings → API keys) with push permission.
3. Build and push from `ruby/`:

```bash
cd ruby
gem build google-maps-scraper-sdk.gemspec
gem push google-maps-scraper-sdk-0.1.2.gem
```

Users:

```bash
gem install google-maps-scraper-sdk
# or: bundle add google-maps-scraper-sdk
require "gmaps_scraper" # GmapsScraper::Client
```

## NuGet (`GmapsLeadFinder.GoogleMapsScraper`)

1. Create an account at https://www.nuget.org and an API key (API Keys → Create).
2. Scope the key to **Push** for new packages / your packages.
3. Pack and push from `dotnet/`:

```bash
cd dotnet
dotnet pack src/GmapsLeadFinder.GoogleMapsScraper -c Release
dotnet pack src/GmapsLeadFinder.GoogleMapsScraper.Cli -c Release
dotnet nuget push src/GmapsLeadFinder.GoogleMapsScraper/bin/Release/*.nupkg \
  --api-key $NUGET_API_KEY --source https://api.nuget.org/v3/index.json
dotnet nuget push src/GmapsLeadFinder.GoogleMapsScraper.Cli/bin/Release/*.nupkg \
  --api-key $NUGET_API_KEY --source https://api.nuget.org/v3/index.json
```

Users:

```bash
dotnet add package GmapsLeadFinder.GoogleMapsScraper
dotnet tool install -g GmapsLeadFinder.GoogleMapsScraper.Cli
```

## After publish

- Confirm:
  - https://www.npmjs.com/package/@gmapsleadfinder/google-maps-scraper
  - https://pypi.org/project/google-maps-scraper-sdk/
  - https://pkg.go.dev/github.com/google-maps-lead-scraper/google-maps-scraper/go
  - https://crates.io/crates/google-maps-scraper-sdk
  - https://packagist.org/packages/gmapsleadfinder/google-maps-scraper
  - https://rubygems.org/gems/google-maps-scraper-sdk
  - https://www.nuget.org/packages/GmapsLeadFinder.GoogleMapsScraper

No GitHub Actions publish workflow is configured yet; releases are manual.
