package main

import (
	"context"
	"log"
	"sync"
	"testdb/database"
	"testdb/utils"
)

func main() {
	err := utils.Init(".")
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

	//1000 goroutines to process messages
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				m, err := consumer.ReadMessage(context.Background())
				if err != nil {
					log.Fatal(err)
					break
				}
				utils.ProcessMessage(pgDb, redis, m.Value)
			}
		}()
	}

	wg.Wait()
}
