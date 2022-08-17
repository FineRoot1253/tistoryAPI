package tistoryAPI

import (
	"context"
	"encoding/json"
	"github.com/fineroot1253/tistoryAPI/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

/**
 	테스트 이후 직접 포스트 삭제할 것
	1. 서비스 생성
	2. 블로그 정보 요청
	3. 카테고리 정보 요청, 이후 결과 전역변수로 등록
	4. 블로그 글 쓰기 요청, 이후 결과 포스트 ID 전역 변수로 등록
	5. 블로그 댓글 쓰기 요청, 이후 결과 댓글 ID 전역 변수로 등록
	6. 블로그 글 목록 요청
	7. 블로그 최신 댓글 목록 요청
	8. 블로그 글 수정 요청
	9. 블로그 댓글 수정 요청
	10. 블로그 댓글 삭제 요청
	11. 블로그 댓글 목록 요청
	12. 블로그 글 첨부 파일 추가 요청
	13. 블로그 글 읽기 요청
*/

var userData model.UserData
var commentData model.CommentData
var commentUpdateData model.CommentUpdateData
var postData model.PostData
var postUpdateData model.PostUpdateData
var blogInfo struct {
	blogName string `json:"blog_name"`
}
var commonService Service

func init() {
	//todo 더미 데이터 초기화
	commonInitData, err := ioutil.ReadFile("testdata/test_common_data.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(commonInitData, &userData); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(commonInitData, &blogInfo); err != nil {
		log.Fatal(err)
	}

	commentDummy, err := ioutil.ReadFile("testdata/test_comment_data.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(commentDummy, &commentData); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(commentDummy, &commentUpdateData); err != nil {
		log.Fatal(err)
	}

	postDummy, err := ioutil.ReadFile("testdata/test_post_data.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(postDummy, &postData); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(postDummy, &postUpdateData); err != nil {
		log.Fatal(err)
	}

}

func TestService(t *testing.T) {

	type args struct {
		ctx      context.Context
		userData model.UserData
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "서비스 생성 테스트:[success]",
			args: args{
				ctx:      context.Background(),
				userData: userData,
			},
			wantErr: false,
		},
		{
			name: "서비스 생성 테스트:[failure]",
			args: args{
				ctx:      context.Background(),
				userData: model.UserData{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		serv, err := NewService(tt.args.ctx, tt.args.userData)
		if err != nil {
			if tt.wantErr {
				log.Println(err)
			} else {
				log.Panicln(err)
			}
		} else {
			token := serv.GetACToken()
			assert.NotEmpty(t, token)
			log.Println("NewService complete: ", token)
			commonService = serv
		}
	}
}

func Test_service_GetBlogInfo(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "개인 블로그 정보 목록 읽기 테스트:[success]",
			wantErr: false,
		},
	}
	var serv Service
	if err := getService(serv); err != nil {
		log.Panicln(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			info, err := serv.GetBlogInfo()
			if err != nil {
				if tt.wantErr {
					log.Println(err)
				} else {
					log.Panicln(err)
				}
			} else {
				assert.NotEmpty(t, info)
				log.Println("GetBlogInfo Complete: ", info)
			}
		})
	}
}

func Test_service_GetCategoryList(t *testing.T) {
	type args struct {
		blogName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "테스트:[success]",
			args:    args{blogName: blogInfo.blogName},
			wantErr: false,
		},
		{
			name:    "테스트:[failure]",
			args:    args{blogName: ""},
			wantErr: true,
		},
	}
	var serv Service
	if err := getService(serv); err != nil {
		log.Panicln(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serv.GetCategoryList(tt.args.blogName)
			if err != nil {
				if tt.wantErr {
					log.Println(err)
				} else {
					log.Panicln(err)
				}
			} else {
				assert.NotEmpty(t, got)
				log.Println("GetCategoryList Complete: ", got)
			}
		})
	}
}

func Test_service_WritePost(t *testing.T) {
	type args struct {
		data model.PostData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "글쓰기 테스트:[success]",
			args: args{
				data: postData,
			},
			wantErr: false,
		},
		{
			name: "글쓰기 테스트:[failure] (필수 필드 포함 전부 생략)",
			args: args{
				data: model.PostData{},
			},
			wantErr: true,
		},
	}
	var serv Service
	if err := getService(serv); err != nil {
		log.Panicln(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serv.WritePost(tt.args.data)
			if err != nil {
				if tt.wantErr {
					log.Println(err)
				} else {
					log.Panicln(err)
				}
			} else {
				assert.NotEmpty(t, got)
				postUpdateData.PostId = got.PostId
				log.Println("WritePost Complete: ", got)
			}
		})
	}
}

