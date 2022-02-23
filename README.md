# Valentine's Day 

This is code written for my school's Valentine's Day organisations. It can be used for statistics, message printing, and conflict resolution.

## Requirements

First of all, this is adressed to a very specific situation. 

1. There are 3 colors of roses, Red, Pink, Yellow.
2. Every person who is ordering this has a classroom that is in this vague format: 12.1 (ie. number.number)
3. Every person who is getting ordered for also has the same classroom type (see above)
4. Each 'order' has the same color of rose, recepient and delivery options
    - This means that a person can order multiple roses with the same options and it be considered the same order
    - This means that if a person wants to send to multiple people, they have place multiple orders
5. Auto parsing excel can only be done with an excel with the following setup:
    1. Ignores first row for headings and the like
    2. The sheet with the results is named 'Sheet1'
    3. This is the order for the results, column by column:
        - ID. This is the order ID
        - Ignore this column
        - Ignore this column
        - Ignore this column
        - Name. The name of the person who ordered the rose
        - Color. The color of the rose. Only parsed if the color is "Red", "Pink" or "Yellow". Case sensetive.
        - Message. A string, where the orderer can write whatever message
        - Name on the message. The orderer can put whatever name they wish for here, and it is what the other person sees
        - Rose amount. This is preferably an actual number rather than "One" and the like, but the code tries to parse human numbers
        - Would you like us to deliver this? The only options are "Get it myself" or "Have us deliver it to them". 
        - Recepient. This has to be in the form of "Name MiddleName LastName AnyOtherName 12.1". All that is important is that there is a classroom corresponding to the format stated above (point 2) after a space at the very end

Other than that, to run the sh files you need fzf, and to get them to look decent you need to be on linux. This is not tested on wayland.
Also need have golang installed, obviously. For svg to png convertion (aaa.sh is generated), you need to have inkscape.

## Setup

1. git clone `https://github.com/ShadiestGoat/ValentinesDayParser.git`
2. go install
3. go build
4. source ./completion.bash

## How to use

You need to first have the data. This can be done through manually creating a data.json file, or through auto parsing an excel with an above mentioned setup. To do excel, you need to build this and then `./valentines source FILEPATH`. If you are adding more onto the existing data, don't worry! This doesn't overwrite already parsed data.

Then, you can do all that want:

- `stats`: Shows stats about the current data. Shows costs, profits, etc. To adjust price per bundle and bundle amount etc you have to edit the actual constants at the top of `stats.go`
    - Optional flag `-r` to filter down to only those who have paid
- `export`: Exports all the orders as messages to ./outputs/*. Also generates the file `aaa.sh` so that you can convert exported svgs into pngs using inkscape
- `bugs`:  Shows all the names that are not in compliance with the standard set in point 5, last column
- `recievers` and `buyers`: These 2 commands are not really meant for human use, it's easier to use the sh files with the same name. They have the same API though, so they are grouped together:
    - Sub-cmd `names`: Puts out a list of recepients, seperated by newlines. 
        - In receivers, filters down by wether or not we deliver their order (if we don't deliver, we ignore it), and wether their assosiated order is paid for. If it isn't paid for, it is ignored.
        - In buyers, you have the option to filter down by if they owe us to not. Use `-f` to filter out those who don't owe us!
    - Sub-cmd `profile`: The rest of the command line arguments are the name of the person needed (without classroom). Outputs the profile of the person. 

Run `sh receivers.sh` and `sh buyers.sh` for a more graphic view of the data. `buyers.sh` filters down people who still owe you, so should be used primarily for dept collection

## Known limitations

- Almost no way to manipulate data, such as editing an order, or adding a custom one, etc. 
- Slow png convertion (no goroutines & having to execute `aaa.sh`)
- Strict and bad excel parsing.
- Uniform SVG generation (ie. no dynamic color or anything like that)
- Incomplete bash completion
- JSON Database