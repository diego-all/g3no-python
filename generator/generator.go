package generator

import (
	"fmt"

	"github.com/diego-all/g3no-python/models"
)

type Attribute struct {
	TipoDato string `json:"tipoDato"`
}

type Entity struct {
	Tipo      string               `json:"tipo"`
	Atributos map[string]Attribute `json:"atributos"`
}

func Generate(projectName string, dbType string, config models.Config, dummy bool) {
	fmt.Printf("Generando proyecto '%s' con base de datos '%s'\n", projectName, dbType)
	fmt.Print("\n")

	fmt.Println("Config from Generate (output python): ", config)

}
