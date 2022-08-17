package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestTyping(t *testing.T) {
	// given
	t2 := TistoryResult[PostResult[PostListItem], PostListItem, EmptyType]{}
	dummys := []struct {
		description string      // description of the testdata case
		data        string      // post data
		result      interface{} // result
	}{
		{
			description: "포스트 목록[success]",
			data: `
{
  "tistory": {
    "status": "200",
    "item": {
      "url": "http://oauth-testdata.tistory.com",
      "secondaryUrl": "",
      "page": "1",
      "count": "10",
      "totalCount": "181",
      "posts": [
        {
          "id": "201",
          "title": "테스트 입니다.",
          "postUrl": "http://oauth-testdata.tistory.com/201",
          "visibility": "0",
          "categoryId": "0",
          "comments": "0",
          "trackbacks": "0",
          "date": "2018-06-01 17:54:28"
        }
      ]
    }
  }
}`,
			result: t2,
		},
		{
			description: "포스트 읽기[success]",
			data: `{
  "tistory":{
    "status":"200",
    "item":{
      "url":"http://oauth.tistory.com",
      "secondaryUrl":"",
      "id":"1",
      "title":"티스토리 OAuth2.0 API 오픈!",
      "content":"Test 입니다.",
      "categoryId":"0",
      "postUrl":"http://oauth.tistory.com/1",
      "visibility":"0",
      "acceptComment":"1",
      "acceptTrackback":"1",
      "tags":{
        "tag":["open", "api"]
      },
      "comments":"0",
      "trackbacks":"0",
      "date":"1303352668"
    }
  }
}`,
			result: TistoryResult[PostResult[PostDetailItem], PostDetailItem, EmptyType]{},
		},
		{
			description: "포스트 작성&수정[success]",
			data: `
			{
  "tistory":{
    "status":"200",
    "postId":"74",
    "url":"http://sampleUrl.tistory.com/74"
  }
}
			`,
			result: TistoryResult[PostWriteResult, EmptyType, EmptyType]{},
		},
	}
	// when

	if err := json.Unmarshal([]byte(dummys[0].data), &t2); err != nil {
		assert.Error(t, err)
	}
	log.Println(t2)
	for _, dummy := range dummys {
		log.Print(dummy.data + " ::: ")
		var temp interface{}
		if err := json.Unmarshal([]byte(dummy.data), &temp); err != nil {
			assert.Error(t, err)
		}
		log.Println(temp)
	}
}
