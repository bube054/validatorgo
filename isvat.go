package validatorgo

// A validator that checks if the string is a valid VAT number if validation is available for the given country code matching ISO 3166-1 alpha-2.
//
// countryCode is one of ("AL", "AR", "AT", "AU", "BE", "BG", "BO", "BR", "BY", "CA", "CH", "CL", "CO", "CR", "CY", "CZ", "DE", "DK", "DO", "EC", "EE", "EL", "ES", "FI", "FR", "GB", "GT", "HN", "HR", "HU", "ID", "IE", "IL", "IN", "IS", "IT", "KZ", "LT", "LU", "LV", "MK", "MT", "MX", "NG", "NI", "NL", "NO", "NZ", "PA", "PE", "PH", "PL", "PT", "PY", "RO", "RS", "RU", "SA", "SE", "SI", "SK", "SM", "SV", "TR", "UA", "UY", "UZ", "VE"). If none is provided will match every one.
func IsVAT(str, countryCode string) bool {
	return false
}
