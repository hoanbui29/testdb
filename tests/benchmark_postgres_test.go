package tests

import (
	"context"
	"io"
	"log"
	"sync"
	"testdb/database"
	"testdb/utils"
	"testing"
)

func BenchmarkPostgres(b *testing.B) {
	err := utils.Init("../")
	if err != nil {
		log.Fatal(err)
	}
	pgDb := database.OpenPg()
	redis := database.OpenRedis()
	consumer := database.GetKafkaConsumer()
	defer consumer.Close()
	defer pgDb.Close()
	defer redis.Close()
	var wg sync.WaitGroup

	b.ResetTimer()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//1000 goroutines to process messages
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				m, err := consumer.ReadMessage(ctx)
				if err != nil {
					if err == io.EOF {
						break
					} else if err == context.DeadlineExceeded {
						log.Println("Deadline exceeded")
						break
					} else {
						log.Fatal(err)
						break
					}
				}
				utils.ProcessMessage(pgDb, redis, m.Value)
			}
		}()
	}
	wg.Wait()
}
