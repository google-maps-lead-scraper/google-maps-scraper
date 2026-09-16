namespace GmapsLeadFinder.GoogleMapsScraper;

public sealed class ClientOptions
{
    public string? ApiKey { get; init; }
    public string? BaseUrl { get; init; }
    public TimeSpan Timeout { get; init; } = TimeSpan.FromSeconds(60);
}

public sealed class ScrapeOptions
{
    public int PollIntervalMs { get; init; } = Client.DefaultPollIntervalMs;
    public int TimeoutMs { get; init; } = Client.DefaultTimeoutMs;
    public int ResultLimit { get; init; } = Client.DefaultResultLimit;
}

public sealed class GetResultsOptions
{
    public int Limit { get; init; } = Client.DefaultResultLimit;
    public string? Cursor { get; init; }
}
