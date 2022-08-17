package model

// PostItemModel 포스트 목록 Item 타입
type PostItemModel interface {
	PostListItem | PostDetailItem | EmptyType
}

// PostData 글쓰기용 DTO
// 글 수정용 DTO도 embed로 함께 쓴다.
// BlogName			블로그 명 (필수)
// Title 			포스트 제목 (필수)
// Content 			포스트 내용
// Visibility 		포스트 공개 여부 상태 (0: 비공개[default], 1: 보호, 3: 발행)
// Category 		포스트 카테고리 (0: 없음[default], 카테고리 ID: 미리 블로그에 생성해둔 ID, 이 카테고리는 카테고리 목록 읽기 API를 통해 확인 가능하다.)
// Published 		배포 시간 (시간[TIMESTAMP]: 현재 시간[default], POST 요청 시간보다 미래의 시간을 넣을 경우 포스트 예약으로 동작)
// Slogan 			문자 주소
// Tag 				태그 (',' 이 따옴표로 구분한다.)
// AcceptComment	댓글 허용 (0: 댓글 거부, 1: 댓글 허용[default])
// Password			보호글용 비밀 번호 (비워두면 공개)
type PostData struct {
	// BlogName 블로그 명 (필수)
	BlogName string `json:"blog_name"`

	// Title 포스트 제목 (필수)
	Title string `json:"title"`

	// Content 포스트 내용
	Content string `json:"content"`

	// Visibility 포스트 공개 여부 상태 (0: 비공개[default], 1: 보호, 3: 발행)
	Visibility string `json:"visibility"`

	// Category 포스트 카테고리 (0: 없음[default], 카테고리 ID: 미리 블로그에 생성해둔 ID, 이 카테고리는 카테고리 목록 읽기 API를 통해 확인 가능하다.)
	Category string `json:"category"`

	// Published 배포 시간 (시간[TIMESTAMP]: 현재 시간[default], POST 요청 시간보다 미래의 시간을 넣을 경우 포스트 예약으로 동작)
	Published string `json:"published"`

	// Slogan 문자 주소
	Slogan string `json:"slogan"`

	// Tag 태그 (',' 이 따옴표로 구분한다.)
	Tag string `json:"tag"`

	// AcceptComment 댓글 허용 (0: 댓글 거부, 1: 댓글 허용[default])
	AcceptComment string `json:"acceptComment"`

	// Password 보호글용 비밀 번호 (비워두면 공개)
	Password string `json:"password"`
}

// PostUpdateData 글 수정용 DTO
// PostId 	포스트 번호 (필수)
// PostData 글 쓰기용 DTO embed
type PostUpdateData struct {

	// PostId 포스트 번호 (필수)
	PostId string `json:"post_id"`

	// PostData 글 쓰기용 DTO embed
	PostData
}

// PostResult 포스트 목록 읽기용 리턴 DTO
// Status	HTTP 상태코드
// Item		포스트 목록 읽기 결과 데이터
type PostResult[M PostItemModel] struct {
	// Status  HTTP 상태코드
	Status string `json:"status"`

	// Item	포스트 목록 읽기 결과 데이터
	Item M `json:"item"`
}

// PostWriteResult 글쓰기, 글 수정용 리턴 DTO
// Status  	HTTP 상태코드
// PostId	포스트 ID
// Url		포스트 URL
type PostWriteResult struct {
	// Status  HTTP 상태코드
	Status string `json:"status"`

	// PostId	포스트 ID
	PostId string `json:"postId"`

	// Url	포스트 URL
	Url string `json:"url"`
}

