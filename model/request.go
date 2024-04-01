package model

type GetSrvTimeReq struct {
}

type PingSrvReq struct {
}

type KLinesReq struct {
	Symbol    string       `schema:"symbol"`
	Interval  IntervalType `schema:"interval"`
	StartTime int64        `schema:"startTime,omitempty"`
	EndTime   int64        `schema:"endTime,omitempty"`
	Limit     int          `schema:"limit,omitempty"`
}
