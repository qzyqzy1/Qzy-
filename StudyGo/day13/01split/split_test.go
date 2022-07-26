package split

import (
	"reflect"
	"testing"
)

//测试

func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{
		"simple":    test{input: "我爱你", sep: "爱", want: []string{"我", "你"}},
		"mult sep":  test{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"mult sep2": test{input: "abcd", sep: "bc", want: []string{"a", "d"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name： %s failed ,want:%v,got:%v", name, tc.want, got)

			}
		})
	}
}
