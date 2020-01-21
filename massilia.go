package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var beginnings = []string{
	"Vé,",
	"Hey,",
	"Minot,",
	"Peuchère,",
	"Ma foi,",
	"Fada,",
	"Ho Gàrri,",
	"Oh tronche d'Àpi,",
	"Doumé,",
	"Parles meilleur,",
	"Fatche de…,",
	"Zou,",
	"Méfi,",
	"Tronche-plate,",
	"Stàssi,",
	"Arretes de marronner de longue,",
	"Bonne mère,",
}

var expressions = []string{
	"tu as l’air tchouche",
	"tu en as perdu la tchatche",
	"tu aimes mettre le ouaï",
	"tu es un mastre",
	"t'engatse-pas",
	"tu t’es fais chourer ton jaune",
	"t'es tout blanquinas",
	"arrete de faire la bèbe",
	"je t’escagass",
	"on va manger des panisses",
	"il est un peu calu",
	"Je crains dégun",
	"tu boulègues",
	"tu vas t’estramasser",
	"tu me gaves ",
	"elle a la scoumoune",
	"tu me nifles!",
	"il a une figuane de gobi",
	"j'ai eu nibe",
	"il a pris un taquet",
	"j’ai quillé le ballon",
	"je me suis gagué",
	"j'ai passé la pièce car c'était tout pègant",
}

var endings = []string{
	"avec ton straou.",
	"au vélodrome.",
	"à Endoume.",
	"au cabanon.",
	"dans le teston.",
	"du jaune.",
	"du pastaga.",
	"dans le cabestron.",
	"ça sent l'aïoli.",
	"avec ta figure de poulpe.",
	"avec tes oursins dans les poches.",
	", c'est une belle de cagade.",
	"une soupe d'esques et te jeter aux goudes. ",
	"dans la Gineste.",
	"comme ce pébron de papé.",
	"sur la Corniche.",
	"devant tous ses collègues.",
	", c’est une trompette.",
	", c'est le ouaille.",
	", c'est une vraie arapède.",
	", c’est une radasse.",
	", c'est une bordille",
	", c'est une cagole.",
	", c'est une bouche.",
}

func generate(count int) string {
	rand.Seed(time.Now().UnixNano())

	beginningsIdx := rand.Perm(len(beginnings))
	expressionsIdx := rand.Perm(len(expressions))
	endingsIdx := rand.Perm(len(endings))

	var p strings.Builder

	for i := 0; i < count; i++ {
		p.WriteString(beginnings[beginningsIdx[i%len(beginningsIdx)]])
		p.WriteRune(' ')
		p.WriteString(expressions[expressionsIdx[i%len(expressionsIdx)]])

		if endings[endingsIdx[i]][0:1] != "," {
			p.WriteRune(' ')
		}

		p.WriteString(endings[endingsIdx[i]])

		if i != count-1 {
			p.WriteRune(' ')
		}
	}

	return p.String()
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: generating ipsum...")
	var p = r.URL.Query()

	length := p.Get("length")
	generated := generateParagraphs(length)
}

func generateParagraphs(count int) []string {
	var result []string

	for i := 0; i < count; i++ {
		result = append(result, generate(5))
	}

	return result
}

func main() {
	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