// PostDetailItem 포스트 상세 보기용 Item 타입
// Url				티스토리 기본 URL
// SecondaryUrl		독립 도메인 URL
// Id				포스트 ID
// Title			포스트 제목
// Content			포스트 내용
// CategoryId		포스트 카테고리 ID
// PostUrl			포스트 대표 주소
// Visibility		포스트 공개 여부 상태 (0: 비공개[default], 1: 보호, 3: 발행)
// AcceptComment	댓글 허용 여부 (0: 댓글 거부, 1: 댓글 허용[default])
// AcceptTrackback	추적 허용 여부 (자세히 파악하지 못한 내용이다. 가이드에도 친절하게 나와있진 않다.)
// Tags	[] Tag 		태그 목록
// Comments			댓글 개수
// Trackbacks		추적 개수
// Date				발생시간 (시간[TIMESTAMP], milliseconds)
type PostDetailItem struct {
	// Url	티스토리 기본 URL
	Url string `json:"url"`

	// SecondaryUrl	독립 도메인 URL
	SecondaryUrl string `json:"secondaryUrl"`

	// Id	포스트 ID
	Id string `json:"id"`

	// Title	포스트 제목
	Title string `json:"title"`

	// Content	포스트 내용
	Content string `json:"content"`

	// CategoryId	포스트 카테고리 ID
	CategoryId string `json:"categoryId"`

	// PostUrl	포스트 대표 주소
	PostUrl string `json:"postUrl"`

	// Visibility	포스트 공개 여부 상태 (0: 비공개[default], 1: 보호, 3: 발행)
	Visibility string `json:"visibility"`

	// AcceptComment	댓글 허용 여부 (0: 댓글 거부, 1: 댓글 허용[default])
	AcceptComment string `json:"acceptComment"`

	// AcceptTrackback	추적 허용 여부 (자세히 파악하지 못한 내용이다. 가이드에도 친절하게 나와있진 않다.)
	AcceptTrackback string `json:"acceptTrackback"`

	// Tags	[] Tag 태그 목록
	Tags struct {
		Tag []string `json:"tag"`
	} `json:"tags"`

	// Comments	댓글 개수
	Comments string `json:"comments"`

	// Trackbacks	추적 개수
	Trackbacks string `json:"trackbacks"`

	// Date	발생시간 (시간[TIMESTAMP], milliseconds)
	Date string `json:"date"`
}

// PostListItem 포스트 목록 보기용 Item 타입
// Url							티스토리 기본 URL
// SecondaryUrl					독립 도메인 URL
// Page							현재 페이지
// Count						현재 페이지 글 개수
// TotalCount					전체 글 수
// Posts [] PostListItemData	글 리스트
type PostListItem struct {
	// Url	티스토리 기본 URL
	Url string `json:"url"`

	// SecondaryUrl	독립 도메인 URL
	SecondaryUrl string `json:"secondaryUrl"`

	// Page	현재 페이지
	Page string `json:"page"`

	// Count	현재 페이지 글 개수
	Count string `json:"count"`

	// TotalCount	전체 글 수
	TotalCount string `json:"totalCount"`

	// Posts [] PostListItemData	글 리스트
	Posts []PostListItemData `json:"posts"`
}

// PostListItemData 포스트 목록 보기용 Item 상세 데이터 타입
// Id			포스트 ID
// Title		포스트 제목
// PostUrl		포스트 URL
// Visibility	포스트 공개 여부 상태 (0: 비공개[default], 15: 보호, 20: 발행)
// CategoryId	카테고리 ID
// Comments		댓글 수
// Trackbacks	트랙백 수
// Date			발행 시간 (시간[TIMESTAMP], YYYY-mm-dd HH:MM:SS)
type PostListItemData struct {
	// Id	포스트 ID
	Id string `json:"id"`

	// Title	포스트 제목
	Title string `json:"title"`

	// PostUrl	포스트 URL
	PostUrl string `json:"postUrl"`

	// Visibility	포스트 공개 여부 상태 (0: 비공개[default], 15: 보호, 20: 발행)
	Visibility string `json:"visibility"`

	// CategoryId	카테고리 ID
	CategoryId string `json:"categoryId"`

	// Comments		댓글 수
	Comments string `json:"comments"`

	// Trackbacks	트랙백 수
	Trackbacks string `json:"trackbacks"`

	// Date		발행 시간 (시간[TIMESTAMP], YYYY-mm-dd HH:MM:SS)
	Date string `json:"date"`
}
