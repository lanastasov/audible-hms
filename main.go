package main

import (
    "fmt"
    "log"
    "os"
    "bufio"
    "strings"
    "strconv"
    "github.com/urfave/cli/v2"
    "github.com/lanastasov/audible-hms/cmd/audiblejson"
)

func main() {
    app := &cli.App{
        Name:  "audible-time-duration",
        Usage: "calculate hours minutes and seconds from a txt file, gets the time from the last 5 characters on each line",
        Action: func(cCtx *cli.Context) error {
		fmt.Printf("hello %q", cCtx.Args().Get(0))
		
		audiblejson.Chapters()

		file, err := os.Open("./audible-chapters-duration.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		var minutes = 0
		var seconds = 0
		var min = 0
		var sec = 0
		for scanner.Scan() {
			var line = scanner.Text()
			s := strings.Split(line[len(line)-5:],":")[1]
			m := strings.Split(line[len(line)-5:],":")[0]
			minutes, err = strconv.Atoi(m)
			seconds, err = strconv.Atoi(s)
			if err != nil {
				fmt.Println("err")
			}
			min += minutes
			sec += seconds
		}

		totalHours := min / 60
		totalMin   := min % 60

		totalmin   := sec / 60
		totalsec   := sec % 60
		totalhour  := totalmin / 60
		totalmin   = totalmin % 60
		
		totalHours += totalhour
		totalMin   += totalmin

		fmt.Println(min, " : " , sec)
		fmt.Println(totalHours, "h <:> ", totalMin, "m <:> ", totalsec, "sec")

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
