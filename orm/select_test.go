package orm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



// SELECT * FROM xxx WHERE id = ?;
func Test_Selector_Build(t *testing.T) {

	testCases := []struct {
		name      string
		q         QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			// 调用 FROM
			name: "from",
			q: NewSelector[TestModel]().From("test_model_tab"),
			wantQuery: &Query{
				SQL: "SELECT * FROM `test_model_tab`;",
			},
			wantErr: nil,
		},
		{
			// 调用 FROM，但是传入空字符串
			name: "from_but_empty",
			q: NewSelector[TestModel]().From(""),
			wantQuery: &Query{
				SQL: "SELECT * FROM `test_model`;",
			},
			wantErr: nil,
		},
		{
			// 调用 FROM，但是传入空字符串
			name: "where",
			q: NewSelector[TestModel]().From("").Where(Predicate{}),
			wantQuery: &Query{
				SQL: "SELECT * FROM `test_model` WHERE id ;",
			},
			wantErr: nil,
		},
	}
                                                
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query, err := tc.q.Build()
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantQuery, query)
		})
	}
}


type TestModel struct {
	Field1 string
	Field2 int64
}