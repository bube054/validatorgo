package validatorgo

import (
	"testing"
)

func TestIsAlphaNoIgnoreAndNoLocale(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsAlphaOpts
		want   bool
	}{
		{
			name:   "Has only english alphabets",
			param1: "helloWORLD",
			param2: IsAlphaOpts{},
			want:   true,
		},
		{
			name:   "Has no english alphabets",
			param1: "123456789",
			param2: IsAlphaOpts{},
			want:   false,
		},
		{
			name:   "Has at least one none english alphabet",
			param1: "HELLO89world",
			param2: IsAlphaOpts{},
			want:   false,
		},
		{
			name:   "Has spaces",
			param1: "HELLO world",
			param2: IsAlphaOpts{},
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsAlpha(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}

func TestIsAlphaNoIgnoreAndALocale(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsAlphaOpts
		want   bool
	}{
		{
			name:   "Has only Bulgarian alphabets",
			param1: "Привет",
			param2: IsAlphaOpts{Locale: "bg-BG"},
			want:   true,
		},
		{
			name:   "Has no Bulgarian alphabets and numbers",
			param1: "7889gygu",
			param2: IsAlphaOpts{Locale: "bg-BG"},
			want:   false,
		},
		{
			name:   "Has at least one non bulgarian alphabet",
			param1: "Пfdsr",
			param2: IsAlphaOpts{Locale: "bg-BG"},
			want:   false,
		},
		{
			name:   "Has spaces",
			param1: "При вет",
			param2: IsAlphaOpts{Locale: "bg-BG"},
			want:   false,
		},
		{
			name:   "Invalid Locale defaults to a latin writing system",
			param1: "При вет",
			param2: IsAlphaOpts{Locale: "bg-BE"},
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsAlpha(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}

func TestIsAlphaAnIgnoreAnNoLocale(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsAlphaOpts
		want   bool
	}{
		{
			name:   "Has english alphabets and includes non english alphabets 89 consecutively",
			param1: "Hello89",
			param2: IsAlphaOpts{Ignore: "89"},
			want:   true,
		},
		{
			name:   "Has english alphabets and includes non english alphabets 12 nonconsecutively",
			param1: "1Hello2",
			param2: IsAlphaOpts{Ignore: "12"},
			want:   true,
		},
		{
			name:   "Has english alphabets and includes non english alphabets not ignored",
			param1: "1Hello23",
			param2: IsAlphaOpts{Ignore: "12"},
			want:   false,
		},
		{
			name:   "Has english alphabets and includes non english alphabets that should be escaped",
			param1: ".Hello*)9",
			param2: IsAlphaOpts{Ignore: ".*)9"},
			want:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsAlpha(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}

func TestIsAlphaAnIgnoreAndALocale(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsAlphaOpts
		want   bool
	}{
		{
			name:   "Has bulgarian alphabets and includes non bulgarian alphabets 89 consecutively, which is ignored",
			param1: "Привет89",
			param2: IsAlphaOpts{Ignore: "89", Locale: "bg-BG"},
			want:   true,
		},
		{
			name:   "Has bulgarian alphabets and includes non bulgarian alphabets 12 nonconsecutively, which is ignored",
			param1: "1Привет2",
			param2: IsAlphaOpts{Ignore: "12", Locale: "bg-BG"},
			want:   true,
		},
		{
			name:   "Has bulgarian alphabets and includes non bulgarian alphabets 12 nonconsecutively, which is not ignored",
			param1: "1Привет2",
			param2: IsAlphaOpts{Ignore: "78", Locale: "bg-BG"},
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsAlpha(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
