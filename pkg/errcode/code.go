package errcode

var (
	MissingArticleId       = NewError(10010000, "缺少文章编号")
	ArticleTitleRequired   = NewError(10010001, "文章标题必填")
	ArticleContentRequired = NewError(10010002, "文章内容必填")
	GetArticleError        = NewError(10010003, "获取文章错误")
	UpdateArticleError     = NewError(10010004, "更新文章错误")
	CreateArticleError     = NewError(10010005, "创建文章错误")
	DeleteArticleError     = NewError(10010006, "删除文章错误")
)