func Test_service_WriteComment(t *testing.T) {
	type args struct {
		data model.CommentData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "댓글 쓰기 테스트:[success]",
			args:    args{data: commentData},
			wantErr: false,
		},
		{
			name:    "댓글 쓰기 테스트:[failure] (필수 필드 포함 전부 생략)",
			args:    args{data: model.CommentData{}},
			wantErr: true,
		},
	}
	var serv Service
	if err := getService(serv); err != nil {
		log.Panicln(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serv.WriteComment(tt.args.data)
			if err != nil {
				if tt.wantErr {
					log.Println(err)
				} else {
					log.Panicln(err)
				}
			} else {
				assert.NotEmpty(t, got)
				log.Println("WriteComment Complete: ", got)
			}
		})
	}
}

func Test_service_GetPostList(t *testing.T) {
	type args struct {
		blogName   string
		pageNumber int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "글 목록 읽기 테스트:[success]",
			args: args{
				blogName:   blogInfo.blogName,
				pageNumber: 1,
			},
			wantErr: false,
		},
		{
			name: "글 목록 읽기 테스트:[failure] (페이지 넘버 디폴트값 오류)",
			args: args{
				blogName:   blogInfo.blogName,
				pageNumber: 0,
			},
			wantErr: true,
		},
		{
			name: "글 목록 읽기 테스트:[failure] (필수값인 블로그 이름 생략)",
			args: args{
				blogName:   "",
				pageNumber: 1,
			},
			wantErr: true,
		},
		{
			name: "글 목록 읽기 테스트:[failure] (필수값인 블로그 이름 생략 + 페이지 넘버 디폴트값 오류)",
			args: args{
				blogName:   "",
				pageNumber: 0,
			},
			wantErr: true,
		},
	}
	var serv Service
	if err := getService(serv); err != nil {
		log.Panicln(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serv.GetPostList(tt.args.blogName, tt.args.pageNumber)
			if err != nil {
				if tt.wantErr {
					log.Println(err)
				} else {
					log.Panicln(err)
				}
			} else {
				assert.NotEmpty(t, got)
				log.Println("GetPostList Complete: ", got)
			}
		})
	}
}

func Test_service_GetNewestCommentList(t *testing.T) {
	type args struct {
		blogName   string
		pageNumber int
		count      int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "최신 댓글 목록 읽기 테스트:[success]",
			args: args{
				blogName:   blogInfo.blogName,
				pageNumber: 1,
				count:      5,
			},
			wantErr: false,
		},
		{
			name: "최신 댓글 목록 읽기 테스트:[failure] (페이지 넘버 디폴트값 오류)",
			args: args{
				blogName:   blogInfo.blogName,
				pageNumber: 0,
				count:      5,
			},
			wantErr: true,
		},
		{
			name: "최신 댓글 목록 읽기 테스트:[failure] (카운트 넘버 디폴트 값 오류)",
			args: args{
				blogName:   blogInfo.blogName,
				pageNumber: 1,
				count:      11,
			},
			wantErr: true,
		},
		{
			name: "최신 댓글 목록 읽기 테스트:[failure] (필수 값인 블로그 이름 생략)",
			args: args{
				blogName:   "",
				pageNumber: 1,
				count:      5,
			},
			wantErr: true,
		},
		{
			name: "최신 댓글 목록 읽기 테스트:[failure] (필수 값인 블로그 이름 생략 + 페이지 넘버 디폴트 값, 카운트 디폴트 값 오류)",
			args: args{
				blogName:   "",
				pageNumber: 0,
				count:      11,
			},
			wantErr: true,
		},
	}
	var serv Service
	if err := getService(serv); err != nil {
		log.Panicln(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serv.GetNewestCommentList(tt.args.blogName, tt.args.pageNumber, tt.args.count)
			if err != nil {
				if tt.wantErr {
					log.Println(err)
				} else {
					log.Panicln(err)
				}
			} else {
				assert.NotEmpty(t, got)
				log.Println("GetNewestCommentList Complete: ", got)
			}
		})
	}
}

func Test_service_UpdatePost(t *testing.T) {
	type args struct {
		data model.PostUpdateData
	}
	tests := []struct {
		name    string
		args    args
		want    model.PostWriteResult
		wantErr bool
	}{
		{
			name: "글 수정 테스트:[success]",
			args: args{
				data: postUpdateData,
			},
			wantErr: false,
		},
		{
			name: "글 수정 테스트:[failure] (필수 필드 포함 전부 생략)",
			args: args{
				data: model.PostUpdateData{},
			},
			wantErr: true,
		},
	}
	var serv Service
	if err := getService(serv); err != nil {
		log.Panicln(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serv.UpdatePost(tt.args.data)
			if err != nil {
				if tt.wantErr {
					log.Println(err)
				} else {
					log.Panicln(err)
				}
			} else {
				assert.NotEmpty(t, got)
				log.Println("UpdatePost Complete: ", got)
			}
		})
	}
}

