package client

import (
	"github.com/fineroot1253/tistoryAPI/model"
	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
	"log"
	"os"
	"testing"
	"time"
)

func Test_service_GetAuthorizeCode(t *testing.T) {

	userData := model.UserData{
		SecretKey:   "de3c1ccd0e25901befb05cd326aa5257b85463d538cea7156c6744ad75e0a29ebd95ffad",
		ClientId:    "de3c1ccd0e25901befb05cd326aa5257",
		RedirectUrl: "https://fineroot1253.tistory.com/",
	}

	authData := model.UserAuthData{
		UserId:  "dark1253@naver.com",
		UserPwd: "mayora13^^",
	}

	service, err := NewService(
		"/Users/hongjungeun/go/src/github.com/fineroot1253/tistoryAPI/bin/chromeDriver/chromedriver",
		7777,
		userData,
		[]selenium.ServiceOption{selenium.Output(os.Stderr)},
	)
	if err != nil {
		log.Panicln(err)
		return
	}

	type args struct {
		authData  model.UserAuthData
		sleepTime time.Duration
		debugMode bool
	}
	tests := []struct {
		name    string
		service Service
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "성공 테스트:[success]",
			service: service,
			args: args{
				authData:  authData,
				sleepTime: 1,
				debugMode: true,
			},
		},
		{},
	}
	for _, tt := range tests {
		code, err := tt.service.GetAuthorizeCode(tt.args.authData, tt.args.sleepTime, tt.args.debugMode)
		if err != nil {
			log.Panicln(err)
		}
		assert.NotEmpty(t, code)
		log.Println(tt.name, " ==> auth code: ", code)
	}
}
