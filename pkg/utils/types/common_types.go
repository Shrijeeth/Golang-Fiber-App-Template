package types

type ParserType uint64

const (
	QueryParserType ParserType = iota + 1
	BodyParserType
)
