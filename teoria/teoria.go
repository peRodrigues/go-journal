package teoria

import (
	"fmt"

	"github.com/peRodrigues/go-journal/teoria/arrays-slices"
	estruturascondicionais "github.com/peRodrigues/go-journal/teoria/estruturas-condicionais"
	"github.com/peRodrigues/go-journal/teoria/structs"
	"github.com/peRodrigues/go-journal/teoria/tipos-de-dados"
	"github.com/peRodrigues/go-journal/teoria/variaveis-constantes"
)

func Teoria() {
	fmt.Println("\nTEORIA:")

	tiposdedados.Tiposdedados()
	variaveisconstantes.Variaveisconstantes()
	arraysslices.Arraysslices()
	structs.Structs()
	estruturascondicionais.Estruturascondicionais()
}