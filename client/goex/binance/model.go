package binance

type BinanceBody struct {
	Body interface{}
}

type AggTrade struct {
	E    string `json:"e,omitempty"`
	BigE uint64 `json:"E,omitempty"`
	S    string `json:"s,omitempty"`
	A    uint64 `json:"a,omitempty"`
	P    string `json:"p,omitempty"`
	Q    string `json:"q,omitempty"`
	F    uint64 `json:"f,omitempty"`
	L    uint64 `json:"l,omitempty"`
	BigT uint64 `json:"T,omitempty"`
	M    bool   `json:"m,omitempty"`
	BigM bool   `json:"M,omitempty"`
}

type Trade struct {
	E    string `json:"e,omitempty"`
	BigE uint64 `json:"E,omitempty"`
	S    string `json:"s,omitempty"`
	A    uint64 `json:"a,omitempty"`
	P    string `json:"p,omitempty"`
	Q    string `json:"q,omitempty"`
	F    uint64 `json:"f,omitempty"`
	L    uint64 `json:"l,omitempty"`
	BigT uint64 `json:"T,omitempty"`
	M    bool   `json:"m,omitempty"`
	BigM bool   `json:"M,omitempty"`
}

type Kline struct {
	E    string `json:"e,omitempty"`
	BigE uint64 `json:"E,omitempty"`
	S    string `json:"s,omitempty"`
	K    struct {
		T    uint64 `json:"t,omitempty"`
		BigT uint64 `json:"T,omitempty"`
		S    string `json:"s,omitempty"`
		I    string `json:"i,omitempty"`
		F    uint64 `json:"f,omitempty"`
		// BigL uint64 `json:"L,omitempty"` // 这根K线期间末一笔成交ID
		O string `json:"o,omitempty"` // 这根K线期间第一笔成交价
		C string `json:"c,omitempty"` // 这根K线期间末一笔成交价
		H string `json:"h,omitempty"` // 这根K线期间最高成交价
		// L    string `json:"l,omitempty"` // 这根K线期间最低成交价
		V    string `json:"v,omitempty"` // 这根K线期间成交量
		N    uint64 `json:"n,omitempty"` // 这根K线期间成交笔数
		X    bool   `json:"x,omitempty"` // 这根K线是否完结(是否已经开始下一根K线)
		Q    string `json:"q,omitempty"` // 这根K线期间成交额
		BigV string `json:"V,omitempty"` // 主动买入的成交量
		BigQ string `json:"Q,omitempty"` // 主动买入的成交额
		BigB string `json:"B,omitempty"` // 忽略此参数
	} `json:"k,omitempty"`
}

type MiniTicker struct {
}

type Ticker struct {
	E    string `json:"e,omitempty"`
	BigE uint64 `json:"E,omitempty"`
	S    string `json:"s,omitempty"`
	P    string `json:"p,omitempty"`
	BigP string `json:"P,omitempty"`
	W    string `json:"w,omitempty"` // 平均价格
	X    string `json:"x,omitempty"` // 整整24小时之前，向前数的最后一次成交价格
	C    string `json:"c,omitempty"` // 最新成交价格
	BigQ string `json:"Q,omitempty"` // 最新成交交易的成交量
	B    string `json:"b,omitempty"` // 目前最高买单价
	BigB string `json:"B,omitempty"` // 目前最高买单价的挂单量
	A    string `json:"a,omitempty"` // 目前最低卖单价
	BigA string `json:"A,omitempty"` // 目前最低卖单价的挂单量
	O    string `json:"o,omitempty"` // 整整24小时前，向后数的第一次成交价格
	H    string `json:"h,omitempty"` // 24小时内最高成交价
	L    string `json:"l,omitempty"` // 24小时内最低成交价
	V    string `json:"v,omitempty"` // 24小时内成交量
	Q    string `json:"Q,omitempty"` // 24小时内成交额
	BigO uint64 `json:"O,omitempty"` // 统计开始时间
	BigC uint64 `json:"C,omitempty"` // 统计结束时间
	BigF uint64 `json:"F,omitempty"` // 24小时内第一笔成交交易ID
	BigL uint64 `json:"L,omitempty"` // 24小时内最后一笔成交交易ID
	N    uint64 `json:"n,omitempty"` // 24小时内成交数
}

type BookTicker struct {
	U    uint64 `json:"u,omitempty"` // order book updateId
	S    string `json:"s,omitempty"` // 交易对
	B    string `json:"b,omitempty"` // 买单最优挂单价格
	BigB string `json:"B,omitempty"` // 买单最优挂单数量
	A    string `json:"a,omitempty"` // 卖单最优挂单价格
	BigA string `json:"A,omitempty"` // 卖单最优挂单数量
}

type DepthUpdate struct {
}
