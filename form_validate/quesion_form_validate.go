package form_validate

type QuestionForm struct {
	Id        	int64    				`form:"id"`
	Author    	string   				`form:"author"`
	QsTitle   	string   				`form:"qs_title"`
	QsDetail  	string   				`form:"qs_detail"`
	IsCreate 	int    	 	 			`form:"_create"`
}

func (*QuestionForm) Messages() {
	//todo
}