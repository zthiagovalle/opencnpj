[![Status do Build](https://github.com/zthiagovalle/opencnpj/actions/workflows/ci.yml/badge.svg)](https://github.com/zthiagovalle/opencnpj/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/zthiagovalle/opencnpj)](https://goreportcard.com/report/github.com/zthiagovalle/opencnpj)
![Go Version](https://img.shields.io/github/go-mod/go-version/zthiagovalle/opencnpj)
[![License](https://img.shields.io/github/license/zthiagovalle/opencnpj)](https://github.com/zthiagovalle/opencnpj/blob/main/LICENSE)
[![Latest Tag](https://img.shields.io/github/v/tag/zthiagovalle/opencnpj)](https://github.com/zthiagovalle/opencnpj/tags)

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
    company, err := client.FindCompanyByCNPJ("06990590000123")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("company: %+v", company)
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

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".

1.  **Fork the Project**
2.  **Create your Feature Branch** (`git checkout -b feature/AmazingFeature`)
3.  **Commit your Changes** (`git commit -m 'feat: Add some AmazingFeature'`)
4.  **Push to the Branch** (`git push origin feature/AmazingFeature`)
5.  **Open a Pull Request**
