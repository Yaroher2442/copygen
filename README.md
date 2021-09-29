# Copygen

[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge&logo=appveyor&logo=appveyor)](https://pkg.go.dev/github.com/switchupcb/copygen)
[![Go Report Card](https://goreportcard.com/badge/github.com/switchupcb/copygen?style=for-the-badge)](https://goreportcard.com/report/github.com/switchupcb/copygen)
[![MIT License](https://img.shields.io/github/license/switchupcb/copygen.svg?style=for-the-badge)](https://github.com/switchupcb/copygen/blob/main/LICENSE)

Copygen is a command-line [code generator](https://github.com/gophersgang/go-codegen) that generates type-to-type and field-to-field struct code without adding any reflection or dependencies to your project. Manual-copy code generated by copygen is [**391x faster**](https://github.com/gotidy/copy#benchmark) than [jinzhu/copier](https://github.com/jinzhu/copier), and adds no allocation to your program. Copygen is the most customizable type-copy generator to-date and features a rich yet simple setup inspired by [goverter](https://github.com/jmattheis/goverter).

| Topic                           | Categories                                                                         |
| :------------------------------ | :--------------------------------------------------------------------------------- |
| [Usage](#Usage)                 | [Types](#types), [Setup](#setup), [Command Line](#command-line), [Output](#output) |
| [Customization](#customization) | [Templates](#templates)                                                            |
| [Matcher](#matcher)             | [Automatch](#automatch)                                                            |

## Usage

Each example has a **README**.

| Example                                                                         | Description                                                       |
| :------------------------------------------------------------------------------ | :---------------------------------------------------------------- |
| main                                                                            | The default example.                                              |
| deepcopy _(Roadmap Feature)_                                                    | Uses templates to create a deepcopy.                              |
| [automatch](https://github.com/switchupcb/copygen/tree/main/examples/automatch) | Uses the automatch feature with depth _(doesn't require fields)_. |
| [error](https://github.com/switchupcb/copygen/tree/main/examples/error)         | Uses templates to return an error (temporarily unsupported).      |

**NOTE: The following guide is set for v0.2 ([view v0.1](https://github.com/switchupcb/copygen/tree/v0.1.0))**

This [example](https://github.com/switchupcb/copygen/blob/main/examples/main) uses three type-structs to generate the `ModelsToDomain()` function.

### Types

`./domain/domain.go`

```go
// The domain package contains business logic models.
package domain

// Account represents a user account.
type Account struct {
	ID     int
	UserID string
	Name   string
	Other  string // The other field is not used.
}
```

`./models/model.go`

```go
// The models package contains data storage models (i.e database).
package models

// Account represents the data model for account.
type Account struct {
	ID       int
	Name     string
	Password string
	Email    string
}

// A user represents the data model for a user.
type User struct {
	ID       int
	Name     int
	UserData string
}
```

### Setup

Setting up copygen is a 2-step process involving a `YML` and `GO` file.

**setup.yml**

```yml
# Define where the code will be generated.
generated:
  setup: ./setup.go
  output: ./copygen.go
  package: copygen

# Define the optional custom templates used to generate the file.
# Templates are currently unsupported.
templates:
  header: ./templates/header.go
  function: ./templates/function.go
```

**setup.go**

Create an interface in the specified setup file with a `type Copy interface`. In each function, specify _the types you want to copy from_ as parameters, and _the type you want to copy to_ as return values.

```go
/* Define the functions that will be generated in the `Copy` Interface. */
type Copy interface {
  // custom: see table below for options
  ModelsToDomain(models.Account, models.User) domain.Account
}
```

You can specify options for your functions using comments.

| Option                                           | Use                                 | Description                                                                                                                                                                                                                                                    | Example                                                                        |
| :----------------------------------------------- | :---------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :----------------------------------------------------------------------------- |
| `alloc`                                          | Allocate and return a new object(s) | Copygen uses no allocation by default which means that <br /> fields are assigned to _objects passed as parameters_. <br /> Use `alloc` to return a new copy of your to-types.                                                                                 | `alloc`                                                                        |
| `map from to`                                   | Manual Field Mapping                | Copygen uses its [automatcher](#automatch) default. Override this using `map` <br /> which uses _regex_ to identify fields that will be mapped to and from eachother.                                                                                          | `map .* package.Type.Field` <br /> `map models.Account.ID domain.Account.ID` |
| `depth field level`                             | Use a specific field depth.         | When fields have fields, Copygen uses the full-field depth by default. <br /> Override this using `depth` which uses a _regex string_ and [depth-level](#depth) integer.                                                                                       | `depth .* 2` <br /> `depth models.Account.* 1`                               |
| `deepcopy field`                                 | Deepcopy from-fields.               | Copygen shallow copies fields by default.  Use `deepcopy` to override this. <br /> For more information, view [Shallow Copy vs. Deep Copy](#shallow-copy-vs-deep-copy).                                                                                        | `deepcopy package.Type.Field` <br /> `deepcopy .*` _(deepcopies all fields)_   |
| `custom: option` <br /> `custom: option: option` | Specify custom options.             | You may want to use custom options in your [templates](#templates). <br /> Options of your choice are passed to the generator object using `custom`. <br /> Custom options are parsed as trim-spaced strings. <br /> Each `:` adds an additional key to a map. | `custom: ignore` <br /> `custom: swap: false`                                  |

_[View a reference on Regex.](https://cheatography.com/davechild/cheat-sheets/regular-expressions/)_

#### Convert

In certain cases, you may want to specify a how a specific field is copied with a function. This can be done by defining a function with a `convert` option.
```go
/* Define the fields this converter is applied to using regex. CONVERTERS ARE ONLY APPLIED TO VALID FIELDS. */
// convert: models.User.ID
// comment: Itoa converts an integer to an ascii value.
func Itoa(i int) string {
  return strconv.Itoa(i)
}
```

### Command Line

Install the command line utility. Copygen is an executable and not a dependency, so use `go install`.

```
go install github.com/switchupcb/copygen@latest
```

Install a specific version by specifying a tag version.
```
go install github.com/switchupcb/copygen@v0.0.0
```

Run the executable with given options.

```bash
# Specify the .yml configuration file.
copygen -yml path/to/yml
```

_The path to the YML file is specified in reference to the current working directory._

### Output

This example outputs a `copygen.go` file with the specified imports and functions.

```go
// Code generated by github.com/switchupcb/copygen
// DO NOT EDIT.

package copygen

import (
	"github.com/switchupcb/copygen/examples/main/converter"
	"github.com/switchupcb/copygen/examples/main/domain"
	"github.com/switchupcb/copygen/examples/main/models"
)

// ModelsToDomain copies a User, Account to a Account.
func ModelsToDomain(tA *domain.Account, fU models.User, fA models.Account) {
	// Account fields
	tA.UserID = c.Itoa(fU.ID)
	tA.ID = fA.ID
	tA.Name = fA.Name

}
```

## Customization

The [error example](https://github.com/switchupcb/copygen/blob/main/examples/main) modifies the .yml to use **custom functions** which `return error`. This is done by modifying the .yml and creating **custom template files**.

#### Templates

Templates can be created using **Go** to customize the generated code algorithm. The `copygen` generator uses the `package tenplates` `Header(*models.Generator)` to generate header code and `Function(*models.Function)` to generate code for each function. As a result, these _(package templates with functions)_ are **required** for your templates to work. View [models.Generator](https://github.com/switchupcb/copygen/blob/main/cli/models/function.go) and [models.Function](https://github.com/switchupcb/copygen/blob/main/cli/models/function.go) for context on the parameters passed to each function. Templates are interpreted by [yaegi](https://github.com/traefik/yaegi) which has limitations on module imports _(that are being fixed)_: As a result, **templates are temporarily unsupported.**

## Matcher

Copygen provides two ways to configure fields: **Manually** and the **Automatcher**. Matching is specified in a `.go` file _(which functions as a schema in relation to other generators)_. Tags are complicated to use with other generators which is why they aren't used.

### Automatch

When fields aren't specified using options, copygen will attempt to automatch type-fields by name. Automatch **supports field-depth** (where types are located within fields) **and recursive types** (where the same type is in another type). Automatch loads types from Go modules _(in GOPATH)_. Ensure your modules are up to date by using `go get -u <insert/module/import/path>`.

#### Depth

A depth level of 0 will match the first-level fields. Increasing the depth level will match more fields.

```go
// depth level
type Account
  // 0
  ID      int
  Name    string
  Email   string
  Basic   domain.T // int
  User    domain.DomainUser
              // 1
              UserID   string
              Name     string
              UserData map[string]interface{}
  // 0
  Log     log.Logger
              // 1
              mu      sync.Mutex
                          // 2
                          state   int32
                          sema    uint32
              // 1
              prefix  string
              flag    int
              out     io.Writer
                          // 2
                          Write   func(p []byte) (n int, err error)
              buf     []byte
```

## Optimization 

### Shallow Copy vs. Deep Copy
The library generates a [shallow copy](https://en.m.wikipedia.org/wiki/Object_copying#Shallow_copy) by default. An easy way to deep-copy fields with the same return type is by using `new()` as/in a converter function or by using a custom template.

### Pointers
Go parameters are _pass-by-value_ which means that a parameter's value _(i.e int, memory address, etc)_ is copied into another location of memory. As a result, passing pointers to functions is more efficient **if the byte size of a pointer is less than the total byte size of the struct member's references**. However, be advised that doing so adds memory to the heap _[which can result in less performance](https://medium.com/@vCabbage/go-are-pointers-a-performance-optimization-a95840d3ef85)_. For more information regarding the use of pointers, read [Pointers vs. Values in Parameters and Return Values](https://stackoverflow.com/questions/23542989/pointers-vs-values-in-parameters-and-return-values/23551970#23551970). For more information on memory, read this article: [What Every Programmer Should Know About Memory](https://lwn.net/Articles/250967/).

## Contributing

You can contribute to this repository by viewing the [Project Structure, Code Specifications, and Roadmap](CONTRIBUTING.md).