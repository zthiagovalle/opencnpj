[![Status do Build](https://github.com/zthiagovalle/opencnpj/actions/workflows/ci.yml/badge.svg)](https://github.com/zthiagovalle/opencnpj/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/zthiagovalle/opencnpj)](https://goreportcard.com/report/github.com/zthiagovalle/opencnpj)
![Go Version](https://img.shields.io/github/go-mod/go-version/zthiagovalle/opencnpj)
[![License](https://img.shields.io/github/license/zthiagovalle/opencnpj)](https://github.com/zthiagovalle/opencnpj/blob/main/LICENSE)
[![Latest Release](https://img.shields.io/github/v/release/zthiagovalle/opencnpj)](https://github.com/zthiagovalle/opencnpj/releases)

# OpenCNPJ Go Client

A simple and efficient Go client for the [OpenCNPJ API](https://opencnpj.org/).

This package provides a straightforward way to fetch Brazilian company registration data using a CNPJ number. It acts as a wrapper around the free and public OpenCNPJ service, offering typed Go structs for easy integration.

## About the OpenCNPJ API

[OpenCNPJ](https://opencnpj.org/) is a free and public API for querying registration data of Brazilian companies. By providing a valid CNPJ, you receive a simple JSON response, ready for use in applications, integrations, and automations.

- **Official Website:** [https://opencnpj.org/](https://opencnpj.org/)
- **Official Repository:** [https://github.com/Hitmasu/opencnpj](https://github.com/Hitmasu/opencnpj)

## Installation

To install the package, use `go get`:

```bash
go get github.com/zthiagovalle/opencnpj
```

## Usage

```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/zthiagovalle/opencnpj"
)

func main() {
    client := opencnpj.NewClient()
    company, err := client.FindByCNPJ("06990590000123")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Legal Name: %s\n", company.LegalName)
    fmt.Printf("Trade Name: %s\n", company.TradeName)
    fmt.Printf("Registration Status: %s\n", company.RegistrationStatus)
    fmt.Printf("Registration Status Date: %s\n", company.RegistrationStatusDate)
    fmt.Printf("Branch Type: %s\n", company.BranchType)
    fmt.Printf("MainActivity Start Date: %s\n", company.MainActivityStartDate)
    fmt.Printf("Primary CNAE: %s\n", company.PrimaryCNAE)
    fmt.Printf("Legal Nature: %s\n", company.LegalNature)
    fmt.Printf("Street: %s\n", company.Street)
    fmt.Printf("Number: %s\n", company.Number)
    fmt.Printf("Complement: %s\n", company.Complement)
    fmt.Printf("Neighborhood: %s\n", company.Neighborhood)
    fmt.Printf("CEP: %s\n", company.CEP)
    fmt.Printf("State: %s\n", company.State)
    fmt.Printf("City: %s\n", company.City)
    fmt.Printf("Email: %s\n", company.Email)
    fmt.Printf("Share Capital: %s\n", company.ShareCapital)
    fmt.Printf("Company Size: %s\n", company.CompanySize)
}
```

## Options

### WithTimeout

You can set a timeout for requests:

```go
client := opencnpj.NewClient(opencnpj.WithTimeout(10 * time.Second))
```

### WithHTTPClient

You can provide your own `http.Client`:

```go
httpClient := &http.Client{
    Timeout: 5 * time.Second,
}
client := opencnpj.NewClient(opencnpj.WithHTTPClient(httpClient))
```

### WithBaseURL

You can set a custom base URL for the API, which is useful for testing or using a proxy:

```go
client := opencnpj.NewClient(opencnpj.WithBaseURL("https://my-proxy.com/%s"))
```
