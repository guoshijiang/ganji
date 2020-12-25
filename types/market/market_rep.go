package market

import "time"


type QsListRet struct {
	QsId       int64  `json:"qs_id"`
	QsAuthor   string `json:"qs_author"`
	QsTitle    string `json:"qs_title"`
	CreateTime time.Time `json:"create_time"`
}


type QsDetailRet struct {
	QsId       int64   `json:"qs_id"`
	QsAuthor   string  `json:"qs_author"`
	QsTitle    string  `json:"qs_title"`
	QsDetail   string  `json:"qs_detail"`
	CreateTime time.Time `json:"create_time"`
}