// Package opencnpj provides a Go client for querying CNPJ data
// through the OpenCNPJ API.
//
// Example usage:
//
//	client := opencnpj.NewClient(5 * time.Second)
//	company, err := client.FindByCNPJ("06990590000123")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(company.LegalName)
package opencnpj

const (
	apiURL = "https://api.opencnpj.org/%s"
)
