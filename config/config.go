// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
  Period time.Duration `config:"period"`
  QueueType string `config:"type"`
  Connection  struct {
    Mysql struct {
      Username string `config:"username"`
      Password string `config:"password"`
      Host string `config:"host"`
      Database string `config:"database"`
    } `config:"mysql"`
  } `config:"connection"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
}
