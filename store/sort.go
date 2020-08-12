package store

type SortType int8

const (
	SortTypeAsc SortType = iota + 1
	SortTypeDesc
)

func SortTypeString(typ SortType) string {
	switch typ {
	case SortTypeAsc:
		return " ASC"
	case SortTypeDesc:
		return " DESC"
	}
	return ""
}
