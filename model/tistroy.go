package model

type Tistory[PI PostItemModel, CI CommentItemModel] interface {
	BlogResult | PostResult[PI] | PostWriteResult | CategoryResult | CommentResult[CI] | CommentWriteResult | CommentDeleteResult | AttachResult
}

type TistoryResult[T Tistory[PI, CI], PI PostItemModel, CI CommentItemModel] struct {
	Tistory T `json:"tistory"`
}

type EmptyType struct {
}
