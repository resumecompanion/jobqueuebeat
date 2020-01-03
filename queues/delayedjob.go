package queues

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/go-sql-driver/mysql"
	"github.com/resumecompanion/jobqueuebeat/config"
)

// DelayedJob for connect redis
type DelayedJob struct {
	Cfg          *config.Config
	DbConnection *sql.DB
}

// Connect is to connect mysql
func (dj *DelayedJob) Connect() {
	var connString string

	logp.Warn("setting SSL")

	tls := SetupTLSConfig(dj.Cfg.Connection.Mysql.SslCa, dj.Cfg.Connection.Mysql.SslCert, dj.Cfg.Connection.Mysql.SslKey)
	mysql.RegisterTLSConfig("custom", &tls)

	connString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?tls=custom", dj.Cfg.Connection.Mysql.Username, dj.Cfg.Connection.Mysql.Password, dj.Cfg.Connection.Mysql.Host, dj.Cfg.Connection.Mysql.Database)

	dj.DbConnection, _ = sql.Open("mysql", connString)

	err := dj.DbConnection.Ping()

	if err != nil {
		logp.Warn(err.Error())
		logp.Warn("could not connect to DB")
		return
	}
}

// CollectMetrics collect fail pending running jobs
func (dj DelayedJob) CollectMetrics() common.MapStr {
	var failedQuery strings.Builder
	failedQuery.WriteString("select count(id) from delayed_jobs where run_at <= '")
	failedQuery.WriteString(time.Now().Format("2006-01-02 15:04:05"))
	failedQuery.WriteString("' AND locked_at IS NULL AND attempts = 0")
	fmt.Println(failedQuery.String())
	return common.MapStr{
		"running_jobs": dj.MetricForQuery("select count(id) from delayed_jobs where locked_at IS NOT NULL AND failed_at IS NULL"),
		"failed_jobs":  dj.MetricForQuery("select count(id) from delayed_jobs where attempts > 0 AND failed_at IS NULL AND locked_at IS NULL"),
		"pending_jobs": dj.MetricForQuery(failedQuery.String()),
	}
}

// MetricForQuery do query to mysql
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
