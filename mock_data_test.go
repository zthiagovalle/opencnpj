package opencnpj

var mockCompanyJSON = `{
  "cnpj": "00000000000000",
  "razao_social": "EMPRESA EXEMPLO LTDA",
  "nome_fantasia": "EXEMPLO",
  "situacao_cadastral": "Ativa",
  "data_situacao_cadastral": "2000-01-01",
  "matriz_filial": "Matriz",
  "data_inicio_atividade": "2000-01-01",
  "cnae_principal": "0000000",
  "cnaes_secundarios": [
    "0000001",
    "0000002"
  ],
  "cnaes_secundarios_count": 2,
  "natureza_juridica": "Sociedade Empresária Limitada",
  "logradouro": "RUA EXEMPLO",
  "numero": "123",
  "complemento": "SALA 1",
  "bairro": "BAIRRO EXEMPLO",
  "cep": "00000000",
  "uf": "SP",
  "municipio": "SAO PAULO",
  "email": "contato@exemplo.com",
  "telefones": [
    {
      "ddd": "11",
      "numero": "900000000",
      "is_fax": false
    }
  ],
  "capital_social": "1000,00",
  "porte_empresa": "Microempresa (ME)",
  "opcao_simples": "",
  "data_opcao_simples": null,
  "opcao_mei": null,
  "data_opcao_mei": null,
  "QSA": [
    {
      "nome_socio": "SOCIO PJ EXEMPLO",
      "cnpj_cpf_socio": "00000000000000",
      "qualificacao_socio": "Sócio Pessoa Jurídica",
      "data_entrada_sociedade": "2000-01-01",
      "identificador_socio": "Pessoa Jurídica",
      "faixa_etaria": "Não se aplica"
    },
    {
      "nome_socio": "SOCIA PF EXEMPLO",
      "cnpj_cpf_socio": "***000000**",
      "qualificacao_socio": "Administrador",
      "data_entrada_sociedade": "2000-01-01",
      "identificador_socio": "Pessoa Física",
      "faixa_etaria": "31 a 40 anos"
    }
  ]
}`
