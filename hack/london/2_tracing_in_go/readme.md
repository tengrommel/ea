# tracing in Go

# Why should I use distributed tracing?

- Instrumentation doesn't show causality
    - or it's pretty hard(logs)
- A main feature is correction between events
- "I want to know why call X didn't work"
- "Why is Y so slow"
- "How do we optimize the performance"

# Tracing Basics
> A trace is a data/execution path through the system
    
    type struct Tracing {
        TraceId: string
        SpanId: int64
        ParentId: int64
    }
