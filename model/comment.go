package model

type CommentItemModel interface {
	CommentNewestListItem | CommentListItem | EmptyType
}

// CommentData 댓글 쓰기용 DTO
// BlogName 블로그 명 (필수)
// PostId	포스트 ID (필수)
// ParentId	부모 댓글 Id (대댓글 인경우에 사용할 것)
// Content	댓글 내용
// Secret	비밀 댓글 여부(0: 공개[default], 1: 비밀)
type CommentData struct {
	// BlogName 블로그 명 (필수)
	BlogName string `json:"blog_name"`

	// PostId	포스트 ID (필수)
	PostId string `json:"post_id"`

	// ParentId	부모 댓글 Id (대댓글 인경우에 사용할 것)
	ParentId string `json:"parent_id"`

	// Content	댓글 내용
	Content string `json:"content"`

	// Secret	비밀 댓글 여부(0: 공개[default], 1: 비밀)
	Secret string `json:"secret"`
}

// CommentUpdateData 댓글 쓰기용 DTO
// CommentId 	댓글 ID (필수)
// CommentData	댓글 쓰기용 DTO embed
type CommentUpdateData struct {
	// CommentId 댓글 ID (필수)
	CommentId string `json:"comment_id"`

	// CommentData 댓글 쓰기용 DTO embed
	CommentData
}

// CommentResult 댓글 목록 읽기용 리턴 DTO
// Status	HTTP 상태코드
// Item		댓글 목록 읽기 결과 데이터
type CommentResult[M CommentItemModel] struct {
	// Status  HTTP 상태코드
	Status string `json:"status"`

	// Item	댓글 목록 읽기 결과 데이터
	Item M `json:"item"`
}

// CommentWriteResult 댓글 쓰기, 댓글 수정용 리턴 DTO
// Status  		HTTP 상태코드
// Result		결과 메시지
// CommentUrl	댓글 URL
type CommentWriteResult struct {
	// Status  HTTP 상태코드
	Status string `json:"status"`

	// Result	결과 메시지
	Result string `json:"result"`

	// CommentUrl	댓글 URL
	CommentUrl string `json:"commentUrl"`
}

// CommentDeleteResult 댓글 삭제용 리턴 DTO
// Status  HTTP 상태코드
type CommentDeleteResult struct {
	// Status  HTTP 상태코드
	Status string `json:"status"`
}

// CommentNewestListItem 포스트 상세 보기용 Item 타입
// Url			티스토리 기본 URL
// SecondaryUrl	독립 도메인 URL
// Comments 	CommentNewestDataList 최신 댓글 목록 데이터
type CommentNewestListItem struct {
	// Url			티스토리 기본 URL
	Url string `json:"url"`

	// SecondaryUrl	독립 도메인 URL
	SecondaryUrl string `json:"secondaryUrl"`

	// Comments 	CommentNewestDataList 최신 댓글 목록 데이터
	Comments CommentNewestDataList `json:"comments"`
}

// CommentListItem 포스트 상세 보기용 Item 타입
// Url			티스토리 기본 URL
// SecondaryUrl	독립 도메인 URL
// PostId		포스트 ID
// TotalCount	총 댓글 개수
// Comments 	CommentDataList 댓글 목록 데이터
type CommentListItem struct {
	// Url			티스토리 기본 URL
	Url string `json:"url"`

	// SecondaryUrl	독립 도메인 URL
	SecondaryUrl string `json:"secondaryUrl"`

	// PostId		포스트 ID
	PostId string `json:"postId"`

	// TotalCount	총 댓글 개수
	TotalCount string `json:"totalCount"`

	// Comments 	CommentDataList 댓글 목록 데이터
	Comments CommentDataList `json:"comments"`
}

// CommentNewestDataList
// Comment [] CommentNewestItemData 최신 댓글 목록 아이템 리스트
type CommentNewestDataList struct {
	// Comment [] CommentNewestItemData 최신 댓글 목록 아이템 리스트
	Comment []CommentNewestItemData
}

// CommentDataList
// Comment [] CommentListItemData 댓글 목록 아이템 리스트
type CommentDataList struct {
	// Comment [] CommentListItemData 댓글 목록 아이템 리스트
	Comment []CommentListItemData `json:"comment"`
}

// CommentNewestItemData 최신 댓글 목록 아이템 타입
// Id		댓글 ID
// Date		댓글 작성시간 (시간[TIMESTAMP], milliseconds)
// PostId	포스트 ID
// Name		작성자 이름
// Homepage	작성자 홈페이지 주소
// Comment	댓글 내용
// Open		댓글 공개여부 (Y: 공개, N: 비공개)
type CommentNewestItemData struct {
	// Id		댓글 ID
	Id string `json:"id"`

	// Date		댓글 작성시간 (시간[TIMESTAMP], milliseconds)
	Date string `json:"date"`

	// PostId	포스트 ID
	PostId string `json:"postId"`

	// Name		작성자 이름
	Name string `json:"name"`

	// Homepage	작성자 홈페이지 주소
	Homepage string `json:"homepage"`

	// Comment	댓글 내용
	Comment string `json:"comment"`

	// Open		댓글 공개여부 (Y: 공개, N: 비공개)
	Open string `json:"open"`
}

// CommentListItemData 댓글 목록 아이템 타입
// Id			댓글 ID
// Date			댓글 작성시간 ()
// Name			작성자 이름
// ParentId		대댓글 ID
// Homepage		작성자 홈페이지 주소
// Visibility	승인 여부 (0: 승인대기, 2: 승인)
// Comment		댓글 내용
// Open			댓글 공개여부 (Y: 공개, N: 비공개)
type CommentListItemData struct {
	// Id			댓글 ID
	Id string `json:"id"`

	// Date			댓글 작성시간 ()
	Date string `json:"date"`

	// Name			작성자 이름
	Name string `json:"name"`

	// ParentId		대댓글 ID
	ParentId string `json:"parentId"`

	// Homepage		작성자 홈페이지 주소
	Homepage string `json:"homepage"`

	// Visibility	승인 여부 (0: 승인대기, 2: 승인)
	Visibility string `json:"visibility"`

	// Comment		댓글 내용
	Comment string `json:"comment"`

	// Open			댓글 공개여부 (Y: 공개, N: 비공개)
	Open string `json:"open"`
}
