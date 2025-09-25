// Package opencnpj provides a Go client for querying CNPJ data
// through the OpenCNPJ API.
//
// Example usage:
//
// client := opencnpj.NewClient()
// company, err := client.FindCompanyByCNPJ("06990590000123")
//
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Printf("company: %+v", company)
package opencnpj

const (
	apiURL = "https://api.opencnpj.org/%s"
)
