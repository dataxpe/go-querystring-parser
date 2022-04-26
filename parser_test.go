package querystring

import (
	"testing"

	"github.com/go-test/deep"
)

func TestParserA(t *testing.T) {
	cond, err := Parse("message: test\\ value AND datetime: [\"2020-01-01T00:00:00\" TO \"2020-12-31T00:00:00\"]")
	if err != nil {
		t.Errorf("parse return error, %s", err)
		return
	}

	expected := Condition(
		&AndCondition{
			Left: &MatchCondition{
				Field: "message",
				Value: "test value",
			},
			Right: &TimeRangeCondition{
				Field:        "datetime",
				Start:        pointer("2020-01-01T00:00:00"),
				End:          pointer("2020-12-31T00:00:00"),
				IncludeStart: true,
				IncludeEnd:   true,
			},
		},
	)

	if diff := deep.Equal(cond, expected); diff != nil {
		t.Errorf("returned condition unexpected: diff= %s", diff)
		return
	}
}

func TestParserB(t *testing.T) {
	cond, err := Parse("number:{1 TO 10}")
	if err != nil {
		t.Errorf("parse return error, %s", err)
		return
	}

	expected := &NumberRangeCondition{
		Field:        "number",
		Start:        pointer("1"),
		End:          pointer("10"),
		IncludeStart: false,
		IncludeEnd:   false,
	}

	if diff := deep.Equal(cond, expected); diff != nil {
		t.Errorf("returned condition unexpected: diff= %s", diff)
		return
	}
}

func TestParserC(t *testing.T) {
	cond, err := Parse("date:{2010-0-01 TO 2020-01-01]")
	if err != nil {
		t.Errorf("parse return error, %s", err)
		return
	}

	expected := &TimeRangeCondition{
		Field:        "date",
		Start:        pointer("2010-0-01"),
		End:          pointer("2020-01-01"),
		IncludeStart: false,
		IncludeEnd:   true,
	}

	if diff := deep.Equal(cond, expected); diff != nil {
		t.Errorf("returned condition unexpected: diff= %s", diff)
		return
	}
}

func TestParserD(t *testing.T) {
	cond, err := Parse("date:2010-01-01")
	if err != nil {
		t.Errorf("parse return error, %s", err)
		return
	}

	expected := &MatchCondition{
		Field: "date",
		Value: "2010-01-01",
	}

	if diff := deep.Equal(cond, expected); diff != nil {
		t.Errorf("returned condition unexpected: diff= %s", diff)
		return
	}
}

func pointer(s string) *string {
	return &s
}
