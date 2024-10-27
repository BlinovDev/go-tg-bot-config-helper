package tgconfighelper

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// define struct to store config
type Config struct {
	Bot struct {
		// For any telegram bot
		Token string `yaml:"token"`
		// For GetUpdates approach
		Timeout int `yaml:"timeout"`
		// For Webhook approach
		WebHookURL string `yaml:"webhook_url"`
		CertFile   string `yaml:"cert_file"`
		KeyFile    string `yaml:"key_file"`
		// Environment(can be set here to prevent its passing through ENV var)
		Environment string `yaml:"environment"`
		// Some reserved configs for cron setup
		FirstCron  string `yaml:"first_cron_time_pattern"`
		SecondCron string `yaml:"second_cron_time_pattern"`
		// Path to your storage file if you don't use DB
		FilePath string `yaml:"file_path"`
	} `yaml:"bot"`
}

// Reads config.yml in project root and returns config if file was parsed to Config successfully
func LoadConfig() (*Config, error) {
	var config Config

	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
