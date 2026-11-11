# Changelog

All notable changes to `validatorgo`. Versions follow [semantic versioning](https://semver.org).

## v1.1.0 (2026-11-11)

### Performance
- Regexes are compiled once and reused instead of on every call. `IsEmail` went from about 265 µs to about 1 µs per call, `IsCreditCard` from about 385 µs to 3.5 µs, and `IsBic` from about 0.9 ms to 16 µs.

### Fixed
- `IsEmail` no longer panics on input such as `junk [1.2.3.4]`, and no longer accepts text after the address, such as `a@example.com<script>`.
- Validators no longer write defaults into the options struct you pass in. This was a data race when one options value is shared across goroutines.
- `IsURL` no longer reorders the caller's `Protocols` slice.

### Behavior changes
These values passed before and now fail, matching validator.js:
- `IsEmail`: leading, trailing or consecutive dots before the `@`; more than 64 bytes before the `@` or 254 after it (`IgnoreMaxLength` skips this check, which previously did nothing); IP domains such as `user@[127.0.0.1]` unless `AllowIpDomain` is set.
- `IsIP`: leading zeros, such as `01.2.3.4`.
- `IsPort`: leading zeros, such as `080`.
- `IsULID`: extra text around the ID.
- `IsRgbColor`: values over 255, missing values, `rgb()` with four values and `rgba()` with three. `rgba(0,0,0,0)` is now valid.

### Other
- A logo, and CI that runs `gofmt`, `go vet` and race-detector tests on every change.
- Misspelled file names fixed (`isethereumaddress.go`, `iswhitelisted.go`); no exported names changed.

## v1.0.0 (2026-05-07)

### Breaking changes
- Every validator returns `(bool, error)` instead of `bool`. The error is a `*ValidationError` with the validator's name, a human-readable `Message` and a machine-readable `Code` such as `invalid_format`.
- Every options field is a pointer, so leaving a field unset (`nil`) is different from setting it to its zero value. Use the helpers `Bool`, `String`, `Int`, `Uint` and `Float64` to set them, for example `&IsEmailOpts{AllowUTF8LocalPart: Bool(false)}`.

### New
- `IsBTCAddress` accepts bech32, taproot and testnet addresses.

### Fixed
- `IsMobilePhone` patterns for about 40 locales.
- `IsDataURI` accepts valid data URIs as defined by RFC 2397.

### Tests
- Test cases ported from validator.js for 28 validators, and every test now checks the returned error.

## v0.3.2 (2026-05-04)

### Fixed
- `IsISRC` accepts letters and digits in the registrant code.

## v0.3.1 (2025-01-11)

### Fixed
- `IsDate`, `IsBefore`, `IsAfter` and the `ToDate` sanitizer.

## v0.3.0 (2025-01-11)

### Changed
- `IsDate` was rewritten.

## v0.2.0 (2025-01-11)

### New
- `IsArray` and `IsObject`.
- Contributing guide and code of conduct.

## v0.1.1-alpha (2024-11-29)

- Documentation updates.

## v0.1.0-alpha (2024-10-07)

First release: string validators and sanitizers ported from validator.js, including `IsEmail`, `IsURL`, `IsIP`, `IsUUID`, `IsCreditCard`, `IsIBAN`, `IsMobilePhone` and 13 sanitizers.
