using GmapsLeadFinder.GoogleMapsScraper;

using var client = new Client();
var me = await client.MeAsync();
Console.WriteLine($"plan={me.GetProperty("plan")} creditsRemaining={me.GetProperty("creditsRemaining")}");

var rows = await client.ScrapeAsync("dentists in Austin TX");
Console.WriteLine($"scraped {rows.Count} places");
if (rows.Count > 0)
{
    var r = rows[0];
    Console.WriteLine($"{r.GetValueOrDefault("Name")} {r.GetValueOrDefault("Phone")} {r.GetValueOrDefault("Website")} {r.GetValueOrDefault("Emails")}");
}
