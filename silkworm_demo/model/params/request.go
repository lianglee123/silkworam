package params

import (
	"strings"
	"fmt"
	"encoding/json"
	"strconv"
)

type Uint64StringSlice []uint64

func (slice Uint64StringSlice) MarshalJSON() ([]byte, error) {
	values := make([]string, len(slice))
	for i, value := range []uint64(slice) {
		values[i] = fmt.Sprintf(`"%v"`, value)
	}
	return []byte(fmt.Sprintf("[%v]", strings.Join(values, ","))), nil
}

func (slice *Uint64StringSlice) UnmarshalJSON(b []byte) error {
	// Try array of strings first.
	var values []string
	err := json.Unmarshal(b, &values)
	if err != nil {
		// Fall back to array of integers:
		var values []uint64
		if err := json.Unmarshal(b, &values); err != nil {
			return err
		}
		*slice = values
		return nil
	}
	*slice = make([]uint64, len(values))
	for i, value := range values {
		value, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		(*slice)[i] = value
	}
	return nil
}

type ListRequest struct {
	Limit           uint64            `form:"limit"`
	Offset          uint64            `form:"offset"`
	SearchFields    []string
	SearchFieldsStr string            `form:"search_fields"`
	SearchValue     string            `form:"search_value"`
	PartnerId       uint64
	ScopeId         uint64
	OrderField      string            `form:"order_field"`
	OrderType       string            `form:"order_type"`
	Id              uint64            `form:"id" json:"id,string"`
	Ids             Uint64StringSlice `form:"ids"`
	Exclude         Exclude
	NoTotal         bool
}

func (r *ListRequest) CorrectSearchFields() {
	r.SearchFields = strings.Split(r.SearchFieldsStr, ",")
}

type Exclude struct {
	Id  uint64            `json:"id,string"`
	Ids Uint64StringSlice `json:"ids"`
}
