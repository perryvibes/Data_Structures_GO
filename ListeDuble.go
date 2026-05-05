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
	masina Masina
	prev   *Nod
	next   *Nod
}

type ListDubla struct {
	// nu vom folosi si 'nrNoduri' pentru simplificare
	head *Nod
	tail *Nod
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

func afisareListaMasiniInceput(lista ListDubla) {
	for lista.head != nil {
		afisareMasina(lista.head.masina)
		lista.head = lista.head.next
	}
}
func afisareListaMasiniSfarsit(lista ListDubla) {
	for lista.tail != nil {
		afisareMasina(lista.tail.masina)
		lista.tail = lista.tail.prev
	}
}

func adaugareMasinaInListaSfarsit(lista *ListDubla, masinaNoua Masina) {
	if lista == nil {
		return
	} else {
		var nou *Nod = &Nod{
			masina: masinaNoua,
			prev:   lista.tail,
			next:   nil,
		}
		if lista.tail != nil {
			lista.tail.next = nou
		} else {
			lista.head = nou
		}
		lista.tail = nou
	}
}
func adaugareMasinaInListaInceput(lista *ListDubla, masinaNoua Masina) {
	if lista == nil {
		return
	} else {
		var nou *Nod = &Nod{
			masina: masinaNoua,
			prev:   nil,
			next:   lista.head,
		}
		if lista.head != nil {
			lista.head.prev = nou
		} else {
			lista.tail = nou
		}
		lista.head = nou
	}
}

func citireLDMasiniDinFisier(numeFisier string) ListDubla {
	file, err := os.Open(numeFisier)
	if err != nil {
		fmt.Println("Eroare deschidere fisier!")
	}
	var listaDubla ListDubla = ListDubla{head: nil, tail: nil}
	for _, value := range citireMasiniDinFisier(file) {
		adaugareMasinaInListaSfarsit(&listaDubla, value)
	}
	defer file.Close()
	return listaDubla
}

func stergeMasinaDupaId(lista *ListDubla, idCautat int) {
	if lista == nil {
		return
	}
	var p *Nod = (*lista).head
	for p != nil && p.masina.id != idCautat {
		p = p.next
	}
	if p == nil {
		return
	}
	// caz pentru primul nod
	if p.prev == nil {
		lista.head = p.next
		if lista.head != nil {
			lista.head.prev = nil
		}
	} else {
		p.prev.next = p.next
	}
	// caz pentru ultimul nod
	if p.next == nil {
		lista.tail = p.prev
		if lista.tail != nil {
			lista.tail.next = nil
		}
	} else {
		p.next.prev = p.prev
	}
}

func calculeazaPretMediu(lista ListDubla) float64 {
	var sum float64
	var counter int
	for lista.head != nil {
		sum += lista.head.masina.pret
		counter++
		lista.head = lista.head.next
	}
	return sum / float64(counter)
}

func getNumeSoferMasinaScumpa(lista ListDubla) string {
	if lista.head == nil {
		return "N/A"
	} else {
		var nodMax *Nod = lista.head
		for lista.head != nil {
			if lista.head.masina.pret > nodMax.masina.pret {
				nodMax = lista.head
			}
			lista.head = lista.head.next
		}
		return nodMax.masina.numeSofer
	}
}

func main() {

	var listaDubla ListDubla = citireLDMasiniDinFisier("cars.txt")
	afisareListaMasiniInceput(listaDubla)
	fmt.Println("\n==========================\n")
	stergeMasinaDupaId(&listaDubla, 10)
	afisareListaMasiniSfarsit(listaDubla)
	fmt.Println("\n==========================\n")
	fmt.Printf("Pretul mediu este: %.2f", calculeazaPretMediu(listaDubla))
	fmt.Println("\n==========================\n")
	fmt.Printf("Soferul cu cea mai scumpa masina este: %s", getNumeSoferMasinaScumpa(listaDubla))
	fmt.Println("\n==========================\n")
}
