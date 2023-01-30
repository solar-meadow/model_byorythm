package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "time"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("enter your born year example(2002): ")
    scanner.Scan()
    birthyear, err := strconv.Atoi(scanner.Text())
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print("enter your born month example( 03 ) it's March: ")
    scanner.Scan()
    birthmonth, err := strconv.Atoi(scanner.Text())
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print("enter your born month example( 30 ) it's 30 March: ")
    scanner.Scan()
    birthday, err := strconv.Atoi(scanner.Text())
    if err != nil {
        log.Fatal(err)
    }
      //Рф = «Физическое состояние», РЭ = «Эмоциональное состояние»,
//РИ = «Интеллектуальное состояние», РТ =  «Творческие возможности»,
    pi := 3.14159265358979323846264338327950288419716939937510582097494459
    day := 0.0
    dateTime := time.Date(birthyear, time.Month(birthmonth), birthday, 0, 0, 0, 0, time.UTC)
    result := []string{}
    result = append(result, ("|      Дата      |    РФ    |   РЭ    |   РИ   |   РТ   |прожитые дни|"))
    fmt.Sprintln("|     Дата     |   РФ   |   РЭ   |   РИ   |   РТ   |   X-число прожитых дней   |")
    for {
        rf := math.Sin(2.0 * pi * day / 22.0)
        re := math.Sin(2 * pi * day / 27)
        ru := math.Sin(2 * pi * day / 32)
        rt := math.Sin(2 * pi * day / 37)
        if day == 7611.0 {
            break
        }
        result = append(result, fmt.Sprintf("|   %s   |   %.1f   |   %.1f   |   %.1f   |   %.1f   |   %.0f   |\n", dateTime.Format("2006-01-02"), rf, re, ru, rt, day))

        fmt.Printf("|   %s   |   %.1f   |   %.1f   |   %.1f   |   %.1f   |   %.1f   |\n", dateTime.Format("2006-01-02"), rf, re, ru, rt, day)
        day += 1.0
        dateTime = dateTime.AddDate(0, 0, 1)
    }
    file, err := os.Create("model-biorythm.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    for _, line := range result {
        fmt.Fprintln(file, line)
    }
}
