package limiter

// func initRedisClient(ctx context.Context) (*redis.Client, error) {
// 	redisConn := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Username: "",
// 		Password: "",
// 		DB:       0,
// 	})
// 	if err := redisConn.Ping(ctx).Err(); err != nil {
// 		return nil, fmt.Errorf("failed to connect to redis: %w", err)
// 	}
// 	return redisConn, nil
// }

// func TestLimitter(t *testing.T) {
// 	logger := logger.New()
// 	ctx := context.Background()

// 	redisClient, err := initRedisClient(ctx)
// 	if err != nil {
// 		t.Fatal("init redis clinet error", err)
// 	}

// 	limiterCfg := Config{
// 		CachePrefix: "test-microservice-name",
// 	}

// 	l, err := New(logger, limiterCfg, redisClient)
// 	if err != nil {
// 		t.Fatal("init limiter error", err)
// 	}

// 	testServiceName := "test-service"

// 	//svc := NewService(testServiceName, WithFormattedLimit("2-S"))
// 	svc := NewService(testServiceName, WithPeriodLimit(time.Second, 2))

// 	if err := l.RegisterServices(svc); err != nil {
// 		t.Fatal("failed to register services", err)
// 	}

// 	svcLimiter, err := l.GetService(testServiceName)
// 	if err != nil {
// 		t.Fatalf("failed to get service with name %s: %v", testServiceName, err)
// 	}

// 	stats, err := svcLimiter.Peek(ctx)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	fmt.Printf("init %+v\n", stats)

// 	for i := 0; i < 12; i++ {
// 		stats, err := svcLimiter.Get(ctx)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		fmt.Printf("%+v\n", stats)

// 		isReached, err := svcLimiter.IsReached(ctx)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		assert.Equal(t, isReached, stats.Reached)

// 		time.Sleep(time.Millisecond * 100)
// 	}
// }
