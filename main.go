package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/xuri/excelize/v2"
)

const FILENAME = "data.json"

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

type RoseColor string

const (
	RED 	RoseColor = "Red"
	YELLOW 	RoseColor = "Yellow"
	PINK 	RoseColor = "Pink"
)

type idMap map[string]bool

type Order struct {
	Color RoseColor
	Recepient string
	Author string
	WeDeliver bool
	Msg string
	ID string
}

func (order Order) GetRecepient() string {
	split := strings.Split(order.Recepient, " ")
	name := strings.Join(split[:len(split)-1], " ")
	return name
}

func (order Order) GetTutor() string {
	split := strings.Split(order.Recepient, " ")
	return split[len(split)-1]
}

type Author struct {
	Name string
	Paid float64
	Ids idMap
	Breakdown []Order
}

type RosesForPerson struct {
	Red int
	Pink int
	Yellow int
}

func (person Author) TotalOwed() float64 {
	return float64(len(person.Breakdown)) * float64(SELL_PRICE) - person.Paid
}

func (person Author) StillOwes() bool {
	return (float64(len(person.Breakdown)) * float64(SELL_PRICE) - person.Paid) != 0 || len(person.Breakdown) == 0
}

var people = map[string]Author{}
var ids = idMap{}

func loadNewData(fileName string) {
	file, err := excelize.OpenFile(fileName)
	panicIfErr(err)
	defer file.Close()
	rows, err := file.GetRows("Sheet1")
	panicIfErr(err)
	
	for i, row := range rows {
		newRow := []string{}
		for _, item := range row {
			newRow = append(newRow, strings.TrimSpace(item))
		}
		row = newRow

		if i == 0 {
			continue
		}

		if _, ok := ids[row[0]]; ok {
			continue
		}

		info := Author{}
		info.Name = row[4]

		name := row[4]
	
		if _, ok := people[name]; ok {
			info = people[name]
		}

		if info.Ids == nil {
			info.Ids = idMap{}
		}

		order := Order{}
	
		order.Color = RoseColor(row[5])
		order.Msg = row[6]
		order.Author = row[7]
		order.ID = row[0]
	
		if row[9] != "Get it myself" {
			order.WeDeliver = true
			order.Recepient = row[10]
		}

		amount := NumParse(row[8])

		for i := 0; i < amount; i++ {
			info.Breakdown = append(info.Breakdown, order)
		}
		
		info.Ids[row[0]] = true

		people[name] = info
	}
}

func save() {
	jsonOrders := []Author{}
	for _, order := range people {
		jsonOrders = append(jsonOrders, order)
	}
	file, err := os.Create(FILENAME)
	panicIfErr(err)
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(jsonOrders)
	panicIfErr(err)
}

func laodSavePoint() {
	file, err := os.Open(FILENAME)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		panic(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	panicIfErr(err)
	rawData := []Author{}
	err = json.Unmarshal(b, &rawData)
	panicIfErr(err)

	for _, order := range rawData {
		people[order.Name] = order
		for id := range order.Ids {
			ids[id] = true
		}
	}
}

var nameReg = regexp.MustCompile(`^((([A-Z]('[A-Z])?([a-z]|ã|ê|í|ç)+?-?)+|de) ?)+?$`)
var yearReg = regexp.MustCompile(`\d\d?\.\d`)

func main() {
	laodSavePoint()
	calculateWidth()

	if len(os.Args) < 2 {
		panic("Needs a command!")
	}

	switch os.Args[1] {
	case "buyers":
		if len(os.Args) < 3 {
			panic("Needs a sub-command!")
		}
		switch os.Args[2] {
		case "names":
			names := ""
			for name, person := range people {
				if len(person.Breakdown) == 0 {
					continue
				}
				if len(os.Args) > 2 {
					if os.Args[3] == "-f" && !person.StillOwes() {
						continue
					}
				}

				names += name + "\n"
			}
			fmt.Println(names[:len(names)-1])
		case "profile":
			BuyerProfile(strings.Join(os.Args[3:], " "))
		default:
			panic("What are you trying to do? Available: `names`, `profile`")
		}
	case "stats":
		real := false
		if len(os.Args) > 2 {
			if os.Args[2] == "-r" {
				real = true
			}
		}

		Stats(real)
	case "source":
		if len(os.Args) > 2 {
			loadNewData(os.Args[2])
			save()
		}
		fmt.Println("Done!")
	case "export":
		svgGen()
	case "bugfix", "bug", "bugs":
		for _, person := range people {
			for _, order := range person.Breakdown {
				if !order.WeDeliver {
					continue
				}
				split := strings.Split(order.Recepient, " ")
				bad := false
				if !yearReg.MatchString(split[len(split)-1]) {
					bad = true
				}
				if !nameReg.MatchString(strings.Join(split[:len(split)-1], " ")) {
					bad = true
				}

				if bad {
					fmt.Println(order.Recepient)
				}
			}
		}
	case "recievers":
		if len(os.Args) < 3 {
			panic("Needs a sub-command!")
		}
		switch os.Args[2] {
		case "names":
			namesMap := map[string]bool{}

			for _, person := range people {
				if person.StillOwes() {
					continue
				}
				for _, order := range person.Breakdown {
					if !order.WeDeliver {
						continue
					}
					namesMap[order.GetRecepient()] = true
				}
			}

			names := ""
			for name := range namesMap {
				names += name + "\n"
			}
			fmt.Println(names[:len(names)-1])
		case "profile":
			RecieverProfile(strings.Join(os.Args[3:], " "))
		default:
			panic("What are you trying to do? Available: `names`, `profile`")
		}
	default:
		panic("That's not an acceptable command!\nAvailable: `buyers`, `stats`, `recievers`, `bugs`" )
	}
}
