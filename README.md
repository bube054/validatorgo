# validatorgo

<img alt="Tag" src="https://img.shields.io/badge/tag-v0.1.0-blue?labelColor=gray"> <img alt="Go Version" src="https://img.shields.io/badge/Go->=1.13-00ADD8?labelColor=gray"> <img alt="Reference" src="https://img.shields.io/badge/-reference-00ADD8?logo=go&labelColor=gray"> <img alt="Tests" src="https://img.shields.io/badge/tests-passing-brightgreen?logo=github&labelColor=gray"> <img alt="Go Report" src="https://img.shields.io/badge/go_report-A%2B-00ADD8"> <img alt="Coverage" src="https://img.shields.io/badge/coverage-100%25-brightgreen?logo=codecov"> <img alt="Contributors" src="https://img.shields.io/badge/contributors-125-blueviolet"> <img alt="License" src="https://img.shields.io/badge/license-MIT-yellow">

A library of string validators and sanitizers, based on the js library [validator.js](https://github.com/validatorjs/validator.js)

## Rationale
Why not use popular go libraries like [Package validator](https://github.com/go-playground/validator) or [govalidator](https://github.com/asaskevich/govalidator)? Because they focus on validating structs rather than standalone strings and do not have a robust library of validators as the original js [validator.js](https://github.com/validatorjs/validator.js) package its based on.
I made this project so that can work that i can work on another go open source library called [ginvalidator](https://github.com/bube054/ginvalidator), a go version for the popular node/express library [express-validator](https://express-validator.github.io/docs/) . Seeing that the most popular libraries in the go ecosystem where either overkill, not as robust, or a combination of both, I decided to make my own library.

## Installation
Using go get.
 ```go
  go get github.com/bube054/validatorgo
 ```
Then import the package into your own code.
 ```go
  import (
    "fmt"
    "github.com/bube054/validatorgo"
  )
 ```
If you are unhappy using the long validatorgo, you can do something like this.
 ```go
  import (
    "fmt"
    vgo "github.com/bube054/validatorgo"
  )
 ```
Simple example usage
 ```go
  func main(){
    id := "5f2a6c69e1d7a4e0077b4e6b"
    validId := vgo.IsMongoID(id)
    fmt.Println(validId) // true
  }
 ```

# Validators
Here is a list of the validators currently available.
| s/n | Validator |  description  | 
|:-----|:--------:| :--------|
| 1 | Contains | A validator that checks if the string contains the seed. <mark>ContainsOpt</mark> is a struct that defaults to <mark>{ IgnoreCase: false, MinOccurrences: 1 }</mark>. <br /> Options: <br /> <mark>IgnoreCase</mark>: Ignore case when doing comparison, default false.<br /> <mark>MinOccurrences</mark>: Minimum number of occurrences for the seed in the string. Defaults to 1.|

# Sanitizers
Here is a list of the validators currently available.
| s/n | Sanitizer |  description  | 
|:-----|:--------:| :--------|
|  |  | |

# Maintainers
- [bube054](https://github.com/bube054) - Attah Gbubemi David (author)
- [Yusufsalami](https://github.com/Yusufsalami) - Yusuf Salami

# Other related projects
- [ginvalidator](https://github.com/bube054/ginvalidator)
- [echovalidator](https://github.com/bube054/echovalidator)
- [fibervalidator](https://github.com/bube054/fibervalidator)
- [chivalidator](https://github.com/bube054/chivalidator)

# License
This project is licensed under the [MIT](https://github.com/bube054/validatorgo/blob/master/LICENSE). See the [LICENSE](https://github.com/bube054/validatorgo/blob/master/LICENSE) file for details.