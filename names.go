package trifle

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"sync"
)

var lastnames = []string{}
var lastnameLock sync.RWMutex

var males = []string{
	"Emmett",
	"Darius",
	"Mac",
	"Ty",
	"Norbert",
	"Micah",
	"Galen",
	"Dale",
	"Herbert",
	"Keenan",
	"Pete",
	"Dominic",
	"Edwardo",
	"Titus",
	"Angel",
	"Wendell",
	"Quinn",
	"Jefferson",
	"Carmine",
	"Johnson",
	"Renaldo",
	"Silas",
	"Thaddeus",
	"Minh",
	"Dwain",
	"Simon",
	"Jerald",
	"Deon",
	"Cliff",
	"Russel",
	"Bryan",
	"Lorenzo",
	"Darrel",
	"Rolando",
	"Dominique",
	"Don",
	"Jarvis",
	"Milo",
	"Blake",
	"Cole",
	"Adolph",
	"Isaias",
	"Denny",
	"Pasquale",
	"Alberto",
	"Josh",
	"Garth",
	"Lewis",
	"Winfred",
	"Ashley",
}

func MaleFirstName() string {
	return males[rand.Intn(len(males))]
}

var females = []string{
	"Vernia",
	"Charlesetta",
	"Nona",
	"Soo",
	"Yen",
	"Barbra",
	"Jeanna",
	"Charleen",
	"Mitzie",
	"Ela",
	"Loreta",
	"Vergie",
	"France",
	"Mariette",
	"Lucinda",
	"Ulrike",
	"Sharen",
	"Glenda",
	"Erlene",
	"Sharie",
	"Shanice",
	"Marvis",
	"Shanita",
	"Elly",
	"Jaleesa",
	"Shameka",
	"Vivan",
	"June",
	"Ha",
	"Milagros",
	"Lakia",
	"Casandra",
	"Emerita",
	"Therese",
	"Colene",
	"Jennell",
	"Traci",
	"Adah",
	"Ozie",
	"Alethea",
	"Angelina",
	"Gia",
	"Lilli",
	"Mitsuko",
	"Felecia",
	"Artie",
	"Danielle",
	"Elli",
	"Zoraida",
	"Shani",
}

func FemaleFirstName() string {
	return females[rand.Intn(len(females))]
}

func LastName() string {
	if len(lastnames) == 0 {
		lastnameLock.Lock()
		if len(lastnames) == 0 {
			fh, err := os.Open("./surname.txt")
			Panicer(err)
			sc := bufio.NewScanner(fh)
			for sc.Scan() {
				lastnames = append(lastnames, strings.TrimSpace(sc.Text()))
			}
		}
	}
	return lastnames[rand.Intn(len(lastnames))]
}

var nameSexGeneratror = TraditionalSex

func Name() string {
	sex := nameSexGeneratror()
	if sex == SexMale {
		return MaleFirstName() + " " + LastName()
	}
	return FemaleFirstName() + " " + LastName()
}
