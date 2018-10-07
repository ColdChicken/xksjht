package structs

type ListArticleFilter struct {
	// 文章的分类，如果不提供则不做分类限制
	Tags []string `json:"tags"`
	// 请求方的起始位置
	CurrentPos int64 `json:"currentPos"`
	// 请求的文章数目
	RequestCnt int64 `json:"requestCnt"`
	// 是否包含文章内容
	ContainContent int64 `json:"containContent"`
	// 专栏
	Catalog string `json:"catalog"`
}

type Article struct {
	Id            int64         `json:"id"`
	Title         string        `json:"title"`
	CreateTime    string        `json:"createTime"`
	EditTime      string        `json:"editTime"`
	Creater       string        `json:"creater"`
	Tags          string        `json:"tags"`
	OriginalTag   int64         `json:"originalTag"`
	Content       string        `json:"content"`
	Catalog       string        `json:"catalog"`
	RawContent    string        `json:"rawContent"`
	ParsedContent *ParserResult `json:"parsedContent"`
}

type UserInfo struct {
	Username string `json:"username"`
}

type CreateArticleRequest struct {
	Creater     string `json:"creater"`
	Tags        string `json:"tags"`
	OriginalTag int64  `json:"originalTag"`
	RawContent  string `json:"rawContent"`
	Catalog     string `json:"catalog"`
}

type UpdateArticleRequest struct {
	ArticleId   int64  `json:"articleId"`
	Tags        string `json:"tags"`
	OriginalTag int64  `json:"originalTag"`
	RawContent  string `json:"rawContent"`
}

type PicInfo struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	RelPath    string `json:"relPath"`
	CreateTime string `json:"createTime"`
	Location   string `json:"location"`
	Type       string `json:"type"`
}
