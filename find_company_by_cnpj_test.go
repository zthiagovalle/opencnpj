package opencnpj

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCompanyByCNPJ(t *testing.T) {
	t.Run("should return a company when the CNPJ is valid", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockCompanyJSON))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL + "/%s"))

		company, err := client.FindCompanyByCNPJ("00000000000000")

		assert.NoError(t, err)
		assert.NotNil(t, company)
		assert.Equal(t, "00000000000000", company.CNPJ)
		assert.Equal(t, "EMPRESA EXEMPLO LTDA", company.LegalName)
		assert.Equal(t, "EXEMPLO", company.TradeName)
		assert.Equal(t, "Ativa", company.RegistrationStatus)
		assert.Equal(t, "2000-01-01", company.RegistrationStatusDate)
		assert.Equal(t, "Matriz", company.BranchType)
		assert.Equal(t, "2000-01-01", company.MainActivityStartDate)
		assert.Equal(t, "0000000", company.PrimaryCNAE)
		assert.Len(t, company.SecondaryCNAEs, 2)
		assert.Equal(t, 2, company.SecondaryCNAEsCount)
		assert.Equal(t, "Sociedade Empresária Limitada", company.LegalNature)
		assert.Equal(t, "RUA EXEMPLO", company.Street)
		assert.Equal(t, "123", company.Number)
		assert.Equal(t, "SALA 1", company.Complement)
		assert.Equal(t, "BAIRRO EXEMPLO", company.Neighborhood)
		assert.Equal(t, "00000000", company.CEP)
		assert.Equal(t, "SP", company.State)
		assert.Equal(t, "SAO PAULO", company.City)
		assert.Equal(t, "contato@exemplo.com", company.Email)
		assert.Len(t, company.Phones, 1)
		assert.Equal(t, "1000,00", company.ShareCapital)
		assert.Equal(t, "Microempresa (ME)", company.CompanySize)
		assert.Empty(t, company.IsOptingForSimples)
		assert.Empty(t, company.SimplesOptionDate)
		assert.Empty(t, company.IsMEI)
		assert.Empty(t, company.MEIOptionDate)
		assert.Len(t, company.Partners, 2)
	})

	t.Run("should return ErrNotFound when the CNPJ is not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL + "/%s"))

		company, err := client.FindCompanyByCNPJ("00000000000000")

		assert.Nil(t, company)
		assert.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("should return ErrTooManyRequests when the rate limit is exceeded", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTooManyRequests)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL + "/%s"))

		company, err := client.FindCompanyByCNPJ("06990590000123")

		assert.Nil(t, company)
		assert.ErrorIs(t, err, ErrTooManyRequests)
	})

	t.Run("should return an error when the status code is unexpected", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL + "/%s"))

		company, err := client.FindCompanyByCNPJ("06990590000123")

		assert.Nil(t, company)
		assert.Error(t, err)
	})

	t.Run("should return an error when the response is not a valid JSON", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`not a json`))
		}))
		defer server.Close()

		client := NewClient(WithBaseURL(server.URL + "/%s"))

		company, err := client.FindCompanyByCNPJ("06990590000123")

		assert.Nil(t, company)
		assert.Error(t, err)
	})

	t.Run("should return an error when the request fails", func(t *testing.T) {
		client := NewClient(WithBaseURL("invalid-url"))

		company, err := client.FindCompanyByCNPJ("06990590000123")

		assert.Nil(t, company)
		assert.Error(t, err)
	})
}
