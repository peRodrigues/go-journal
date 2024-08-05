package estruturascondicionais

import "fmt"

func Estruturascondicionais() {
	fmt.Println("Estruturas Condicionais:")
	if irapraia() {
		fmt.Println("Você vai à pria!")
	} else {
		fmt.Println("Deixe para ir outro dia...")
	}
}

func irapraia() bool {
	chuva := false
	cansado := 60

	if chuva || cansado >= 60 {
		return false
	}

	return true
}