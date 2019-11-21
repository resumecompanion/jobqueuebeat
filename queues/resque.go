package queues

import (
	"fmt"
	"strings"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/go-redis/redis"
	"github.com/resumecompanion/jobqueuebeat/config"
)

// Resque for connect redis
type Resque struct {
	Cfg          *config.Config
	DbConnection *redis.Client
}

// Connect is to connect redis
func (rQ *Resque) Connect() {
	connAddr := fmt.Sprintf("%s:%s", rQ.Cfg.Connection.Resque.Host, rQ.Cfg.Connection.Resque.Port)
	rQ.DbConnection = redis.NewClient(&redis.Options{
		Addr:     connAddr,
		Password: rQ.Cfg.Connection.Resque.Password,
		DB:       0, // use default DB
	})

	_, err := rQ.DbConnection.Ping().Result()
	if err != nil {
		logp.Warn("could not connect to redis")
		return
	}
}

// CollectMetrics is to collecting all required output
func (rQ Resque) CollectMetrics() common.MapStr {
	r := common.MapStr{
		"schedule_jobs": rQ.scheduleJobs(),
		"failed_jobs":   rQ.failedJobs(),
	}

	queues := rQ.queuesList()
	for _, queue := range queues {
		// queues in resque looks like `environment_default` with suffix
		// so in staging queue name is resque:queue:staging_default
		queueName := strings.Split(queue, "_")[1]

		k := fmt.Sprintf("%s_jobs", queueName)
		r[k] = rQ.queueData(queue)
	}

	return r
}

func (rQ Resque) failedJobs() int64 {
	fJ, _ := rQ.DbConnection.LLen("resque:failed").Result()
	return fJ
}

// Resque will add key with prefix resque:delayed when schedule delayed jobs
func (rQ Resque) scheduleJobs() int {
	result, _ := rQ.DbConnection.Keys("resque:delayed:*").Result()

	rC := len(result)
	return rC
}

func (rQ Resque) queueData(q string) int64 {
	qN := fmt.Sprintf("resque:queue:%s", q)
	queueCount, _ := rQ.DbConnection.LLen(qN).Result()

	return queueCount
}

func (rQ Resque) queuesList() []string {
	queueList, _ := rQ.DbConnection.SMembers("resque:queues").Result()
	return queueList
}
