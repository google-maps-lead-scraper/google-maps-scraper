/**
 * Run multiple keywords sequentially (one running job per user).
 * Requires GMF_API_KEY.
 */
import { Client, JobConflictError } from "../src/index.js";

const KEYWORDS = [
  "dentists in Austin TX",
  "coffee shops in Austin TX",
];

async function main() {
  const client = new Client();
  for (const keyword of KEYWORDS) {
    console.log(`=== scraping ${JSON.stringify(keyword)} ===`);
    try {
      const rows = await client.scrape(keyword);
      console.log(`got ${rows.length} rows`);
    } catch (err) {
      if (err instanceof JobConflictError) {
        console.error("Another job is running (HTTP 409). Wait and retry.");
      }
      throw err;
    }
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
