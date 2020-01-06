package config

import "time"

// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations
type Config struct {
	Period     time.Duration `config:"period"`
	Connection struct {
		Mysql struct {
			Username string `config:"username"`
			Password string `config:"password"`
			Host     string `config:"host"`
			Database string `config:"database"`
			Type     string `config:"type"`
			SslCa    string `config:"sslca"`
			SslKey   string `config:"sslkey"`
			SslCert  string `config:"sslcert"`
		} `config:"mysql"`
		Sidekiq struct {
			Password string `config:"password"`
			Host     string `config:"host"`
			Port     string `config:"port"`
			Type     string `config:"type"`
		} `config:"sidekiq"`
		Resque struct {
			Password string `config:"password"`
			Host     string `config:"host"`
			Port     string `config:"port"`
			Type     string `config:"type"`
		} `config:"resque"`
	} `config:"connection"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
}
