package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"regexp"
)

var heightReg = regexp.MustCompile(`{{%HEIGHT}}`)
var textReg = regexp.MustCompile(`{{%MSGS}}`)
var numReg = regexp.MustCompile(`\d+`)

func svgGen() {
	template, err := ioutil.ReadFile("template.svg")
	panicIfErr(err)

	for _, person := range people {
		for _, order := range person.Breakdown {
			file, err := os.Create("outputs/" + order.ID + ".svg")
			if err != nil {
				panic(err)
			}
			defer file.Close()
			from := order.Author
			if len(order.Author) == 0 {
				from = "Anonymous"
			}

			str := ""
			if order.WeDeliver {
				str += `<text y="30.194239"><tspan>To:</tspan><tspan font-size="12px">   ` + order.GetRecepient() + "</tspan></text>\n"
				str += `<text y="42.807598"><tspan>From:</tspan><tspan font-size="`
				if len(from) < 30 {
					str += "12"
				} else if len(from) < 45 {
					str += "10"
				} else if len(from) < 60 {
					str += "8.5"
				} else {
					str += "7"
				}
				str += `px">   ` + from + "</tspan></text>\n"
			} else {
				if len(order.Msg) == 0 {
					continue
				}
			}


			y := "55"
			if !order.WeDeliver {
				y = "45"
			}

			str += `<text y="` + y + `"><tspan>Message:</tspan><tspan font-size="`

			msgLines := strings.Split(order.Msg, "\n")

			avgLen := 0.0

			for _, s := range msgLines {
				avgLen += float64(len(s))
			}

			avgLen = avgLen/float64(len(msgLines))

			if avgLen < 30 {
				str += "12"
			} else if avgLen < 45 {
				str += "10"
			} else if avgLen < 60 {
				str += "8.5"
			} else {
				str += "8"
			}

			str += `px">   `

			for i, msg := range msgLines {
				if i != 0 {
					str += `<tspan dy="1.25em" x="45">`
				} else {
					str += `<tspan>`
				}
				str += msg
				str += "</tspan>"
			}

			str += "</tspan></text>"
			height := 30 + 9*len(msgLines)
			fileContent := heightReg.ReplaceAll(template, []byte(fmt.Sprint(height)))
			fileContent = textReg.ReplaceAll(fileContent, []byte(str))
			file.Write(fileContent)
		}
	}

	files, err := ioutil.ReadDir("outputs")
	panicIfErr(err)
	strB := ""
	for _, file := range files {
		name := "outputs/" + file.Name()
		content, err := ioutil.ReadFile(name)
		panicIfErr(err)
		if len(content) == 0 {
			continue
		}

		heightStart := string(content)[27:]
		heightStr := numReg.FindString(heightStart)
		height, err := strconv.ParseInt(heightStr, 10, 64)
		fmt.Println(height)
		panicIfErr(err)
		strB += fmt.Sprintf("inkscape -w 594 -h %v %v -o pngs/%v.png\n", height*2, name, file.Name()[:len(file.Name())-4])
		// cmd := exec.Command("/bin/inkscape", "-w 594", "-h " + fmt.Sprint(height*2), name, "-o " + "pngs/" + file.Name()[:len(file.Name())-4] + ".png")
		// cmd.Run()
		// panic(fmt.Sprintf("%#v", cmd))
		// err = cmd.Run()
		// panicIfErr(err)
	}

	fs, err := os.Create("aaa.sh")
	panicIfErr(err)
	fs.WriteString(strB)
}