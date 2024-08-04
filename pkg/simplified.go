package pkg

type SimplifiedVehicleInfo struct {
	Placa   string `json:"placa"`
	Renavam string `json:"renavam"`
	Chassi  string `json:"chassi"`
	Modelo  string `json:"modelo"`
	Ano     struct {
		Fabricacao string `json:"fabricacao"`
		Modelo     string `json:"modelo"`
	} `json:"ano"`
	Cor          string `json:"cor"`
	TipoVeiculo  string `json:"tipo_veiculo"`
	Proprietario struct {
		Nome string `json:"nome"`
		CPF  string `json:"cpf"`
	} `json:"proprietario"`

	Localidade struct {
		Municipio string `json:"municipio"`
		UF        string `json:"uf"`
	} `json:"localidade"`
	Situacao   string   `json:"situacao"`
	Restricoes []string `json:"restricoes"`
	Infracoes  []struct {
		Tipo     string `json:"tipo"`
		Situacao string `json:"situacao"`
		Valor    string `json:"valor"`
	} `json:"infracoes"`
	Combustivel    string `json:"combustivel"`
	Potencia       string `json:"potencia"`
	Cilindradas    string `json:"cilindradas"`
	PesoBrutoTotal string `json:"peso_bruto_total"`
}

func ConvertToSimplifiedVehicleInfo(original *Dados) SimplifiedVehicleInfo {
    simplified := SimplifiedVehicleInfo{
        Placa:          original.Placa,
        Renavam:        original.Renavam,
        Chassi:         original.Chassi,
        Modelo:         original.Modelo,
        Cor:            original.Cor,
        TipoVeiculo:    original.TipoVeiculo,
        Situacao:       original.Situacao,
        Combustivel:    original.Combustivel,
        Potencia:       original.Potencia,
        Cilindradas:    original.Cilindradas,
        PesoBrutoTotal: original.PesoBrutoTotal,
    }

    simplified.Ano.Fabricacao = original.Ano.Fabricacao
    simplified.Ano.Modelo = original.Ano.Modelo

    simplified.Proprietario.Nome = original.Proprietario.Nome
    simplified.Proprietario.CPF = original.Proprietario.CPF

    simplified.Localidade.Municipio = original.Localidade.Municipio
    simplified.Localidade.UF = original.Localidade.UF

    if original.Restricoes.LimiteRestricaoTrib != "" {
        simplified.Restricoes = []string{
            original.Restricoes.Restricao1,
            original.Restricoes.Restricao2,
            original.Restricoes.Restricao3,
            original.Restricoes.Restricao4,
        }
        simplified.Restricoes = removeEmptyStrings(simplified.Restricoes)
    }

    if original.Pagamentos.Infracoes.Situacao != "" && original.Pagamentos.Infracoes.Situacao != "" {
        infracao := struct {
            Tipo     string `json:"tipo"`
            Situacao string `json:"situacao"`
            Valor    string `json:"valor"`
        }{
            Tipo:     original.Pagamentos.Infracoes.Tipo,
            Situacao: original.Pagamentos.Infracoes.Situacao,
            Valor:    original.Pagamentos.Infracoes.PagValor,
        }
        simplified.Infracoes = append(simplified.Infracoes, infracao)
    }

    if original.Multas.Status != "" {
        multa := struct {
            Tipo     string `json:"tipo"`
            Situacao string `json:"situacao"`
            Valor    string `json:"valor"`
        }{
            Tipo:     "Multa",
            Situacao: original.Multas.Status,
            Valor:    original.Multas.Valor,
        }
        simplified.Infracoes = append(simplified.Infracoes, multa)
    }

    return simplified
}

func removeEmptyStrings(s []string) []string {
    var r []string
    for _, str := range s {
        if str != "" {
            r = append(r, str)
        }
    }
    return r
}