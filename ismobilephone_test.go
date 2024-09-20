package validatorgo

import "testing"

func TestIsMobilePhone(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 []string
		param3 IsMobilePhoneOpts
		want   bool
	}{
		{name: "Valid nigerian number", param1: "08070448986", param2: []string{"en-NG"}, param3: IsMobilePhoneOpts{}, want: true},
		{name: "Invalid nigerian number", param1: "090666666567", param2: []string{"en-NG"}, param3: IsMobilePhoneOpts{}, want: false},
		{name: "Valid nigerian number, strict", param1: "+2348120989668", param2: []string{"en-NG"}, param3: IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid nigerian number, strict", param1: "+23481209895", param2: []string{"en-NG"}, param3: IsMobilePhoneOpts{StrictMode: true}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsMobilePhone(test.param1, test.param2, test.param3)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
