package weread

import (
	"testing"
)

func TestParseBooks(t *testing.T) {
	body := []byte(`{
		"synckey": 1705804212,
		"totalBookCount": 171,
		"noBookReviewCount": 0,
		"books": [ {
			"bookId": "34630808",
			"book": {
				"bookId": "34630808",
				"title": "写作是门手艺",
				"author": "刘军强",
				"cover": "https://cdn.weread.qq.com/weread/cover/86/YueWen_34630808/s_YueWen_34630808.jpg",
				"version": 2145347742,
				"format": "epub",
				"type": 0,
				"price": 20.8,
				"originalPrice": 0,
				"soldout": 0,
				"bookStatus": 1,
				"payType": 1048577,
				"centPrice": 2080,
				"finished": 1,
				"maxFreeChapter": 12,
				"free": 0,
				"mcardDiscount": 0,
				"ispub": 1,
				"extra_type": 5,
				"cpid": -3805127,
				"publishTime": "2020-07-01 00:00:00",
				"centPrice": 2080,
				"categories": [
					{
						"categoryId": 1400000,
						"subCategoryId": 1400003,
						"categoryType": 0,
						"title": "教育学习-教育"
					}
				],
				"hasLecture": 0,
				"lastChapterIdx": 55,
				"paperBook": {
					"skuId": "12944320"
				},
				"maxFreeInfo": {
					"maxFreeChapterIdx": 12,
					"maxFreeChapterUid": 67,
					"maxFreeChapterRatio": 7
				},
				"copyrightChapterUids": [
					57
				],
				"hasKeyPoint": true,
				"blockSaveImg": 0,
				"language": "zh",
				"hideUpdateTime": false,
				"isEPUBComics": 0,
				"webBookControl": 0
			},
			"reviewCount": 2,
			"reviewLikeCount": 1,
			"reviewCommentCount": 0,
			"noteCount": 0,
			"bookmarkCount": 0,
			"sort": 1685853159
		}]
	}`)
	var expectedTitle = "写作是门手艺"

	books, err := ParseBooks(body)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	actualBookTitle := books[0].BookInfo.Title
	if actualBookTitle != expectedTitle {
		t.Errorf("Expected %v, but got %v", expectedTitle, actualBookTitle)
	}

}
