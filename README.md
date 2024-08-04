# Consulte a Placa Veicular pelo terminal

Consulte a Placa pelo terminal

## Instalação

### Go install

```bash
go install github.com/publi0/placa@latest
```

### From source

```bash
go build -o placa main.go
```

### From release

<https://github.com/publi0/placa/releases>

## Uso

```bash
./placa 00000000000000
```

## Saída

### Por padrão o output é o yaml simplificado

```yaml
placa: ABC1234
renavam: "123445"
chassi: 9BR53ZEC148509041
modelo: TOYOTA/COROLLA 
ano:
    fabricacao: "2003"
    modelo: "2004"
cor: Preta
tipo_veiculo: Automovel
proprietario:
    nome: Jose Silva
    cpf: "123456"
localidade:
    municipio: SAO PAULO
    uf: SP
situacao: S
restricoes: []
infracoes: []
combustivel: Gasolina
potencia: "110"
cilindradas: "1598"
peso_bruto_total: "110"
```

## Flags

### --ouput -o

Retorna o output em json ou yaml

```bash
./placa 00000000000000 --output json
```

### --raw -r

Retorna o output em texto puro

```bash
./placa 00000000000000 --raw
```

### --full -f

Retorna o output completo

```bash
./placa 00000000000000 --full
```

```yaml
dados:
    situacao: S
    chassi: string
    motor: string
    renavam: string
    placa: string
    situacao_chassi: string
    combustivel: string
    potencia: string
    capacidade_carga: string
    nacionalidade: string
    modelo: string
    cor: string
    cilindradas: string
    carroceria: string
    cap_maxima_tracao: string
    peso_bruto_total: string
    quantidade_passageiro: string
    caixa_cambio: string
    aixos: string
    eixo_traseiro_dif: string
    terceiro_eixo: string
    tipo_veiculo: string
    especie_veiculo: string
    tipo_doc_importadora: string
    ident_importadora: string
    di: string
    registro_di: string
    unidade_local_srf: string
    placa_modelo_antigo: string
    placa_modelo_novo: string
    placa_nova: string
    restricoes:
        restricao_1: string
        restricao_2: string
        restricao_3: string
        restricao_4: string
        limite_restricao_trib: string
    multas:
        comentimento: string
        status: string
        valor: string
    pagamentos:
        infracoes:
            tipo: string
            banco: string
            situacao: string
            pagValor: string
    faturado:
        faturado: string
        tipo_pessoa: string
        uf: string
    ano:
        fabricacao: string
        modelo: string
    localidade:
        municipio: string
        uf: string
    proprietario:
        nome: string
        cpf: string
```

## Fonte de dados

URL base: <https://consultanacional.org/>
