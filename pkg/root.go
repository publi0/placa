package pkg

import (
	"regexp"

	"github.com/spf13/cobra"
)

var (
	outputFormat string
	isRaw        bool
	IsFull       bool
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "yaml", "Output format (yaml or json)")
	RootCmd.PersistentFlags().BoolVarP(&isRaw, "raw", "r", false, "Print raw output")
	RootCmd.PersistentFlags().BoolVarP(&IsFull, "full", "f", false, "Print full output")
}

var RootCmd = &cobra.Command{
	Use:   "placa [placa]",
	Short: "Consulte a Placa de um veiculo.",
	Long:  "Consulte a Placa de um veiculo e obtenha informações detalhadas sobre ele.",
	Args:  cobra.ExactArgs(1),
	Run:   fetchPlaca,
}

func fetchPlaca(cmd *cobra.Command, args []string) {
	var data any
	placa := args[0]

	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	placa = reg.ReplaceAllString(placa, "")

	if placa == "" {
		cmd.Println("Placa não pode ser vazio.")
		return
	}

	data, err := fetchPlacaData(cmd.Context(), placa)
	if err != nil {
		panic(err)
	}

	if !IsFull {
		g := data.(PlacaData).Dados
		data = ConvertToSimplifiedVehicleInfo(&g)
	}

	switch outputFormat {
	case "json":
		PrintJSON(data, isRaw)
	case "yaml":
		PrintYAML(data, isRaw)
	default:
		cmd.Println("Unsupported output format. Please choose 'json' or 'yaml'.")
	}
}
