package weread

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Synckey           int    `json:"synckey"`
	TotalBookCount    int    `json:"totalBookCount"`
	NoBookReviewCount int    `json:"noBookReviewCount"`
	Books             []Book `json:"books"`
}

type Book struct {
	BookID             string   `json:"bookId"`
	BookInfo           BookInfo `json:"book"`
	ReviewCount        int      `json:"reviewCount"`
	ReviewLikeCount    int      `json:"reviewLikeCount"`
	ReviewCommentCount int      `json:"reviewCommentCount"`
	NoteCount          int      `json:"noteCount"`
	BookmarkCount      int      `json:"bookmarkCount"`
	Sort               int      `json:"sort"`
}

type BookInfo struct {
	BookID               string      `json:"bookId"`
	Title                string      `json:"title"`
	Author               string      `json:"author"`
	Cover                string      `json:"cover"`
	Version              int         `json:"version"`
	Format               string      `json:"format"`
	Type                 int         `json:"type"`
	Price                float64     `json:"price"`
	OriginalPrice        float64     `json:"originalPrice"`
	Soldout              int         `json:"soldout"`
	BookStatus           int         `json:"bookStatus"`
	PayType              int         `json:"payType"`
	CentPrice            int         `json:"centPrice"`
	Finished             int         `json:"finished"`
	MaxFreeChapter       int         `json:"maxFreeChapter"`
	Free                 int         `json:"free"`
	McardDiscount        int         `json:"mcardDiscount"`
	Ispub                int         `json:"ispub"`
	Cpid                 int         `json:"cpid"`
	PublishTime          string      `json:"publishTime"`
	HasLecture           int         `json:"hasLecture"`
	LastChapterIdx       int         `json:"lastChapterIdx"`
	PaperBook            PaperBook   `json:"paperBook"`
	MaxFreeInfo          MaxFreeInfo `json:"maxFreeInfo"`
	OtherType            []OtherType `json:"otherType"`
	CopyrightChapterUids []int       `json:"copyrightChapterUids"`
	HasKeyPoint          bool        `json:"hasKeyPoint"`
	LimitShareChat       int         `json:"limitShareChat"`
	BlockSaveImg         int         `json:"blockSaveImg"`
	Language             string      `json:"language"`
	HideUpdateTime       bool        `json:"hideUpdateTime"`
	IsEPUBComics         int         `json:"isEPUBComics"`
	WebBookControl       int         `json:"webBookControl"`
}

type PaperBook struct {
	SkuId string `json:"skuId"`
}

type MaxFreeInfo struct {
	MaxFreeChapterIdx   int `json:"maxFreeChapterIdx"`
	MaxFreeChapterUid   int `json:"maxFreeChapterUid"`
	MaxFreeChapterRatio int `json:"maxFreeChapterRatio"`
}

type OtherType struct {
	Type            string `json:"type"`
	ShowType        bool   `json:"showType"`
	TranslateStatus string `json:"translateStatus"`
	TranslateDone   bool   `json:"translateDone"`
}

func ParseBooks(body []byte) ([]Book, error) {
	var webooks Response
	err := json.Unmarshal(body, &webooks)
	if err != nil {
		fmt.Println("Error while parsing books:", err)
		return nil, err
	}

	return webooks.Books, nil
}
