package main

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/hirokisan/search-fetcher/model/list"
)

const (
	BaseUrl = "https://www.google.com/search?"
	Num     = "10"
)

var (
	app     = kingpin.New("search-fetcher", "A command-line search fetch application.")
	post    = kingpin.Command("run", "search and store data on file")
	count   = post.Flag("count", "count of files to create").Short('c').Default(Num).Int()
	opener  = post.Flag("opener", "opener to divide people to segment").Short('o').Required().String()
	keyword = post.Flag("keyword", "meta keyword represent complain, suggestion, problem, ...").Short('k').Required().String()
)

func main() {
	switch kingpin.Parse() {
	case "run":
		Run()
	}
}

func Run() error {
	q1 := *opener
	q2 := *keyword
	num := *count
	now := time.Now().Format("2006-01-02-15-04")
	query := url.Values{}
	query.Add("num", strconv.Itoa(num))
	query.Add("q", q1)
	query.Add("q", q2)
	dirName := "./data/" + q1 + "_" + q2 + "_" + now + "/"
	if err := os.MkdirAll(dirName, 0777); err != nil {
		return err
	}

	tUrl := BaseUrl + query.Encode()

	doc, err := goquery.NewDocument(tUrl)
	if err != nil {
		return err
	}

	lists := make(list.Lists, num, num)
	doc.Find("h3").Each(func(i int, s *goquery.Selection) {
		lists[i].Title = s.Text()
	})
	doc.Find(".r a").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		lists[i].Url = list.AdjustUrl(url)
	})
	for _, l := range lists {
		if m, _ := regexp.MatchString("youtube", l.Url); m {
			continue
		}
		if m, _ := regexp.MatchString("twitter", l.Url); m {
			continue
		}
		if m, _ := regexp.MatchString("facebook", l.Url); m {
			continue
		}
		if m, _ := regexp.MatchString("instagram", l.Url); m {
			continue
		}
		if m, _ := regexp.MatchString("naver", l.Url); m {
			continue
		}
		if m, _ := regexp.MatchString("search.yahoo.co.jp", l.Url); m {
			continue
		}
		if m, _ := regexp.MatchString("books.google.co.jp", l.Url); m {
			continue
		}
		fmt.Println(l.Title)
		file, err := os.Create(dirName + l.Title + ".txt")
		if err != nil {
			return err
		}
		io.WriteString(file, l.Url+"\n")
		dc, err := goquery.NewDocument(l.Url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		dc.Find("h2").Each(func(_ int, s *goquery.Selection) {
			text := s.Text()
			text = list.AdjustText(text)
			io.WriteString(file, text+"\n")
		})
		dc.Find("p").Each(func(_ int, s *goquery.Selection) {
			text := s.Text()
			text = list.AdjustText(text)
			io.WriteString(file, text+"\n")
		})
		// アメログ
		dc.Find(".articleText").Each(func(_ int, s *goquery.Selection) {
			text := s.Text()
			text = list.AdjustText(text)
			io.WriteString(file, text+"\n")
		})
		dc.Find(".comment-item .body").Each(func(_ int, s *goquery.Selection) {
			text := s.Text()
			text = list.AdjustText(text)
			io.WriteString(file, text+"\n")
		})
	}
	return nil
}
