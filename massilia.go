package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var beginnings = []string{
	"Vé ",
	"Hey ",
	"Minot ",
	"Peuchère ",
	"Ma foi ",
	"Fada ",
	"Ho Gàrri ",
	"Oh tronche d'Àpi ",
	"Doumé ",
	"Parles meilleur ",
	"Fatche de… ",
	"Zou ",
	"Méfi ",
	"Tronche-plate ",
	"Stàssi ",
	"Arretes de marronner de longue ",
	"Bonne mère ",
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
	"elle a la scoumoune ",
	"tu me nifles!",
	"il a une figuane de gobi",
	"j'ai eu nibe",
	"il a pris un taquet ",
	"j’ai quillé le ballon",
	"je me suis gagué",
	"j'ai passé la pièce car c'était tout pègant",
}

var endings = []string{
	" avec ton straou.",
	" au vélodrome.",
	" à Endoume.",
	" au cabanon.",
	" dans le teston.",
	" du jaune.",
	" du pastaga",
	" dans le cabestron.",
	" ça sent l'aïoli.",
	" avec ta figure de poulpe.",
	" avec tes oursins dans les poches.",
	", c'est une une belle de cagade.",
	" une soupe d'esques et te jeter aux goudes. ",
	" dans la Gineste.",
	" comme ce pébron de papé.",
	" sur la Corniche.",
	" devant tous ses collègues.",
	", c’est une trompette.",
	", c'est le ouaille.",
	", c'est une vraie arapède.",
	", c’est une radasse.",
	", c'est une bordille",
	", c'est une cagole.",
	", c'est une bouche.",
}

func randomizeList(list []string) []string {
	shuffled := list
	rand.Seed(time.Now().UnixNano())

	// randomize an array
	rand.Shuffle(len(list), func(i, j int) {
		tmp := list[i]
		list[i] = list[j]
		list[j] = tmp
	})

	return shuffled
}

func buildSentence(e []string) string {
	return strings.Join(e, "")
}

func main() {
	amount := 15
	shuffledBeginnings := randomizeList(beginnings)
	shuffledExpressions := randomizeList(expressions)
	shuffledEndings := randomizeList(endings)

	for i := 0; i < amount; i++ {
		exprs := []string{shuffledBeginnings[i], shuffledExpressions[i], shuffledEndings[i]}
		fmt.Println(buildSentence(exprs))
	}

	fmt.Println(beginnings)
}
