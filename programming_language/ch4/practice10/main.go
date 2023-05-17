package main

import (
	"fmt"
	"go_practice/programming_language/ch4/github"
	"log"
	"os"
	"time"
)

func main() {
	// lastMonth := time.Now().AddDate(0, -1, 0).Format("2006-01-02")
	// qStrings := append(os.Args[1:], "created:>"+lastMonth)
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)
	dateMap := make(map[string][]string)
	for _, item := range result.Items {
		date := item.CreatedAt
		diffDayFromNow := time.Since(date).Hours() / 24
		output := fmt.Sprintf("#%-5d %9.9s %.55s", item.Number, item.User.Login, item.Title)
		if diffDayFromNow < 30 {
			dateMap["lastMonth"] = append(dateMap["lastMonth"], output)
		} else if diffDayFromNow >= 30 && diffDayFromNow < 365 {
			dateMap["lastYear"] = append(dateMap["lastYear"], output)
		} else {
			dateMap["moreLastYear"] = append(dateMap["moreLastYear"], output)
		}
	}
	fmt.Println("----------１ヶ月未満に作成されたイシュー----------")
	for _, output := range dateMap["lastMonth"] {
		fmt.Println(output)
	}
	fmt.Println("\n----------１年未満に作成されたイシュー----------")
	for _, output := range dateMap["lastYear"] {
		fmt.Println(output)
	}
	fmt.Println("\n----------１年以上前に作成されたイシュー----------")
	for _, output := range dateMap["moreLastYear"] {
		fmt.Println(output)
	}
}
