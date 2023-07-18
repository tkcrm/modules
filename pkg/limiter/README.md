# Limiter

Thit package contains simple limiter

## Features

- Sliding window limiter for requests per period
- Register multiple services in one limiter
- It can be used in multiple instances of microservices
- It can be used in multiple stages with one redis
- It can be used with HTTP middleware and custom services logic
- Trade safe instance
- Formated config for enviroment variables: 10-S / 50-M / 100-H / 1000-D

## How to use in custom microservice

```go
// define config
limiterCfg := Config{
    CachePrefix: "dev-service-name",
}

// init limiter instance
l, err := limiter.New(logger, limiterCfg, redisClient)
if err != nil {
    log.Fatal("init limiter error", err)
}

// register new service
if err := l.RegisterServices(
    // 50 requests per second
    limiter.NewService("test-service", limiter.WithFormattedLimit("50-S")),
    // 1000 requests per 10 minutes
    limiter.NewService("test-service2", limiter.WithPeriodLimit(time.Minute * 10, 1000)),
); err != nil {
    log.Fatal("failed to register services", err)
}

// get registered service
svcLimiter, err := l.GetService(testServiceName)
if err != nil {
    log.Fatalf("failed to get service with name %s: %v", testServiceName, err)
}

// get current limit stats
stats, err := svcLimiter.Peek(ctx)
if err != nil {
    log.Fatal(err)
}

// return is limit reached
isReached, err := svcLimiter.IsReached(ctx)
if err != nil {
    log.Fatal(err)
}

// increment one request and get new stats
stats, err := svcLimiter.Get(ctx)
if err != nil {
    log.Fatal(err)
}

// increment custom integer value
stats, err := svcLimiter.Increment(ctx, 10)
if err != nil {
    log.Fatal(err)
}

// reset data
stats, err := svcLimiter.Reset(ctx)
if err != nil {
    log.Fatal(err)
}

// increment one request and get new stats for user ip
userIP := "192.143.203.16"
stats, err := svcLimiter.Get(ctx, limiter.WithCacheKey(userIP))
if err != nil {
    log.Fatal(err)
}
```

## How to use with fiber

```go
func IPRateLimit() fiber.Handler {
    // 1. Configure
    // define config
    limiterCfg := Config{
        CachePrefix: "dev-s-apiservice",
    }

    // init limiter instance
    l, err := New(logger, limiterCfg, redisClient)
    if err != nil {
        log.Fatal("init limiter error", err)
    }

    // register new service
    if err := l.RegisterServices(
        limiter.NewService("http-ip-middleware", WithFormattedLimit("50-S")),
    ); err != nil {
        log.Fatal("failed to register services", err)
    }

    // get registered service
    svcLimiter, err := l.GetService("http-ip-middleware")
    if err != nil {
        log.Fatalf("failed to get service with name %s: %v", "http-ip-middleware", err)
    }

    // 2. Return middleware handler
    return func(c *fiber.Ctx) error {
        ctx := c.Context()
        limiterCtx, err := svcLimiter.Get(ctx, limiter.WithCacheKey(c.IP()))
        if err != nil {
            log.Printf("IPRateLimit - ipRateLimiter.Get - err: %v, %s on %s", err, c.IP(), c.Path())
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "success": false,
                "message": err,
            })
        }

        c.Set("X-RateLimit-Limit", strconv.FormatInt(limiterCtx.Limit, 10))
        c.Set("X-RateLimit-Remaining", strconv.FormatInt(limiterCtx.Remaining, 10))
        c.Set("X-RateLimit-Reset", strconv.FormatInt(limiterCtx.Reset, 10))

        if limiterCtx.Reached {
            log.Printf("Too Many Requests from %s on %s", c.IP(), c.Path())
            return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
                "success": false,
                "message": "Too Many Requests on " + c.Path(),
            })
        }
        return c.Next()
    }
}
```
