package main

import (
	"os"
	"strconv"
	"time"
)

const (
	wDayTitle = "Su Mo Tu We Th Fr Sa"
	wDaySeparator = "      \n"
	space = " "
	space2 = "  "
	space3 = "   "
	newline = "\n"
	newline2 = "\n\n"
)

var monthNumMap = map[string]int {
	"january":int(time.January), "jan":int(time.January),
	"february":int(time.February), "feb":int(time.February),
	"march":int(time.March), "mar":int(time.March),
	"april":int(time.April), "apr":int(time.April),
	"may":int(time.May),
	"june":int(time.June), "jun":int(time.June),
	"july":int(time.July), "jul":int(time.July),
	"august":int(time.August), "aug":int(time.August),
	"september":int(time.September), "sep":int(time.September),
	"october":int(time.October), "oct":int(time.October),
	"november":int(time.November), "nov":int(time.November),
	"december":int(time.December), "dec":int(time.December),
}

func usage() {
	os.Stderr.WriteString("usage: cal [[month] year]\n")
	os.Exit(1)
}

func printCal(from, to time.Month, year, monthsPerRow int) {
	monthName := func(m time.Month) string {
		return m.String() + space + strconv.Itoa(year)
	}

	if int(to - from) > 0 {
		monthName = func(m time.Month) string {
			return m.String()
		}

		for i:=(len(wDayTitle)*monthsPerRow+3*(monthsPerRow-1)-
			len(strconv.Itoa(year)))/2; i>0; i-- {
			os.Stdout.WriteString(space)
		}
		os.Stdout.WriteString(strconv.Itoa(year)+newline2)
	}

	for months, i := make([]time.Time, monthsPerRow), 0;
	    from <= to;
	    from=time.Month(int(from)+len(months)) {

		if int(to) - int(from) + 1 < len(months) {
			months = months[0:int(to)-int(from)+1]
		}

		/* print month names */
		for i=0; i < len(months); i++{
			months[i] = time.Date(year,
					time.Month(int(from)+i),
					1,0,0,0,0,time.UTC)

			if i > 0 {
				os.Stdout.WriteString(space3)
			}

			name, j := monthName(months[i].Month()), 0
			for ; j<(len(wDayTitle)-len(name))/2; j++ {
				os.Stdout.WriteString(space)
			}
			os.Stdout.WriteString(name)

			for j+=len(name); j<len(wDayTitle); j++ {
				os.Stdout.WriteString(space)
			}
		}
		os.Stdout.WriteString(newline)

		/* print weekday names */
		for i=0; i < len(months); i++ {
			if i > 0 { os.Stdout.WriteString(space3) }
			os.Stdout.WriteString(wDayTitle)
		}
		os.Stdout.WriteString(newline)

		row: for i=0; ; i=(i+1)%len(months) {
			if i > 0 { os.Stdout.WriteString(space3) }

			for wday := time.Sunday; wday <= time.Saturday; wday++ {
				if wday > time.Sunday {
					os.Stdout.WriteString(space)
				}

				if wday < months[i].Weekday() ||
				   months[i].Month() > time.Month(int(from)+i)  {
					os.Stdout.WriteString(space2)
					continue
				}

				if months[i].Day() < 10 {
					os.Stdout.WriteString(space)
				}

				os.Stdout.WriteString(strconv.Itoa(months[i].Day()))
				months[i] = months[i].AddDate(0,0,1)

				if months[i].Month() >= time.Month(int(from)+len(months)) ||
				   months[i].Year() > year {
					break row
				}
			}

			if i == len(months)-1 { os.Stdout.WriteString(newline) }
		}

		os.Stdout.WriteString(newline2)
	}
}

func Cal(month int, year int) {
	var (
		hasMonth = month >= 1 && month <= 12
		hasYear = year >= 1 && year <= 9999
		mon = time.Now().Month()
	)

	if hasMonth { mon = time.Month(month) }
	if !hasYear { year = time.Now().Year() }

	if hasMonth || (!hasMonth && !hasYear) {
		printCal(mon,mon,year,1)
	} else {
		/* print months in 3-columns view */
		printCal(time.January,time.December,year,3)
	}
}

func main() {
	var m, y int
	var e error
	switch len(os.Args) {
		case 2, 3:
			y, e = strconv.Atoi(os.Args[len(os.Args)-1])
			if e != nil || y < 1 || y > 9999 {
				os.Stderr.WriteString("cal: bad argument\n")
				break
			}

			if len(os.Args) == 3 {
				var ok bool

				if m, ok = monthNumMap[os.Args[1]]; !ok {
					m, e = strconv.Atoi(os.Args[1])
				}

				if e != nil || m < 1 || m > 12 {
					break
				}
			}

			fallthrough
		case 1:
			Cal(m, y)
			os.Exit(0)
	}
	usage()
}
