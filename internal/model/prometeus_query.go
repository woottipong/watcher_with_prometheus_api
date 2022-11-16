package model

type Response struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	ResultType string   `json:"resultType"`
	Results    []Result `json:"result"`
}

type Result struct {
	Metric map[string]interface{} `json:"metric"`
	Value  [2]string              `json:"value"`
}
