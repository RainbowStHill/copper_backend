package identity_server_test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"

	server "github.com/rainbowsthill/copper_backend/service/id/server"
)

func TestGenerate(t *testing.T) {
	startsAt, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2024-12-25 00:00:00 +0800 CST")
	opt := server.SnowflakeOpt{
		TimestampBits:    41,
		DataCenterIDBits: 5,
		InstanceIDBits:   5,
		IncrIDBits:       12,

		StartsAt:   startsAt.UnixMilli(),
		DataCenter: 12,
		Instance:   2,
	}

	generator := server.NewSnowflake(opt)

	const coroutineNum = 10
	const timesPerCoroutine = 1000

	results := []server.Unique{}
	var mtx sync.Mutex

	var wg sync.WaitGroup

	for range [coroutineNum]struct{}{} {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for range [timesPerCoroutine]struct{}{} {
				id := generator.Generate()
				if id == nil {
					continue
				}

				mtx.Lock()
				results = append(results, id)
				mtx.Unlock()
			}
		}()
	}

	wg.Wait()

	if len(results) < coroutineNum*timesPerCoroutine {
		t.Fatalf("Some test cases has been failed:\n\tPassed: %d\n\tTotal: %d", len(results), coroutineNum*timesPerCoroutine)
	}

	fmt.Printf("Result samples: ")
	for i := 0; i < len(results); i += timesPerCoroutine {
		res, err := strconv.Atoi(results[i].String())
		if err != nil {
			t.Fatalf("Invalid snowflake id generated: %s", results[i].String())
		}
		fmt.Printf("%x\t", res)
	}
	fmt.Printf("...")
}
