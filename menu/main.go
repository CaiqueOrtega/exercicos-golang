package main

import (
	"fmt"
	"strings"
)

func main() {
	var choice int

	for choice != 10 {
		fmt.Println("_________________________")
		fmt.Println("           MENU          ")
		fmt.Println("_________________________")
		fmt.Println("1. Two-Fer               ")
		fmt.Println("2. Troca                 ")
		fmt.Println("3. Ajuste Salarial       ")
		fmt.Println("4. FizzBuzz              ")
		fmt.Println("5. Area da Circunferência")
		fmt.Println("6. Soma dos quadrados    ")
		fmt.Println("7. Quadrado da soma      ")
		fmt.Println("8. Palíndromo            ")
		fmt.Println("9. ScrabbleScore         ")
		fmt.Println("_________________________")
		fmt.Println("10. Sair do Sistema       ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Erro ao ler entrada:", err)
			continue
		}

		switch choice {
		case 1:
			TwoFer()
		case 2:
			Troca()
		case 3:
			AjusteSalarial()
		case 4:
			FizzBuzz()
		case 5:
			AreaDaCircunferencia()
		case 6:
			SomaDosQuadrados()
		case 7:
			QuadradoDaSoma()
		case 8:
			Palindromo()
		case 9:
			ScrabbleScore()
		case 10:
			fmt.Println("Saindo do Sistema")
		default:
			fmt.Println("Opção Inválida")
		}
	}

}

func TwoFer() {
	var name string

	fmt.Println("______________________")
	fmt.Println("        Two-Fer       ")
	fmt.Println("______________________")
	fmt.Println(">Entre com um Nome:")
	fmt.Scan(&name)

	if name == "" {
		name = "você"
	}

	fmt.Println("Um para", name, "um para mim")
}

func Troca() {
	var a, b string

	fmt.Println("______________________")
	fmt.Println("        Troca         ")
	fmt.Println("______________________")
	fmt.Println(">Entre com primeiro valor:")
	fmt.Scan(&a)
	fmt.Println(">Entre com segundo valor:")
	fmt.Scan(&b)

	a, b = b, a

	fmt.Printf("Primeiro valor %s, segundo valor %s \n", a, b)
}

func AjusteSalarial() {
	var salario float32
	var reajuste float32

	fmt.Println("______________________")
	fmt.Println("   Ajuste Salarial    ")
	fmt.Println("______________________")
	fmt.Println(">Entre com o salário:")
	fmt.Scan(&salario)
	fmt.Println(">Entre com o percentual do reajuste:")
	fmt.Scan(&reajuste)
	reajuste = reajuste / 100
	salario += salario * reajuste

	fmt.Printf("O valor do salário com reajuste é de: %0.2f \n", salario)
}

func FizzBuzz() {
	var inicial int8
	var final int8

	fmt.Println("__________________________________________")
	fmt.Println("                Fizz Buzz                 ")
	fmt.Println("__________________________________________")
	fmt.Println(">Entre com o número inicial do intervalo:")
	fmt.Scan(&inicial)
	fmt.Println(">Entre com o número final do intervalo:")
	fmt.Scan(&final)

	for i := inicial; i < final; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func AreaDaCircunferencia() {
	var raio float32

	fmt.Println("__________________________________________")
	fmt.Println("          area da Circunferência          ")
	fmt.Println("__________________________________________")
	fmt.Println(">Entre com o valor do raio:")
	fmt.Scan(&raio)

	area := raio * raio * 3.14159265

	fmt.Printf("A área da circunferência com raio %.2f é %.2f\n", raio, area)
}

func SomaDosQuadrados() {
	var valor1, valor2, valor3 int

	fmt.Println("__________________________________________")
	fmt.Println("            Soma dos Quadrados            ")
	fmt.Println("__________________________________________")
	fmt.Println(">Entre com o primeiro valor:")
	fmt.Scan(&valor1)
	fmt.Println(">Entre com o segundo valor:")
	fmt.Scan(&valor2)
	fmt.Println(">Entre com o terceiro valor:")
	fmt.Scan(&valor3)

	somaQuadrados := valor1*valor1 + valor2*valor2 + valor3*valor3

	fmt.Printf("A soma dos quadrados dos valores %d, %d e %d é: %d\n", valor1, valor2, valor3, somaQuadrados)
}

func QuadradoDaSoma() {
	var valor1, valor2, valor3 int

	fmt.Println("__________________________________________")
	fmt.Println("           Quadrado da Soma               ")
	fmt.Println("__________________________________________")
	fmt.Println(">Entre com o primeiro valor:")
	fmt.Scan(&valor1)
	fmt.Println(">Entre com o segundo valor:")
	fmt.Scan(&valor2)
	fmt.Println(">Entre com o terceiro valor:")
	fmt.Scan(&valor3)

	soma := valor1 + valor2 + valor3
	quadradoSoma := soma * soma

	fmt.Printf("O quadrado da soma dos valores %d, %d e %d é: %d\n", valor1, valor2, valor3, quadradoSoma)
}

func Palindromo() {
	var palavra string

	fmt.Println("____________________")
	fmt.Println("      Palíndromo    ")
	fmt.Println("____________________")
	fmt.Println(">Entre com uma palavra:")
	fmt.Scan(&palavra)

	palavra = strings.ToLower(palavra)
	palavra = strings.ReplaceAll(palavra, " ", "")

	isPalindromo := true
	n := len(palavra)

	for i := 0; i < n; i++ {
		if palavra[i] != palavra[n-1-i] {
			isPalindromo = false

			break
		}
	}

	if isPalindromo {
		fmt.Printf("'%s' é um palíndromo!\n", palavra)
	} else {
		fmt.Printf("'%s' não é um palíndromo.\n", palavra)
	}
}

func ScrabbleScore() {

	scores := map[rune]int{
		'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'l': 1, 'n': 1, 'r': 1, 's': 1, 't': 1,
		'd': 2, 'g': 2,
		'b': 3, 'c': 3, 'm': 3, 'p': 3,
		'f': 4, 'h': 4, 'v': 4, 'w': 4, 'y': 4,
		'k': 5,
		'j': 8, 'x': 8,
		'q': 10, 'z': 10,
	}

	var palavra string

	fmt.Println("____________________________________")
	fmt.Println("         Scrabble Score             ")
	fmt.Println("____________________________________")
	fmt.Print("> Entre com a palavra: ")
	fmt.Scan(&palavra)

	palavra = strings.ToLower(palavra)
	palavra = strings.ReplaceAll(palavra, " ", "")

	score := 0
	for _, char := range palavra {
		if val, ok := scores[char]; ok {
			score += val
		}
	}

	fmt.Printf("Scrabble Score de '%s': %d\n", palavra, score)
}
