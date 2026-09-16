use google_maps_scraper_sdk::{Client, ClientOptions};

fn main() {
    let client = Client::new(ClientOptions::default()).expect("client");
    let me = client.me().expect("me");
    println!(
        "plan={} creditsRemaining={}",
        me.plan, me.credits_remaining
    );

    let rows = client
        .scrape("dentists in Austin TX", Default::default())
        .expect("scrape");
    println!("scraped {} places", rows.len());
    if let Some(row) = rows.first() {
        println!(
            "{:?} {:?} {:?} {:?}",
            row.get("Name"),
            row.get("Phone"),
            row.get("Website"),
            row.get("Emails")
        );
    }
}
