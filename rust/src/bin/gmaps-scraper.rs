use google_maps_scraper_sdk::{Client, ClientOptions, ScrapeOptions};
use std::env;
use std::fs;
use std::io::{self, Write};
use std::path::Path;
use std::process;

fn main() {
    let mut args: Vec<String> = env::args().skip(1).collect();
    if args.is_empty() || args.iter().any(|a| a == "-h" || a == "--help") {
        print_help();
        if args.is_empty() {
            process::exit(1);
        }
        return;
    }
    if args.iter().any(|a| a == "--version") {
        println!("0.1.1");
        return;
    }

    let cmd = args.remove(0);
    match cmd.as_str() {
        "me" => cmd_me(),
        "scrape" => {
            if args.is_empty() || args[0].starts_with('-') {
                eprintln!("error: scrape requires a keyword");
                process::exit(1);
            }
            let keyword = args.remove(0);
            let mut out: Option<String> = None;
            let mut opts = ScrapeOptions::default();
            let mut i = 0;
            while i < args.len() {
                match args[i].as_str() {
                    "--out" => {
                        i += 1;
                        out = Some(args.get(i).cloned().unwrap_or_default());
                    }
                    "--poll-interval-ms" => {
                        i += 1;
                        opts.poll_interval_ms = args.get(i).and_then(|s| s.parse().ok());
                    }
                    "--timeout-ms" => {
                        i += 1;
                        opts.timeout_ms = args.get(i).and_then(|s| s.parse().ok());
                    }
                    other => {
                        eprintln!("error: unknown argument {other}");
                        process::exit(1);
                    }
                }
                i += 1;
            }
            cmd_scrape(&keyword, out.as_deref(), opts);
        }
        other => {
            eprintln!("error: unknown command {other}");
            print_help();
            process::exit(1);
        }
    }
}

fn print_help() {
    println!(
        "gmaps-scraper (Rust) 0.1.1

Usage:
  gmaps-scraper me
  gmaps-scraper scrape \"<keyword>\" [--out leads.json|leads.csv]

Env:
  GMF_API_KEY     required
  GMF_BASE_URL    optional"
    );
}

fn client() -> Client {
    match Client::new(ClientOptions::default()) {
        Ok(c) => c,
        Err(e) => {
            eprintln!("error: {e}");
            process::exit(1);
        }
    }
}

fn cmd_me() {
    let c = client();
    match c.me() {
        Ok(me) => {
            println!("{}", serde_json::to_string_pretty(&me).unwrap());
        }
        Err(e) => {
            eprintln!("error: {e}");
            process::exit(1);
        }
    }
}

fn cmd_scrape(keyword: &str, out: Option<&str>, opts: ScrapeOptions) {
    let c = client();
    let rows = match c.scrape(keyword, opts) {
        Ok(r) => r,
        Err(e) => {
            eprintln!("error: {e}");
            process::exit(1);
        }
    };
    if let Err(e) = write_output(&rows, out) {
        eprintln!("error: {e}");
        process::exit(1);
    }
}

fn write_output(
    rows: &[google_maps_scraper_sdk::PlaceRow],
    out: Option<&str>,
) -> io::Result<()> {
    match out {
        None => {
            let s = serde_json::to_string_pretty(rows).unwrap();
            println!("{s}");
            Ok(())
        }
        Some(path) if path.to_ascii_lowercase().ends_with(".csv") => write_csv(path, rows),
        Some(path) => {
            if let Some(parent) = Path::new(path).parent() {
                if !parent.as_os_str().is_empty() {
                    fs::create_dir_all(parent)?;
                }
            }
            let s = serde_json::to_string_pretty(rows).unwrap();
            fs::write(path, format!("{s}\n"))?;
            eprintln!("Wrote {} rows to {path}", rows.len());
            Ok(())
        }
    }
}

fn write_csv(path: &str, rows: &[google_maps_scraper_sdk::PlaceRow]) -> io::Result<()> {
    let mut fields: Vec<String> = Vec::new();
    for row in rows {
        for k in row.keys() {
            if !fields.iter().any(|f| f == k) {
                fields.push(k.clone());
            }
        }
    }
    if fields.is_empty() {
        fields.push("Name".into());
    }
    if let Some(parent) = Path::new(path).parent() {
        if !parent.as_os_str().is_empty() {
            fs::create_dir_all(parent)?;
        }
    }
    let mut f = fs::File::create(path)?;
    writeln!(f, "{}", fields.join(","))?;
    for row in rows {
        let line: Vec<String> = fields
            .iter()
            .map(|k| csv_escape(row.get(k).map(|s| s.as_str()).unwrap_or("")))
            .collect();
        writeln!(f, "{}", line.join(","))?;
    }
    eprintln!("Wrote {} rows to {path}", rows.len());
    Ok(())
}

fn csv_escape(v: &str) -> String {
    if v.contains(',') || v.contains('"') || v.contains('\n') {
        format!("\"{}\"", v.replace('"', "\"\""))
    } else {
        v.to_string()
    }
}
