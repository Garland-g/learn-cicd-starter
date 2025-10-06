package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	TESTS := []struct {
		name   string
		input  string
		output any
	}{
		{
			name:   "None",
			input:  "",
			output: ErrNoAuthHeaderIncluded,
		},
		{
			name:   "Malformed",
			input:  "bad",
			output: errors.New("malformed authorization header"),
		},
		{
			name:   "Good",
			input:  "ApiKey 458tg7h45",
			output: "458tg7h45",
		},
	}

	for _, testcase := range TESTS {
		t.Run(testcase.name, func(t *testing.T) {
			headers := http.Header{}
			headers.Add("Authorization", testcase.input)
			out, err := GetAPIKey(headers)
			if err != nil {
				testErr, ok := testcase.output.(error)
				if !ok {
					panic("bad test")
				}
				if testErr.Error() != err.Error() {
					t.Fatalf("Expected: %v, Got: %v", testErr, err)
				}
			} else if out != testcase.output {
				t.Fatalf("Expected: %v, Got: %v", testcase.output, out)
			}
		})
	}
}

func TestFailure(t *testing.T) {
	t.Fail()
}
