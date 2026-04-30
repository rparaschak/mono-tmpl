package testkit

import "github.com/rparaschak/mono-tmpl/api/modules/sample/contracts"

func InputNamed(name string) contracts.SampleInputDTO {
	input := WarsawInput()
	input.Name = name
	return input
}

func WithName(input contracts.SampleInputDTO, name string) contracts.SampleInputDTO {
	input.Name = name
	return input
}

func WarsawInput() contracts.SampleInputDTO {
	return contracts.SampleInputDTO{
		Name:      "warsaw",
		Latitude:  52.2297,
		Longitude: 21.0122,
	}
}

func KrakowInput() contracts.SampleInputDTO {
	return contracts.SampleInputDTO{
		Name:      "krakow",
		Latitude:  50.0497,
		Longitude: 19.9445,
	}
}

func WroclawInput() contracts.SampleInputDTO {
	return contracts.SampleInputDTO{
		Name:      "wroclaw",
		Latitude:  51.1079,
		Longitude: 17.0385,
	}
}

func GdanskInput() contracts.SampleInputDTO {
	return contracts.SampleInputDTO{
		Name:      "gdansk",
		Latitude:  54.3520,
		Longitude: 18.6466,
	}
}

func PoznanInput() contracts.SampleInputDTO {
	return contracts.SampleInputDTO{
		Name:      "poznan",
		Latitude:  52.4064,
		Longitude: 16.9252,
	}
}

func AlphaOneInput() contracts.SampleInputDTO {
	return contracts.SampleInputDTO{
		Name:      "alpha-1",
		Latitude:  52.2297,
		Longitude: 21.0122,
	}
}

func AlphaTwoInput() contracts.SampleInputDTO {
	return contracts.SampleInputDTO{
		Name:      "alpha-2",
		Latitude:  50.0497,
		Longitude: 19.9445,
	}
}

func BetaOneInput() contracts.SampleInputDTO {
	return contracts.SampleInputDTO{
		Name:      "beta-1",
		Latitude:  51.1079,
		Longitude: 17.0385,
	}
}
