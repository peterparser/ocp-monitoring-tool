package util

type Query struct {
	PrettyName string `yaml:"prettyName,omitempty"`
	Expression string `yaml:"expression"`
	Start      string `yaml:"start,omitempty"`
	End        string `yaml:"end,omitempty"`
	Result     string
}

type Configuration struct {
	OcpOauthUrl        string  `yaml:"ocpOauthUrl"`
	PrometheusEndpoint string  `yaml:"prometheusEndpoint"`
	Username           string  `yaml:"username"`
	Password           string  `yaml:"password"`
	Queries            []Query `yaml:"queries"`
}

type QueryResult struct {
	Result     string
	PrettyName string
}

type QueryPerformer interface {
	PerformQuery(Url string, token string, queries []Query, plotterChan chan<- QueryResult)
}
