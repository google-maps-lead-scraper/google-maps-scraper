using System.Text;
using System.Text.Json;
using GmapsLeadFinder.GoogleMapsScraper;
using GmapsLeadFinder.GoogleMapsScraper.Exceptions;

static void PrintHelp()
{
    Console.WriteLine("""
        Usage: gmaps-scraper <command> [options]

        Commands:
          me                          Show plan and credits
          scrape "<keyword>" [opts]   Scrape one keyword and print/save rows

        Options for scrape:
          --out PATH                  Output path (.json or .csv). Default: JSON to stdout
          --poll-interval-ms MS       Job poll interval (default: 2000)
          --timeout-ms MS             Max wait for job completion (default: 600000)

        Global:
          --version                   Print version
          -h, --help                  Show this help
        """);
}

static void WriteOutput(IReadOnlyList<Dictionary<string, string>> rows, string? outPath)
{
    if (outPath is null)
    {
        Console.WriteLine(JsonSerializer.Serialize(rows, new JsonSerializerOptions { WriteIndented = true }));
        return;
    }

    var path = Path.GetFullPath(outPath);
    var dir = Path.GetDirectoryName(path);
    if (!string.IsNullOrEmpty(dir))
    {
        Directory.CreateDirectory(dir);
    }

    if (Path.GetExtension(path).Equals(".csv", StringComparison.OrdinalIgnoreCase))
    {
        var fieldnames = new List<string>();
        var seen = new HashSet<string>(StringComparer.Ordinal);
        foreach (var row in rows)
        {
            foreach (var key in row.Keys)
            {
                if (seen.Add(key))
                {
                    fieldnames.Add(key);
                }
            }
        }

        if (fieldnames.Count == 0)
        {
            fieldnames.Add("Name");
        }

        var sb = new StringBuilder();
        sb.AppendLine(string.Join(",", fieldnames.Select(EscapeCsv)));
        foreach (var row in rows)
        {
            sb.AppendLine(string.Join(",", fieldnames.Select(k => EscapeCsv(row.GetValueOrDefault(k, "")))));
        }

        File.WriteAllText(path, sb.ToString(), Encoding.UTF8);
    }
    else
    {
        File.WriteAllText(
            path,
            JsonSerializer.Serialize(rows, new JsonSerializerOptions { WriteIndented = true }) + "\n",
            Encoding.UTF8);
    }

    Console.Error.WriteLine($"Wrote {rows.Count} rows to {path}");
}

static string EscapeCsv(string value)
{
    if (value.Contains('"') || value.Contains(',') || value.Contains('\n') || value.Contains('\r'))
    {
        return "\"" + value.Replace("\"", "\"\"") + "\"";
    }

    return value;
}

var argv = args.ToList();
if (argv.Count == 0 || argv.Contains("-h") || argv.Contains("--help"))
{
    PrintHelp();
    return argv.Count == 0 ? 1 : 0;
}

if (argv.Contains("--version"))
{
    Console.WriteLine(Client.Version);
    return 0;
}

var cmd = argv[0];
argv.RemoveAt(0);

try
{
    switch (cmd)
    {
        case "me":
        {
            using var client = new Client();
            var me = await client.MeAsync();
            Console.WriteLine(JsonSerializer.Serialize(me, new JsonSerializerOptions { WriteIndented = true }));
            return 0;
        }
        case "scrape":
        {
            string? outPath = null;
            var pollIntervalMs = Client.DefaultPollIntervalMs;
            var timeoutMs = Client.DefaultTimeoutMs;
            string? keyword = null;

            for (var i = 0; i < argv.Count; i++)
            {
                var a = argv[i];
                if (a == "--out" && i + 1 < argv.Count)
                {
                    outPath = argv[++i];
                }
                else if (a == "--poll-interval-ms" && i + 1 < argv.Count)
                {
                    pollIntervalMs = int.Parse(argv[++i]);
                }
                else if (a == "--timeout-ms" && i + 1 < argv.Count)
                {
                    timeoutMs = int.Parse(argv[++i]);
                }
                else if (!a.StartsWith('-') && keyword is null)
                {
                    keyword = a;
                }
                else
                {
                    throw new ArgumentException($"unknown argument {a}");
                }
            }

            if (string.IsNullOrWhiteSpace(keyword))
            {
                throw new ArgumentException("keyword is required");
            }

            using var client = new Client();
            var rows = await client.ScrapeAsync(keyword, new ScrapeOptions
            {
                PollIntervalMs = pollIntervalMs,
                TimeoutMs = timeoutMs,
            });
            WriteOutput(rows, outPath);
            return 0;
        }
        default:
            throw new ArgumentException($"unknown command {cmd}");
    }
}
catch (ApiException ex)
{
    Console.Error.WriteLine($"error: {ex.Message}");
    if (ex.StatusCode is not null)
    {
        Console.Error.WriteLine($"status: {ex.StatusCode}");
    }

    return 1;
}
catch (ArgumentException ex)
{
    Console.Error.WriteLine($"error: {ex.Message}");
    return 1;
}
