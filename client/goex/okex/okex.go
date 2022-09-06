package okex

const Okex = "Okex"

type FUTURES struct {
}

type Instruments struct {
	Arg struct {
		Channel  string `json:"channel,omitempty"`
		InstType string `json:"instType,omitempty"`
	} `json:"arg,omitempty"`
	Data []FUTURES `json:"data,omitempty"`
}
