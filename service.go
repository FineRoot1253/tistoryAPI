package tistoryAPI

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/fineroot1253/tistoryAPI/model"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

// Service TistoryAPI 인터페이스
// 초기화 조건 : Selenium 설치후 authorization code를 받아야한다.
// 블로그 명 : 블로그 URL의 'xxx.tistory.com' xxx 부분을 의미 한다.
type Service interface {

	// WritePost 글 작성하기
	// @Param model.PostData
	// return model.PostWriteResult, error
	WritePost(data model.PostData) (model.PostWriteResult, error)
	// WriteComment 댓글 작성하기
	// @Param model.CommentData
	// return model.CommentWriteResult, error
	WriteComment(data model.CommentData) (model.CommentWriteResult, error)
	// AttachFiles 파일 첨부하기
	// @Param string	// 블로그 명
	// return model.AttachResult, error
	AttachFiles(blogName string, filePath string) (model.AttachResult, error)

	// GetBlogInfo 블로그 정보 가져오기
	// return model.BlogResult, error
	GetBlogInfo() (model.BlogResult, error)

	// GetPostList 글 목록 가져오기
	// @Param string int	// 블로그 명, 글 목록 페이지 넘버
	// return model.PostResult[model.PostListItem], error
	GetPostList(blogName string, pageNumber int) (model.PostResult[model.PostListItem], error)
	// GetPost 글 상세 데이터 가져오기
	// @Param string	// 블로그 명
	// return model.PostResult[model.PostDetailItem], error
	GetPost(blogName, postId string) (model.PostResult[model.PostDetailItem], error)
	// GetNewestCommentList 최신 댓글 목록 가져오기
	// @Param string, int, int	// 블로그 명, 댓글 목록 페이지 넘버, 댓글 페이지당 댓글 수 (기본=10, 최대=10)
	// return model.CommentResult[model.CommentNewestListItem], error
	GetNewestCommentList(blogName string, pageNumber int, count int) (model.CommentResult[model.CommentNewestListItem], error)
	// GetCommentList 댓글 목록 가져오기
	// @Param model.CommentData
	// return model.CommentResult[model.CommentListItem], error
	GetCommentList(blogName, postId string) (model.CommentResult[model.CommentListItem], error)
	// GetCategoryList 카테고리 목록 가져오기
	// @Param string	//블로그 명
	// return model.CategoryResult, error
	GetCategoryList(blogName string) (model.CategoryResult, error)

	// UpdatePost 글 수정하기
	// @Param model.PostUpdateData
	// return model.PostWriteResult, error
	UpdatePost(data model.PostUpdateData) (model.PostWriteResult, error)
	// UpdateComment 댓글 수정하기
	// @Param model.CommentUpdateData
	// return model.CommentWriteResult, error
	UpdateComment(data model.CommentUpdateData) (model.CommentWriteResult, error)

	// DeleteComment 댓글 삭제하기
	// @Param uint64, uint64
	// return model.CommentDeleteResult, error
	DeleteComment(blogName string, postId string, commentId string) (model.CommentDeleteResult, error)

	// GetACToken 엑세스 토큰 확인
	// return string
	GetACToken() string
}

type service struct {
	// 각종 API에 사용되는 클라이언트
	client http.Client
	ctx    context.Context

	accessToken string
}

