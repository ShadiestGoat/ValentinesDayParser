package main

import (
	"fmt"
	"math"
)

// Price you are selling at
const SELL_PRICE = 2
// One time cost like shipping
const CONSTFEE = 2.1
// Bundle size. If you are buying 1 flower at a time and not wholesale, then set this to be 1
const BUY_AT_A_TIME float64 = 20
// Price to buy a bundle
const BUY_PRICE = 28

func Stats(real bool) {
	revenue := 0.0

	redTot := 0
	yellowTot := 0
	pinkTot := 0
	totDeliveries := 0

	for _, person := range people {
		if real {
			if person.StillOwes() {
				continue
			}
		}

		for _, order := range person.Breakdown {
			switch order.Color {
			case RED:
				redTot++
			case PINK:
				pinkTot++
			case YELLOW:
				yellowTot++
			}

			if order.WeDeliver {
				totDeliveries++
			}
		}

		if real {
			revenue += person.Paid
		} else {
			revenue += float64(len(person.Breakdown) * SELL_PRICE)
		}
	}

	flowerTot := pinkTot + yellowTot + redTot
	thingsBought := (nearestAtX(float64(pinkTot), BUY_AT_A_TIME) + nearestAtX(float64(redTot), BUY_AT_A_TIME) + nearestAtX(float64(yellowTot), BUY_AT_A_TIME))
	cost := thingsBought*BUY_PRICE

	suffix := " (Estimated)"

	if real {
		suffix = " (Real)"
	}

	fmt.Print(sperator("┌", "┐"))
	fmt.Print(centerText("Valentines Day" + suffix))
	fmt.Print(centerText(""))
	fmt.Print(centerThings(NumToSpace(redTot, 6), "├ Red Orders", "Red Orders ┤"))
	fmt.Print(centerThings(NumToSpace(yellowTot, 6), "├ Yellow Orders", "Yellow Orders ┤", ))
	fmt.Print(centerThings(NumToSpace(pinkTot, 6), "├ Pink Orders", "Pink Orders ┤"))
	fmt.Print(centerThings(NumToSpace(flowerTot, 6), "├ Total Orders", "Total Orders ┤", ))
	fmt.Print(centerThings(NumToSpace(thingsBought, 6), "├ Bundles Bought", "Bundles Bought ┤"))
	logSeperator()
	fmt.Print(centerThings(NumToSpace(totDeliveries, 6), "├ Total Deliveries", "Total Deliveries ┤"))
	logSeperator()
	fmt.Print(centerThings(NumToSpace(cost, 6), "├ Cost", "Cost ┤"))
	fmt.Print(centerThings(NumToSpace(int(math.Floor(revenue)), 6), "├ Revenue", "Revenue ┤"))
	fmt.Print(centerThings(NumToSpace(int(math.Floor(revenue - float64(cost) - CONSTFEE)), 6), "├ Profit", "Profit ┤"))
	fmt.Print(sperator("└", "┘"))
}
func logSeperator() {
	fmt.Print(sperator("├", "┤"))
}

func BuyerProfile(name string) {
	authorInfo := people[name]
	price := 0
	redTot := 0
	pinkTot := 0
	yellowTot := 0
	
	outputStr := sperator("┌", "┐") 
	outputStr += centerText("Valentines Day")
	outputStr += centerText("Buyers Profile")
	outputStr += centerText("")
	outputStr += sperator("├", "┤")

	outputStr += centerText("Items")
	outputStr += centerText("")

	for _, order := range authorInfo.Breakdown {
		price += SELL_PRICE
		id := order.ID
		back := false
		for len(id) != 4 {
			if back {
				id += " "
			} else {
				id = " " + id
			}
			back = !back
		}

		switch order.Color {
		case RED:
			redTot++
			outputStr += centerThings(id, "├ Red Rose", "Red Rose ┤")
		case PINK:
			pinkTot++
			outputStr += centerThings(id, "├ Pink Rose", "Pink Rose ┤")
		case YELLOW:
			yellowTot++
			outputStr += centerThings(id, "├ Yellow Rose", "Yellow Rose ┤")
		}

	}
	outputStr += sperator("├", "┤")

	
	outputStr += centerThings(NumToSpace(redTot, 6), "├ Red Orders", "Red Orders ┤")
	outputStr += centerThings(NumToSpace(yellowTot, 6), "├ Yellow Orders", "Yellow Orders ┤", )
	outputStr += centerThings(NumToSpace(pinkTot, 6), "├ Pink Orders", "Pink Orders ┤")
	outputStr += sperator("├", "┤")
	outputStr += centerThings(NumToSpace(redTot + yellowTot + pinkTot, 6), "├ Total Orders", "Total Orders ┤", )

	outputStr += centerThings(NumToSpace(len(authorInfo.Breakdown)*SELL_PRICE, 6), "├ Total Cost", "Total Cost ┤")
	
	outputStr += sperator("├", "┤")

	outputStr += centerThings(NumToSpace(int(authorInfo.TotalOwed()), 6), "├ Owed To Us", "Owed To Us ┤")

	outputStr += sperator("└", "┘")
	fmt.Print(outputStr)
}

func RecieverProfile(name string) {
	info := RosesForPerson{}

	for _, person := range people {
		for _, order := range person.Breakdown {
			if !order.WeDeliver {
				continue
			}
			if order.GetRecepient() != name {
				continue
			}
			switch order.Color {
			case RED:
				info.Red++
			case YELLOW:
				info.Yellow++
			case PINK:
				info.Pink++
			}
		}
	}

	str := sperator("┌", "┐")
	str += centerText("For Person") + centerText("")

	if info.Red != 0 {
		str += centerThings(NumToSpace(info.Red, 6), "├ Red Roses", "Red Roses ┤")
	}
	if info.Yellow != 0 {
		str += centerThings(NumToSpace(info.Yellow, 6), "├ Yellow Roses", "Yellow Roses ┤")
	}
	if info.Pink != 0 {
		str += centerThings(NumToSpace(info.Pink, 6), "├ Pink Roses", "Pink Roses ┤")
	}

	str += sperator("├", "┤")

	str += centerThings(NumToSpace(info.Red + info.Pink + info.Yellow, 6), "├ Total Roses", "Total Roses ┤")

	str += sperator("└", "┘")

	fmt.Print(str)
}

// type DeliveryInfo struct {
// 	RedNum int
// 	PinkNum int
// 	YellowNum int
// }

// type RoomDeliveryInfo struct {
// 	DeliveryInfo
// 	Room string
// 	Orders map[string]DeliveryInfo // for people
// }

// func (info *DeliveryInfo) Add(order Order) {
// 	switch order.Color {
// 	case RED:
// 		info.RedNum++
// 	case PINK:
// 		info.PinkNum++
// 	case YELLOW:
// 		info.YellowNum++
// 	}
// }

// func (info *RoomDeliveryInfo) AddMajor(order Order) {
// 	info.Add(order)
// 	personInfo := info.Orders[order.GetRecepient()]
// 	personInfo.Add(order)
// 	info.Orders[order.GetRecepient()] = personInfo
// }