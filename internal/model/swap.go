package model

type Pair struct {
	Token                  Token   `yaml:"token"`
	Strategy               string  `yaml:"strategy"`
	ExpectedProfitForward  float64 `yaml:"expected_profit_forward"`
	ExpectedProfitBackward float64 `yaml:"expected_profit_backward"`
	ExpectedProfileStep    float64 `yaml:"expected_profit_step"`
}

type Token struct {
	Input  string `yaml:"input"`
	Output string `yaml:"output"`
}

type Swap struct {
	ReturnAmount     uint64
	SpreadAmount     uint64
	CommissionAmount uint64
	Changed          float64
}
