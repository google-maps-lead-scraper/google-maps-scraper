# TypeScript client for GMaps Lead Finder

```bash
npm install @gmapsleadfinder/google-maps-scraper
export GMF_API_KEY=gmf_your_key_here
npx gmaps-scraper me
```

```ts
import { Client } from "@gmapsleadfinder/google-maps-scraper";

const rows = await new Client().scrape("dentists in Austin TX");
```

npm: https://www.npmjs.com/package/@gmapsleadfinder/google-maps-scraper  
Docs: [../docs/typescript.md](../docs/typescript.md) · API: https://gmapsleadfinder.com/docs/api