func Test_service_UpdateComment(t *testing.T) {
	type args struct {
		data model.CommentUpdateData
	}
	tests := []struct {
		name    string
		args    args
		want    model.CommentWriteResult
		wantErr bool
	}{
		{
			name: "댓글 수정 테스트:[success]",
			args: args{
				data: commentUpdateData,
			},
			wantErr: false,
		},
		{
			name:    "댓글 수정 테스트:[failure] (필수 필드 포함 전부 생략)",
			args:    args{data: model.CommentUpdateData{}},
			wantErr: true,
		},
	}
	var serv Service
	if err := getService(serv); err != nil {
		log.Panicln(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serv.UpdateComment(tt.args.data)
			if err != nil {
				if tt.wantErr {
					log.Println(err)
				} else {
					log.Panicln(err)
				}
			} else {
				assert.NotEmpty(t, got)
				log.Println("GetCategoryList Complete: ", got)
			}
		})
	}
}

func Test_service_DeleteComment(t *testing.T) {
	type args struct {
		blogName  string
		postId    string
		commentId string
	}
	tests := []struct {
		name    string
		args    args
		want    model.CommentDeleteResult
		wantErr bool
	}{
		{
			name: "댓글 삭제 테스트:[success]",
			args: args{
				blogName:  blogInfo.blogName,
				postId:    commentUpdateData.PostId,
				commentId: commentUpdateData.CommentId,
			},
			wantErr: false,
		},
		{
			name: "댓글 삭제 테스트:[failure] (필수 필드 포함 전부 생략)",
			args: args{
				blogName:  "",
				postId:    "",
				commentId: "",
			},
			wantErr: true,
		},
	}
	var serv Service
	if err := getService(serv); err != nil {
		log.Panicln(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serv.DeleteComment(tt.args.blogName, tt.args.postId, tt.args.commentId)
			if err != nil {
				if tt.wantErr {
					log.Println(err)
				} else {
					log.Panicln(err)
				}
			} else {
				assert.NotEmpty(t, got)
				log.Println("GetCategoryList Complete: ", got)
			}
		})
	}
}

func Test_service_GetCommentList(t *testing.T) {
	type args struct {
		blogName string
		postId   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "글 댓글 목록 읽기 테스트:[success]",
			args: args{
				blogName: blogInfo.blogName,
				postId:   commentUpdateData.PostId,
			},
			wantErr: false,
		},
		{
			name: "글 댓글 목록 읽기 테스트:[failure] (필수 필드 포함 전부 생략)",
			args: args{
				blogName: "",
				postId:   "",
			},
			wantErr: true,
		},
	}
	var serv Service
	if err := getService(serv); err != nil {
		log.Panicln(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serv.GetCommentList(tt.args.blogName, tt.args.postId)
			if err != nil {
				if tt.wantErr {
					log.Println(err)
				} else {
					log.Panicln(err)
				}
			} else {
				assert.NotEmpty(t, got)
				log.Println("GetCategoryList Complete: ", got)
			}
		})
	}
}

func Test_service_AttachFiles(t *testing.T) {
	type args struct {
		blogName string
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    model.AttachResult
		wantErr bool
	}{
		{
			name: "테스트:[success]",
			args: args{
				blogName: blogInfo.blogName,
				filePath: "testdata/square-gopher.png",
			},
			wantErr: false,
		},
		{
			name: "테스트:[failure] (필수 필드 포함 전부 생략)",
			args: args{
				blogName: "",
				filePath: "",
			},
			wantErr: true,
		},
	}
	var serv Service
	if err := getService(serv); err != nil {
		log.Panicln(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serv.AttachFiles(tt.args.blogName, tt.args.filePath)
			if err != nil {
				if tt.wantErr {
					log.Println(err)
				} else {
					log.Panicln(err)
				}
			} else {
				assert.NotEmpty(t, got)
				log.Println("GetCategoryList Complete: ", got)
			}
		})
	}
}

func Test_service_GetPost(t *testing.T) {
	type args struct {
		blogName string
		postId   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "테스트:[success]",
			args: args{
				blogName: blogInfo.blogName,
				postId:   commentUpdateData.PostId,
			},
			wantErr: false,
		},
		{
			name: "테스트:[failure]",
			args: args{
				blogName: "",
				postId:   "",
			},
			wantErr: true,
		},
	}
	var serv Service
	if err := getService(serv); err != nil {
		log.Panicln(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serv.GetPost(tt.args.blogName, tt.args.postId)
			if err != nil {
				if tt.wantErr {
					log.Println(err)
				} else {
					log.Panicln(err)
				}
			} else {
				assert.NotEmpty(t, got)
				log.Println("GetCategoryList Complete: ", got)
			}
		})
	}
}

func getService(serv Service) error {
	if commonService != nil {
		serv = commonService
	} else {
		var err error
		serv, err = NewService(context.Background(), userData)
		if err != nil {
			return err
		}
	}
	return nil
}
