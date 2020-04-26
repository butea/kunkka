package registry

type Endpoint struct {
	Name string            `json:"name"`
	Meta map[string]string `json:"meta"`
}

type Node struct {
	Id   string            `json:"id"`
	Addr string            `json:"addr"`
	Ver  string            `json:"ver"`
	Meta map[string]string `json:"meta"`
}

type Service struct {
	Name      string               `json:"name"`
	Endpoints map[string]*Endpoint `json:"endpoints"`
	Nodes     []*Node              `json:"nodes"`
	Meta      map[string]string    `json:"meta"`
}
