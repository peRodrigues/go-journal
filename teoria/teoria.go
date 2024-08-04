package teoria

import (
	"fmt"
	arraysslices "github.com/peRodrigues/go-journal/teoria/arrays-slices"
	tiposdedados "github.com/peRodrigues/go-journal/teoria/tipos-de-dados"
	variaveisconstantes "github.com/peRodrigues/go-journal/teoria/variaveis-constantes"
)

func Teoria() {
	fmt.Println("\nTEORIA:")

	tiposdedados.Tiposdedados()
	variaveisconstantes.Variaveisconstantes()
	arraysslices.Arraysslices()
}