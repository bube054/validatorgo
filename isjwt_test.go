package validatorgo

import "testing"

func TestIsJWT(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid JWT
		{name: "Valid JWT with HS256", param1: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", want: true},
		{name: "Valid JWT with RS256", param1: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMjM0NSIsImV4cCI6MTYwNjMwOTg5OH0.DQOqtr6lUVxH3nIRcfT7jq0MI5QBSwyFgxj1gPP6U5A", want: true},
		{name: "Valid JWT with custom claims", param1: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImphbmVkbyIsImFkbWluIjp0cnVlfQ.VCkOUqBwdp4hXdZoBh0F4WtYXRVZtFD93PcH6ozZ1Qs", want: true},
		{name: "Valid JWT with loggedInAs claim", param1: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dnZWRJbkFzIjoiYWRtaW4iLCJpYXQiOjE0MjI3Nzk2Mzh9.gzSraSYS8EXBxLN_oWnFSRgCzcmJmMjLiuyu5CSpyHI", want: true},
		{name: "Valid JWT with lorem claim", param1: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb3JlbSI6Imlwc3VtIn0.ymiJSsMJXR6tMSr8G9usjQ15_8hKPDv_CArLhxw28MI", want: true},
		{name: "Valid JWT with array claim", param1: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkb2xvciI6InNpdCIsImFtZXQiOlsibG9yZW0iLCJpcHN1bSJdfQ.rRpe04zbWbbJjwM43VnHzAboDzszJtGrNsUxaqQ-GQ8", want: true},
		{name: "Valid JWT with nested object claims", param1: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqb2huIjp7ImFnZSI6MjUsImhlaWdodCI6MTg1fSwiamFrZSI6eyJhZ2UiOjMwLCJoZWlnaHQiOjI3MH19.YRLPARDmhGMC3BBk_OhtwwK21PIkVCqQe8ncIRPKo-E", want: true},
		// Invalid JWT
		{name: "Missing periods (.)", param1: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", want: false},
		{name: "Improperly Base64-encoded payload", param1: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9%lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", want: false},
		{name: "No signature", param1: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ", want: false},
		{name: "Invalid JSON structure in payload", param1: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwLCJuYW1lIjoiSm9obiBEb2UifQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", want: false},
		{name: "Only header no dots", param1: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", want: false},
		{name: "Invalid chars with dollar sign", param1: "$Zs.ewu.su84", want: false},
		{name: "Invalid chars with special symbols", param1: "ks64$S/9.dy$§kz.3sd73b", want: false},
		{name: "Empty string", param1: "", want: false},
		{name: "Only two parts missing signature", param1: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NSIsIm5hbWUiOiJKb2huIERvZSIsImlhdCI6MTUxNjIzOTAyMn0", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsJWT(test.param1)

			assertValidation(t, result, test.want, err)
		})
	}
}
