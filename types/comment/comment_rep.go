package comment


type CommentListRep struct {
	Id          int64  `json:"id"`
	GoodsId     int64  `json:"goods_id"`
	UserId      int64  `json:"user_id"`
	Title       string `json:"title"`
	Star        int8   `json:"star"`
	Content     string `json:"content"`
	ImgOne    string `json:"img_one_id"`
	ImgTwo    string `json:"img_two_id"`
	ImgThree  string `json:"img_three_id"`
}
