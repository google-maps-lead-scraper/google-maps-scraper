/**
 * Minimal scrape example. Requires GMF_API_KEY.
 * Run after build: npx tsx examples/quickstart.ts
 * or: node --import tsx examples/quickstart.ts
 */
import { Client } from "../src/index.js";

const KEYWORD = "dentists in Austin TX";

async function main() {
  const client = new Client();
  const me = await client.me();
  console.log(`plan=${me.plan} creditsRemaining=${me.creditsRemaining}`);

  const rows = await client.scrape(KEYWORD);
  console.log(`scraped ${rows.length} places for ${JSON.stringify(KEYWORD)}`);
  if (rows[0]) {
    console.log(rows[0].Name, rows[0].Phone, rows[0].Website, rows[0].Emails);
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
