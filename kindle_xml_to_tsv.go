package kindlexmltotsv

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"time"
)

type Response struct {
	AddUpdateList AddUpdateList `xml:"add_update_list"`
}

type AddUpdateList struct {
	MetaData []MetaData `xml:"meta_data"`
}

type MetaData struct {
	ASIN            string   `xml:"ASIN"`
	Title           string   `xml:"title"`
	Authors         []string `xml:"authors>author"`
	Publishers      []string `xml:"publishers>publisher"`
	PublicationDate string   `xml:"publication_date"`
	PurchaseDate    string   `xml:"purchase_date"`
	TextbookType    string   `xml:"textbook_type"`
	CdeContenttype  string   `xml:"cde_contenttype"`
	ContentType     string   `xml:"content_type"`
}

func Convert(name string) (string, error) {
	buf, err := os.ReadFile(name)
	if err != nil {
		return "", fmt.Errorf("os.ReadFile: %w", err)
	}
	r := Response{}
	err = xml.Unmarshal(buf, &r)
	if err != nil {
		return "", fmt.Errorf("xml.Unmarshal: %w", err)
	}
	// fmt.Printf("%#v", r)
	tsv := "ASIN\tタイトル\t著者\t出版社\t出版日\t購入日\t開始日\t完了日\n"
	for _, v := range r.AddUpdateList.MetaData {
		author := strings.Join(v.Authors, ";")
		publisher := strings.Join(v.Publishers, ";")
		pubdt, err := time.Parse("2006-01-02T15:04:05Z0700", v.PublicationDate)
		if err != nil {
			pubdt = time.Time{}
			// return "", fmt.Errorf("time.Parse PublicationDate(%s): %w", v.ASIN, err)
		}
		pubdate := pubdt.Local().Format("2006-01-02")
		purdt, err := time.Parse("2006-01-02T15:04:05Z0700", v.PurchaseDate)
		if err != nil {
			purdt = time.Time{}
			// return "", fmt.Errorf("time.Parse PurchaseDate(%s): %w", v.ASIN, err)
		}
		purdate := purdt.Local().Format("2006-01-02")
		// line := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		// 	v.ASIN, v.Title, author, publisher, pubdate, purdate, v.TextbookType, v.CdeContenttype, v.ContentType)
		line := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t\t\n",
			v.ASIN, v.Title, author, publisher, pubdate, purdate)
		tsv = tsv + line
	}
	return tsv, nil
}
