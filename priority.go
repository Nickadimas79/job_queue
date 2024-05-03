package main

type Priority int

// Variables set in rank of importance with #1 taking highest priority
const (
	TIME_CRITICAL = iota + 1
	NOT_TIME_CRITICAL
)

func StringToPrio(str string) int {
	switch {
	case str == "TIME_CRITICAL":
		return TIME_CRITICAL
	case str == "NOT_TIME_CRITICAL":
		return NOT_TIME_CRITICAL
	}

	return 0
}

func (p Priority) String() string {
	return [...]string{"TIME_CRITICAL", "NOT_TIME_CRITICAL"}[p-1]
}

func (p Priority) EnumIndex() int {
	return int(p)
}
