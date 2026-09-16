/**
 * Create a job, poll manually, export CSV. Requires GMF_API_KEY.
 */
import { writeFile } from "node:fs/promises";
import { Client, TERMINAL_STATUSES } from "../src/index.js";
import type { PlaceRow } from "../src/index.js";

const KEYWORD = "dentists in Austin TX";
const OUT = "leads.csv";

function toCsv(rows: PlaceRow[]): string {
  const fieldnames: string[] = [];
  const seen = new Set<string>();
  for (const row of rows) {
    for (const key of Object.keys(row)) {
      if (!seen.has(key)) {
        seen.add(key);
        fieldnames.push(key);
      }
    }
  }
  const escape = (v: string) =>
    /[",\n\r]/.test(v) ? `"${v.replace(/"/g, '""')}"` : v;
  const lines = [fieldnames.join(",")];
  for (const row of rows) {
    lines.push(fieldnames.map((k) => escape(String(row[k] ?? ""))).join(","));
  }
  return lines.join("\n") + "\n";
}

async function main() {
  const client = new Client();
  const created = await client.createJob(KEYWORD);
  const jobId = created.jobId;
  console.log(`jobId=${jobId}`);

  for (;;) {
    const job = await client.getJob(jobId);
    const status = (job.status ?? "").toLowerCase();
    console.log(`status=${status} rowCount=${job.rowCount}`);
    if (TERMINAL_STATUSES.has(status)) {
      if (status === "failed") {
        throw new Error(job.error || "job failed");
      }
      break;
    }
    await new Promise((r) => setTimeout(r, 2000));
  }

  const rows: PlaceRow[] = [];
  let cursor: string | null = "0";
  while (cursor !== null) {
    const page = await client.getResults(jobId, { limit: 100, cursor });
    if (page.rows) rows.push(...page.rows);
    cursor =
      page.nextCursor === undefined || page.nextCursor === null
        ? null
        : String(page.nextCursor);
  }

  await writeFile(OUT, toCsv(rows), "utf8");
  console.log(`Wrote ${rows.length} rows to ${OUT}`);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
