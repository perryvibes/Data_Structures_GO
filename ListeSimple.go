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

// Functia de afisare Masini
func afisareMasini(vector []Masina) {
	for _, value := range vector {
		fmt.Println(value)
	}
}

// adaugare la final
func adaugaMasinaInLista(cap **Nod, masinaNoua Masina) {
	var nou *Nod
	nou.masina = &masinaNoua
	nou.next = nil
	if (*cap) == nil {
		(*cap) = nou
	} else {
		var p *Nod = (*cap)
		for p.next != nil {
			p = p.next
		}
		p.next = nou
	}
}

func main() {

}
