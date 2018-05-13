package queues

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/common"
  "fmt"
  "strings"
  "time"
	"github.com/resumecompanion/jobqueuebeat/config"
)

type DelayedJob struct {
  Cfg *config.Config
  DbConnection *sql.DB

}


func (dj *DelayedJob) Connect() {
  connString := fmt.Sprintf("%s:%s@%s/%s", dj.Cfg.Connection.Mysql.Username, dj.Cfg.Connection.Mysql.Password, dj.Cfg.Connection.Mysql.Host, dj.Cfg.Connection.Mysql.Database )
  dj.DbConnection, _ = sql.Open("mysql", connString)

  err := dj.DbConnection.Ping()


  if err != nil {
    logp.Warn("could not connect to DB")
    return 
  }
}

func (dj DelayedJob) CollectMetrics() common.MapStr {
  var failedQuery strings.Builder
  failedQuery.WriteString("select count(id) from delayed_jobs where run_at <= '")
  failedQuery.WriteString(time.Now().Format("2006-01-02 15:04:05"))
  failedQuery.WriteString("' AND locked_at IS NULL AND attempts = 0")
  fmt.Println(failedQuery.String())
  return common.MapStr {
    "running_jobs": dj.MetricForQuery("select count(id) from delayed_jobs where locked_at IS NOT NULL AND failed_at IS NULL"),
    "failed_jobs": dj.MetricForQuery("select count(id) from delayed_jobs where attempts > 0 AND failed_at IS NULL AND locked_at IS NULL"),
    "pending_jobs": dj.MetricForQuery(failedQuery.String()),
  }
}

func (dj DelayedJob) MetricForQuery(query string) int {
  rows, rowErr := dj.DbConnection.Query(query)

  if rowErr != nil {
    fmt.Println("Row Error")
    return 0
  }

  defer rows.Close()

  var count int
  for rows.Next() {
    rows.Scan(&count)
  }

  return count
}
