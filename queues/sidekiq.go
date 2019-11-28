package queues

import (
	"fmt"
	"strconv"
	"time"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/go-redis/redis"
	"github.com/resumecompanion/jobqueuebeat/config"
)

// Sidekiq for connect redis
type Sidekiq struct {
	Cfg          *config.Config
	DbConnection *redis.Client
}

// Connect is to connect redis
func (sk *Sidekiq) Connect() {
	connAddr := fmt.Sprintf("%s:%s", sk.Cfg.Connection.Sidekiq.Host, sk.Cfg.Connection.Sidekiq.Port)
	sk.DbConnection = redis.NewClient(&redis.Options{
		Addr:     connAddr,
		Password: sk.Cfg.Connection.Sidekiq.Password,
		DB:       0, // use default DB
	})

	_, err := sk.DbConnection.Ping().Result()
	if err != nil {
		logp.Warn("could not connect to redis")
		return
	}
}

// CollectMetrics is to collecting all required output
func (sk Sidekiq) CollectMetrics() common.MapStr {
	r := common.MapStr{
		"schedule_jobs": sk.scheduleJobs(),
		"failed_jobs":   sk.failedJobs(),
	}

	queues := sk.queuesList()
	for _, queue := range queues {
		k := fmt.Sprintf("%s_jobs", queue)
		r[k] = sk.queueData(queue)
	}

	return r
}

func (sk Sidekiq) failedJobs() int {
	tProcess := fmt.Sprintf("stat:failed:%s", time.Now().Format("2006-01-02"))
	fJ, _ := sk.DbConnection.Get(tProcess).Result()
	r, _ := strconv.Atoi(fJ)
	return r
}

func (sk Sidekiq) scheduleJobs() int {
	result, _ := sk.DbConnection.ZRange("schedule", 0, -1).Result()
	return len(result)
}

func (sk Sidekiq) queueData(q string) int {
	queueName := fmt.Sprintf("queue:%s", q)
	queueCount, _ := sk.DbConnection.LRange(queueName, 0, -1).Result()

	return len(queueCount)
}

func (sk Sidekiq) queuesList() []string {
	queueList, _ := sk.DbConnection.SMembers("queues").Result()
	return queueList
}
