package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type Restricoes struct {
	Restricao1          string `json:"restricao_1"`
	Restricao2          string `json:"restricao_2"`
	Restricao3          string `json:"restricao_3"`
	Restricao4          string `json:"restricao_4"`
	LimiteRestricaoTrib string `json:"limite_restricao_trib"`
}

type Multas struct {
	Comentimento string `json:"comentimento"`
	Status       string `json:"status"`
	Valor        string `json:"valor"`
}

type Infracoes struct {
	Tipo     string `json:"tipo"`
	Banco    string `json:"banco"`
	Situacao string `json:"situacao"`
	PagValor string `json:"pagValor"`
}

type Pagamentos struct {
	Infracoes Infracoes `json:"infracoes"`
}

type Faturado struct {
	Faturado   string `json:"faturado"`
	TipoPessoa string `json:"tipo_pessoa"`
	UF         string `json:"uf"`
}

type Ano struct {
	Fabricacao string `json:"fabricacao"`
	Modelo     string `json:"modelo"`
}

type Localidade struct {
	Municipio string `json:"municipio"`
	UF        string `json:"uf"`
}

type Proprietario struct {
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
}

type Dados struct {
	Situacao             string       `json:"situacao"`
	Chassi               string       `json:"chassi"`
	Motor                string       `json:"motor"`
	Renavam              string       `json:"renavam"`
	Placa                string       `json:"placa"`
	SituacaoChassi       string       `json:"situacao_chassi"`
	Combustivel          string       `json:"combustivel"`
	Potencia             string       `json:"potencia"`
	CapacidadeCarga      string       `json:"capacidade_carga"`
	Nacionalidade        string       `json:"nacionalidade"`
	Modelo               string       `json:"modelo"`
	Cor                  string       `json:"cor"`
	Cilindradas          string       `json:"cilindradas"`
	Carroceria           string       `json:"carroceria"`
	CapMaximaTracao      string       `json:"cap_maxima_tracao"`
	PesoBrutoTotal       string       `json:"peso_bruto_total"`
	QuantidadePassageiro string       `json:"quantidade_passageiro"`
	CaixaCambio          string       `json:"caixa_cambio"`
	Aixos                string       `json:"aixos"`
	EixoTraseiroDif      string       `json:"eixo_traseiro_dif"`
	TerceiroEixo         string       `json:"terceiro_eixo"`
	TipoVeiculo          string       `json:"tipo_veiculo"`
	EspecieVeiculo       string       `json:"especie_veiculo"`
	TipoDocImportadora   string       `json:"tipo_doc_importadora"`
	IdentImportadora     string       `json:"ident_importadora"`
	DI                   string       `json:"di"`
	RegistroDI           string       `json:"registro_di"`
	UnidadeLocalSRF      string       `json:"unidade_local_srf"`
	PlacaModeloAntigo    string       `json:"placa_modelo_antigo"`
	PlacaModeloNovo      string       `json:"placa_modelo_novo"`
	PlacaNova            string       `json:"placa_nova"`
	Restricoes           Restricoes   `json:"restricoes"`
    Multas               Multas       `json:"multas"`
	Pagamentos           Pagamentos   `json:"pagamentos"`
	Faturado             Faturado     `json:"faturado"`
	Ano                  Ano          `json:"ano"`
	Localidade           Localidade   `json:"localidade"`
	Proprietario         Proprietario `json:"proprietario"`
}

type PlacaData struct {
	Dados  Dados `json:"dados"`
}

const (
	ConsultaURL = "https://consultanacional.org/wp-json/custom/v1/consulta_placa?placa="
)

func fetchPlacaData(ctx context.Context, placa string) (PlacaData, error) {
	client := &http.Client{}
	placa = strings.ToUpper(placa)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ConsultaURL+placa, nil)
	if err != nil {
		return PlacaData{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return PlacaData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PlacaData{}, errors.New("placa não encontrada")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PlacaData{}, err
	}

	var placaData PlacaData

	json.Unmarshal(body, &placaData)

	return placaData, nil
}
