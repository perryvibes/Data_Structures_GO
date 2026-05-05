package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Masina struct {
	id        int
	nrUsi     int
	pret      float64
	model     string
	numeSofer string
	serie     rune
}

type Nod struct {
	masina *Masina
	next   *Nod
}

// Functia de citire Masina din fisier pentru toate exercitiile

func citireMasiniDinFisier(fptr *os.File) []Masina {
	var masini []Masina
	var masinaScanata Masina
	scanner := bufio.NewScanner(fptr)
	for scanner.Scan() {
		line := scanner.Text()
		detaliiMasina := strings.Split(line, ",")
		masinaScanata.id, _ = strconv.Atoi(detaliiMasina[0])
		masinaScanata.nrUsi, _ = strconv.Atoi(detaliiMasina[1])
		masinaScanata.pret, _ = strconv.ParseFloat(detaliiMasina[2], 64)
		masinaScanata.model = detaliiMasina[3]
		masinaScanata.numeSofer = detaliiMasina[4]
		masinaScanata.serie = []rune(detaliiMasina[5])[0]
		masini = append(masini, masinaScanata)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	return masini
}

// Functia de afisare Masina
func afisareMasina(masina Masina) {
	fmt.Printf("Id: %d\n", masina.id)
	fmt.Printf("Nr. usi : %d\n", masina.nrUsi)
	fmt.Printf("Pret: %.2f\n", masina.pret)
	fmt.Printf("Model: %s\n", masina.model)
	fmt.Printf("Nume sofer: %s\n", masina.numeSofer)
	fmt.Printf("Serie: %c\n\n", masina.serie)
}

// Functia de afisare lista de masini

func afisareListaMasini(lista *Nod) {
	for lista != nil {
		afisareMasina(*lista.masina)
		lista = lista.next
	}
}

// adaugare la final
func adaugaMasinaInListaFinal(cap **Nod, masinaNoua Masina) {
	nou := &Nod{
		masina: &masinaNoua,
		next:   nil,
	}
	if (*cap) == nil {
		(*cap) = nou
	} else {
		p := *cap
		for p.next != nil {
			p = p.next
		}
		p.next = nou
	}
}

// adaugare la inceput
func adaugaMasinaInListaInceput(cap **Nod, masinaNoua Masina) {
	nou := &Nod{
		masina: &masinaNoua,
		next:   nil,
	}
	if (*cap) == nil {
		*cap = nou
	} else {
		nou.next = *cap
		*cap = nou
	}
}

func citireListaMasinaDinFisier(numeFisier string) *Nod {
	file, err := os.Open(numeFisier)
	var lista *Nod = nil
	if err != nil {
		fmt.Println("Error opening file!")
	}
	for _, value := range citireMasiniDinFisier(file) {
		adaugaMasinaInListaFinal(&lista, value)
	}
	defer file.Close()
	return lista
}

func calculeazaPretMediu(lista *Nod) float64 {
	var sum float64
	var count float64
	for lista != nil {
		sum += lista.masina.pret
		count++
		lista = lista.next
	}
	return sum / count
}

func stergeMasiniDinSerie(lista **Nod, idCautat int) {
	if (*lista) == nil {
		return
	}
	for *lista != nil && (*lista).masina.id == idCautat {
		// go nu are deallocate deci vom rupe legaturile nodului si GC se va ocupa de dezalocare
		*lista = (*lista).next
	}
	if *lista != nil {
		var nodAnterior *Nod = *lista
		var nodCurent *Nod = (*lista).next
		for nodCurent != nil {
			if nodCurent.masina.id == idCautat {
				nodAnterior.next = nodCurent.next
				nodCurent = nodCurent.next
			} else {
				nodAnterior = nodCurent
				nodCurent = nodCurent.next
			}
		}
	}
}

func calcularePretMasinaUnuiSofer(lista *Nod, nume string) float64 {
	var pretTotal float64 = 0
	for lista != nil {
		if lista.masina.numeSofer == nume {
			pretTotal += lista.masina.pret
		}
		lista = lista.next
	}
	return pretTotal
}

func getCeaMaiScumpaMasina(lista *Nod) string {
	if lista == nil {
		return "N/A"
	} else {
		var nodMax *Nod = lista
		for lista != nil {
			if lista.masina.pret > nodMax.masina.pret {
				nodMax = lista
			}
			lista = lista.next
		}
		return nodMax.masina.model
	}
}

// Functia de dezalocare nu o mai construim deoarece GO are Garbage Collector.

func main() {

	cap := citireListaMasinaDinFisier("cars.txt")
	afisareListaMasini(cap)
	fmt.Println("=========================")
	fmt.Printf("Pretul mediu: %.2f\n", calculeazaPretMediu(cap))
	fmt.Println("=========================")
	stergeMasiniDinSerie(&cap, 10)
	afisareListaMasini(cap)
	fmt.Println("=========================")
	fmt.Printf("Cea mai scumpa masina este: %s\n", getCeaMaiScumpaMasina(cap))
	fmt.Println("=========================")
}
