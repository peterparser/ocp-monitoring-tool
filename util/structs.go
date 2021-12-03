package util

type PromQuery struct {
	Expression string `yaml:"expression"`
	Start      string `yaml:"start,omitempty"`
	End        string `yaml:"end,omitempty"`
}

type Configuration struct {
	OcpOauthUrl        string      `yaml:"ocpOauthUrl"`
	PrometheusEndpoint string      `yaml:"prometheusEndpoint"`
	Username           string      `yaml:"username"`
	Password           string      `yaml:"password"`
	Queries            []PromQuery `yaml:"queries"`
}
