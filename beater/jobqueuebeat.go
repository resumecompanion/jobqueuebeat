package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/resumecompanion/jobqueuebeat/config"
	"github.com/resumecompanion/jobqueuebeat/queues"
)

type Jobqueuebeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Jobqueuebeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Jobqueuebeat) Run(b *beat.Beat) error {
	logp.Info("jobqueuebeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	counter := 1

	var t interface{}

	if bt.config.Connection.Mysql.Username != "" {
		t = queues.DelayedJob{
			Cfg: &bt.config,
		}
	} else if bt.config.Connection.Sidekiq.Host != "" {
		t = queues.Sidekiq{
			Cfg: &bt.config,
		}
	} else {
		t = queues.Resque{
			Cfg: &bt.config,
		}
	}
	djb, dok := t.(queues.DelayedJob)
	skb, sok := t.(queues.Sidekiq)
	rsb, rok := t.(queues.Resque)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		var fields common.MapStr
		if dok {
			djb.Connect()
			fields = djb.CollectMetrics()
			fields["background_runner"] = djb.Cfg.Connection.Mysql.Type
			djb.DbConnection.Close()
		} else if sok {
			skb.Connect()
			fields = skb.CollectMetrics()
			fields["background_runner"] = skb.Cfg.Connection.Sidekiq.Type
		} else if rok {
			rsb.Connect()
			fields = rsb.CollectMetrics()
			fields["background_runner"] = rsb.Cfg.Connection.Resque.Type
		}
		fields["type"] = b.Info.Name
		fields["counter"] = counter

		event := beat.Event{
			Timestamp: time.Now(),
			Fields:    fields,
		}

		bt.client.Publish(event)
		logp.Info("Event sent")
		counter++
	}
}

func (bt *Jobqueuebeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