// NewService Tistory API 생성함수
// 생성시 http 통신을 통해 Access Token을 받고 세팅한다.
// 이 생성 과정중 에러 발생 가능성이 가장 높으므로 잘 테스트 해보고 사용 할 것
// @Params context.Context model.UserData	// 컨텍스트, Tistory Open API 유저 데이터
// return Service, error
func NewService(ctx context.Context, userData model.UserData) (Service, error) {

	getAccessTokenPath := TISTORY_OAUTH_ACCESSTOKEN_GET_PATH +
		"?client_id=" + userData.ClientId +
		"&client_secret=" + userData.SecretKey +
		"&redirect_uri=" + userData.RedirectUrl +
		"&code=" + userData.AuthorizationCode +
		"&grant_type=authorization_code"

	client := http.Client{}

	reqWithCtx, err := http.NewRequestWithContext(ctx, http.MethodGet, getAccessTokenPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(reqWithCtx)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &service{client, ctx, string(all)}, nil

}

func (s service) WritePost(data model.PostData) (model.PostWriteResult, error) {

	// SET Result Record and Input Record
	result := model.TistoryResult[model.PostWriteResult, model.EmptyType, model.EmptyType]{Tistory: model.PostWriteResult{}}

	sendUrl := TISTORY_API_URL + "/post/write" +
		"?access_token=" + s.accessToken +
		"&output=json" +
		"&blogName=" + data.BlogName +
		"&title=" + data.Title +
		"&content=" + data.Content +
		"&visibility=" + data.Visibility +
		"&category=" + data.Category +
		"&published=" + data.Published +
		"&slogan=" + data.Slogan +
		"&tag=" + data.Tag +
		"&acceptComment=" + data.AcceptComment +
		"&password=" + data.Password

	resp, err := s.client.Post(sendUrl, "application/json", http.NoBody)
	if err != nil {
		return result.Tistory, err
	}

	if err := compute[model.PostWriteResult, model.EmptyType, model.EmptyType](&result.Tistory, resp); err != nil {
		return result.Tistory, err
	}
	return result.Tistory, nil
}

func (s service) WriteComment(data model.CommentData) (model.CommentWriteResult, error) {
	// SET Result Record and Input Record
	result := model.TistoryResult[model.CommentWriteResult, model.EmptyType, model.EmptyType]{Tistory: model.CommentWriteResult{}}
	sendUrl := TISTORY_API_URL + "/post/write" +
		"?access_token=" + s.accessToken +
		"&output=json" +
		"&blogName=" + data.BlogName +
		"&postId=" + data.PostId +
		"&parentId=" + data.ParentId +
		"&content=" + data.Content +
		"&secret=" + data.Secret

	resp, err := s.client.Post(sendUrl, "application/json", http.NoBody)
	if err != nil {
		return result.Tistory, err
	}

	if err := compute[model.CommentWriteResult, model.EmptyType, model.EmptyType](&result.Tistory, resp); err != nil {
		return result.Tistory, err
	}
	return result.Tistory, nil
}

func (s service) AttachFiles(blogName string, filePath string) (model.AttachResult, error) {

	result := model.TistoryResult[model.AttachResult, model.EmptyType, model.EmptyType]{Tistory: model.AttachResult{}}

	sendUrl := TISTORY_API_URL + "/post/attach" +
		"?access_token=" + s.accessToken +
		"&blogName=" + blogName

	// file stream open
	openedFile, err := os.Open(filePath)
	if err != nil {
		return result.Tistory, nil
	}

	// multipart form 생성
	// writer에 body 버퍼 참조
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	multiPartFile, err := writer.CreateFormFile("uploadedfile", openedFile.Name())
	if err != nil {
		return result.Tistory, nil
	}
	if _, err := io.Copy(multiPartFile, openedFile); err != nil {
		return result.Tistory, nil
	}

	// file stream close
	if err := writer.Close(); err != nil {
		return result.Tistory, nil
	}

	resp, err := s.client.Post(sendUrl, writer.FormDataContentType(), body)
	if err != nil {
		return result.Tistory, err
	}

	if err := compute[model.AttachResult, model.EmptyType, model.EmptyType](&result.Tistory, resp); err != nil {
		return result.Tistory, err
	}
	return result.Tistory, nil
}

func (s service) GetBlogInfo() (model.BlogResult, error) {

	sendUrl := TISTORY_API_URL + "/blog/info" +
		"?access_token=" + s.accessToken +
		"&output=json"

	result := model.TistoryResult[model.BlogResult, model.EmptyType, model.EmptyType]{Tistory: model.BlogResult{}}

	resp, err := s.client.Get(sendUrl)
	if err != nil {
		return result.Tistory, err
	}

	if err := compute[model.BlogResult, model.EmptyType, model.EmptyType](&result.Tistory, resp); err != nil {
		return result.Tistory, err
	}
	return result.Tistory, nil

}

func (s service) GetPostList(blogName string, pageNumber int) (model.PostResult[model.PostListItem], error) {
	sendUrl := TISTORY_API_URL + "/post/list" +
		"?access_token=" + s.accessToken +
		"&output=json" +
		"&blogName=" + blogName +
		"&page=" + strconv.Itoa(pageNumber)
	result := model.TistoryResult[model.PostResult[model.PostListItem], model.PostListItem, model.EmptyType]{Tistory: model.PostResult[model.PostListItem]{}}
	resp, err := s.client.Get(sendUrl)
	if err != nil {
		return result.Tistory, err
	}

	if err := compute[model.PostResult[model.PostListItem], model.PostListItem, model.EmptyType](&result.Tistory, resp); err != nil {
		return result.Tistory, err
	}
	return result.Tistory, nil
}

func (s service) GetPost(blogName, postId string) (model.PostResult[model.PostDetailItem], error) {
	sendUrl := TISTORY_API_URL + "/post/read" +
		"?access_token=" + s.accessToken +
		"&output=json" +
		"&blogName=" + blogName +
		"&postId=" + postId
	result := model.TistoryResult[model.PostResult[model.PostDetailItem], model.PostDetailItem, model.EmptyType]{Tistory: model.PostResult[model.PostDetailItem]{}}
	resp, err := s.client.Get(sendUrl)
	if err != nil {
		return result.Tistory, err
	}

	if err := compute[model.PostResult[model.PostDetailItem], model.PostDetailItem, model.EmptyType](&result.Tistory, resp); err != nil {
		return result.Tistory, err
	}
	return result.Tistory, nil
}

func (s service) GetNewestCommentList(blogName string, pageNumber int, count int) (model.CommentResult[model.CommentNewestListItem], error) {
	sendUrl := TISTORY_API_URL + "/comment/newest" +
		"?access_token=" + s.accessToken +
		"&output=json" +
		"&blogName=" + blogName +
		"&page=" + fmt.Sprint(pageNumber) +
		"&count=" + fmt.Sprint(count)
	result := model.TistoryResult[model.CommentResult[model.CommentNewestListItem], model.EmptyType, model.CommentNewestListItem]{Tistory: model.CommentResult[model.CommentNewestListItem]{}}
	resp, err := s.client.Get(sendUrl)
	if err != nil {
		return result.Tistory, err
	}

	if err := compute[model.CommentResult[model.CommentNewestListItem], model.EmptyType, model.CommentNewestListItem](&result.Tistory, resp); err != nil {
		return result.Tistory, err
	}
	return result.Tistory, nil
}

func (s service) GetCommentList(blogName, postId string) (model.CommentResult[model.CommentListItem], error) {
	sendUrl := TISTORY_API_URL + "/comment/list" +
		"?access_token=" + s.accessToken +
		"&output=json" +
		"&blogName=" + blogName +
		"&postId=" + postId
	result := model.TistoryResult[model.CommentResult[model.CommentListItem], model.EmptyType, model.CommentListItem]{Tistory: model.CommentResult[model.CommentListItem]{}}
	resp, err := s.client.Get(sendUrl)
	if err != nil {
		return result.Tistory, err
	}

	if err := compute[model.CommentResult[model.CommentListItem], model.EmptyType, model.CommentListItem](&result.Tistory, resp); err != nil {
		return result.Tistory, err
	}
	return result.Tistory, nil
}

func (s service) GetCategoryList(blogName string) (model.CategoryResult, error) {
	sendUrl := TISTORY_API_URL + "/category/list" +
		"?access_token=" + s.accessToken +
		"&output=json" +
		"&blogName=" + blogName
	result := model.TistoryResult[model.CategoryResult, model.EmptyType, model.EmptyType]{Tistory: model.CategoryResult{}}
	resp, err := s.client.Get(sendUrl)
	if err != nil {
		return result.Tistory, err
	}

	if err := compute[model.CategoryResult, model.EmptyType, model.EmptyType](&result.Tistory, resp); err != nil {
		return result.Tistory, err
	}
	return result.Tistory, nil
}

func (s service) UpdatePost(data model.PostUpdateData) (model.PostWriteResult, error) {
	// SET Result Record and Input Record
	result := model.TistoryResult[model.PostWriteResult, model.EmptyType, model.EmptyType]{Tistory: model.PostWriteResult{}}
	sendUrl := TISTORY_API_URL + "/post/modify" +
		"?access_token=" + s.accessToken +
		"&output=json" +
		"&blogName=" + data.BlogName +
		"&postId=" + data.PostId +
		"&title=" + data.Title +
		"&content=" + data.Content +
		"&visibility=" + data.Visibility +
		"&category=" + data.Category +
		"&published=" + data.Published +
		"&slogan=" + data.Slogan +
		"&tag=" + data.Tag +
		"&acceptComment=" + data.AcceptComment +
		"&password=" + data.Password
	resp, err := s.client.Post(sendUrl, "application/json", http.NoBody)
	if err != nil {
		return result.Tistory, err
	}

	if err := compute[model.PostWriteResult, model.EmptyType, model.EmptyType](&result.Tistory, resp); err != nil {
		return result.Tistory, err
	}
	return result.Tistory, nil
}

func (s service) UpdateComment(data model.CommentUpdateData) (model.CommentWriteResult, error) {
	// SET Result Record and Input Record
	result := model.TistoryResult[model.CommentWriteResult, model.EmptyType, model.EmptyType]{Tistory: model.CommentWriteResult{}}
	sendUrl := TISTORY_API_URL + "/comment/modify" +
		"?access_token=" + s.accessToken +
		"&output=json" +
		"&blogName=" + data.BlogName +
		"&postId=" + data.PostId +
		"&parentId=" + data.ParentId +
		"&commentId=" + data.CommentId +
		"&content=" + data.Content +
		"&secret=" + data.Secret
	resp, err := s.client.Post(sendUrl, "application/json", http.NoBody)
	if err != nil {
		return result.Tistory, err
	}

	if err := compute[model.CommentWriteResult, model.EmptyType, model.EmptyType](&result.Tistory, resp); err != nil {
		return result.Tistory, err
	}

	return result.Tistory, nil
}

func (s service) DeleteComment(blogName string, postId string, commentId string) (model.CommentDeleteResult, error) {
	// SET Result Record and Input Record
	result := model.TistoryResult[model.CommentDeleteResult, model.EmptyType, model.EmptyType]{Tistory: model.CommentDeleteResult{}}
	sendUrl := TISTORY_API_URL + "/comment/delete" +
		"?access_token=" + s.accessToken +
		"&output=json" +
		"&blogName=" + blogName +
		"&postId=" + postId +
		"&commentId=" + commentId

	resp, err := s.client.Post(sendUrl, "application/json", http.NoBody)
	if err != nil {
		return result.Tistory, err
	}

	if err := compute[model.CommentDeleteResult, model.EmptyType, model.EmptyType](&result.Tistory, resp); err != nil {
		return result.Tistory, err
	}
	return result.Tistory, nil
}

func compute[T model.Tistory[PI, CI], PI model.PostItemModel, CI model.CommentItemModel](result *T, resp *http.Response) error {

	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(all, result); err != nil {
		return err
	}
	return nil
}

func (s service) GetACToken() string {
	return s.accessToken
}
