namespace GmapsLeadFinder.GoogleMapsScraper.Exceptions;

public class ApiException : Exception
{
    public int? StatusCode { get; }
    public object? Body { get; }

    public ApiException(string message, int? statusCode = null, object? body = null, Exception? inner = null)
        : base(message, inner)
    {
        StatusCode = statusCode;
        Body = body;
    }
}

public class AuthenticationException : ApiException
{
    public AuthenticationException(string message, int? statusCode = null, object? body = null)
        : base(message, statusCode, body) { }
}

public class PlanNotAllowedException : ApiException
{
    public PlanNotAllowedException(string message, int? statusCode = null, object? body = null)
        : base(message, statusCode, body) { }
}

public class InsufficientCreditsException : ApiException
{
    public InsufficientCreditsException(string message, int? statusCode = null, object? body = null)
        : base(message, statusCode, body) { }
}

public class JobConflictException : ApiException
{
    public JobConflictException(string message, int? statusCode = null, object? body = null)
        : base(message, statusCode, body) { }
}

public class NotFoundException : ApiException
{
    public NotFoundException(string message, int? statusCode = null, object? body = null)
        : base(message, statusCode, body) { }
}

public class BadRequestException : ApiException
{
    public BadRequestException(string message, int? statusCode = null, object? body = null)
        : base(message, statusCode, body) { }
}

public class TimeoutException : ApiException
{
    public TimeoutException(string message, int? statusCode = null, object? body = null)
        : base(message, statusCode, body) { }
}
