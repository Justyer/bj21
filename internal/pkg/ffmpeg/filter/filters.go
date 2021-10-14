package filter

import "strings"

type Filters []*Filter

func (filters Filters) SingleParams() (params []string) {
	for _, filter := range filters {
		params = append(params, filter.String())
	}
	return
}

func (filters Filters) Params() (params []string) {
	params = append(params, "-filter_complex", filters.String())
	return
}

func (filters Filters) String() string {
	return strings.Join(filters.SingleParams(), ";")
}
