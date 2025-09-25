package opencnpj

// Company represents the full structure of a company's data from the API.
// Fields are in English, mapped from the Portuguese JSON via struct tags.
type Company struct {
	CNPJ                   string    `json:"cnpj"`
	LegalName              string    `json:"razao_social"`
	TradeName              string    `json:"nome_fantasia"`
	RegistrationStatus     string    `json:"situacao_cadastral"`
	RegistrationStatusDate string    `json:"data_situacao_cadastral"`
	BranchType             string    `json:"matriz_filial"`
	MainActivityStartDate  string    `json:"data_inicio_atividade"`
	PrimaryCNAE            string    `json:"cnae_principal"`
	SecondaryCNAEs         []string  `json:"cnaes_secundarios"`
	SecondaryCNAEsCount    int       `json:"cnaes_secundarios_count"`
	LegalNature            string    `json:"natureza_juridica"`
	Street                 string    `json:"logradouro"`
	Number                 string    `json:"numero"`
	Complement             string    `json:"complemento"`
	Neighborhood           string    `json:"bairro"`
	CEP                    string    `json:"cep"`
	State                  string    `json:"uf"`
	City                   string    `json:"municipio"`
	Email                  string    `json:"email"`
	Phones                 []Phone   `json:"telefones"`
	ShareCapital           string    `json:"capital_social"`
	CompanySize            string    `json:"porte_empresa"`
	IsOptingForSimples     string    `json:"opcao_simples"`
	SimplesOptionDate      string    `json:"data_opcao_simples"`
	IsMEI                  string    `json:"opcao_mei"`
	MEIOptionDate          string    `json:"data_opcao_mei"`
	Partners               []Partner `json:"qsa"`
}

// Phone represents a phone number structure.
type Phone struct {
	DDD    string `json:"ddd"`
	Number string `json:"numero"`
	IsFax  bool   `json:"is_fax"`
}

// Partner represents a partner or administrator in the company.
type Partner struct {
	Name          string `json:"nome_socio"`
	Document      string `json:"cnpj_cpf_socio"`
	Qualification string `json:"qualificacao_socio"`
	EntryDate     string `json:"data_entrada_sociedade"`
	Identifier    string `json:"identificador_socio"`
	AgeRange      string `json:"faixa_etaria"`
}
