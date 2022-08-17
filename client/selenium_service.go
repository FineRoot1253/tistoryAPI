package client

import (
	"fmt"

	"github.com/fineroot1253/tistoryAPI"
	"github.com/fineroot1253/tistoryAPI/model"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

type Service interface {
	GetAuthorizeCode(authData model.UserAuthData, sleepTime time.Duration, debugMode bool) (string, error)
}

type service struct {
	driverPath          string
	port                int
	startUrl            string
	chromeDriverService selenium.Service
}

func NewService(driverPath string, port int, userData model.UserData, serviceOpts []selenium.ServiceOption) (Service, error) {

	// 초기 로딩 URL
	startUrl := tistoryAPI.TISTORY_OAUTH_AUTHENTICATIONTOKEN_GET_PATH +
		"?client_id=" + userData.ClientId +
		"&redirect_uri=" + userData.RedirectUrl +
		"&response_type=code"

	// option을 통해 셀레니움 초기화
	// 크롬만 쓰도록 강제하기 위해 드라이버 서비스만 주입받도록 만들지 않았다.
	// 크롬인지 검사는 로직을 넣는건 오버 엔지니어링이기 땜시롱
	chromeDriverService, err2 := selenium.NewChromeDriverService(driverPath, port, serviceOpts...)
	if err2 != nil {
		return nil, err2
	}

	return service{driverPath: driverPath, port: port, startUrl: startUrl, chromeDriverService: *chromeDriverService}, nil
}

// GetAuthorizeCode 크롬 드라이버를 통해 Authentication code를 구하는 연산
// id, pwd를 통해 로그인을 하고 auth를 얻는다.
// sleepTime을 통해 Sec 단위로 대기 시간을 조절할 수 있다.
// 너무 짦으면 현재 드라이버가 막힐 수도 있다. 이땐 새로운 셀레니움 서비스를 얻어야한다.
// debugMode를 통해 유닛 테스트으로써 실행후 드라이버를 종료할 것인지 정할 수 있다.
func (s service) GetAuthorizeCode(authData model.UserAuthData, sleepTime time.Duration, debugMode bool) (string, error) {

	authorizeCode := ""

	if debugMode {
		// 크롬 드라이버 실행
		defer s.chromeDriverService.Stop()
	}

	// 크롬 드라이버 접속 설정
	caps := selenium.Capabilities{"browserName": "chrome"}

	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			"--headless",   // 창 띄우지 않기
			"--no-sandbox", //
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
		},
	}

	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", s.port))
	if err != nil {
		log.Panicln(err)
		return authorizeCode, err
	}

	defer wd.Quit()

	// 초기로딩: 로그인 화면을 띄우기 위함
	if err := wd.Get(s.startUrl); err != nil {
		log.Panicln(err)
		return authorizeCode, err
	}
	time.Sleep(sleepTime * time.Second)
	initLoginBtn, err := wd.FindElement(selenium.ByCSSSelector, `a.btn_login.link_kakao_id`)
	if err != nil {
		log.Panicln(err)
		return authorizeCode, err
	}
	initLoginBtn.Click()
	time.Sleep(sleepTime * time.Second)

	// ID SET
	idElement, err2 := wd.FindElement(selenium.ByID, `id_email_2`)
	if err2 != nil {
		source, err := wd.PageSource()
		if err != nil {
			log.Panicln(err)
			return authorizeCode, err2
		}
		log.Panicln(source)
		log.Panicln(err2)
		return authorizeCode, err2
	}
	idElement.SendKeys(authData.UserId)
	time.Sleep(sleepTime * time.Second)

	// PWD SET
	pwdElement, err := wd.FindElement(selenium.ByID, `id_password_3`)
	if err != nil {

		log.Panicln(err)
		return authorizeCode, err
	}
	pwdElement.SendKeys(authData.UserPwd)
	time.Sleep(sleepTime * time.Second)

	// CLICK LOGIN
	loginBtn, err := wd.FindElement(selenium.ByXPATH, `//*[@id="login-form"]/fieldset/div[8]/button[1]`)
	if err != nil {
		log.Panicln(err)
		return authorizeCode, err
	}
	loginBtn.Click()
	time.Sleep(sleepTime * time.Second)
	resultUrl1, err := wd.CurrentURL()
	if err != nil {
		log.Panicln(err)
		return authorizeCode, err
	}
	log.Println("CurrentURL: ", resultUrl1)

	// 초기로딩: 로그인 화면을 띄우기 위함
	if err := wd.Get(s.startUrl); err != nil {
		log.Panicln(err)
		return authorizeCode, err
	}

	confirmBtn, err := wd.FindElement(selenium.ByClassName, `confirm`)
	if err != nil {

		log.Panicln(err)
		return authorizeCode, err
	}
	confirmBtn.Click()
	time.Sleep(sleepTime * time.Second)

	// 허가코드 결과는 Location에 쿼리스트링으로 붙어있다.
	// 결과URL를 긁어 쿼리스트링을 파싱한다.
	resultUrl, err := wd.CurrentURL()
	if err != nil {
		log.Panicln(err)
		return authorizeCode, err
	}
	u, err := url.Parse(resultUrl)
	if err != nil {
		log.Panicln(err)
		return authorizeCode, err
	}
	parseQuery, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		log.Panicln(err)
		return authorizeCode, err
	}

	authorizeCode = parseQuery.Get("code")

	return authorizeCode, nil

}

func genTxId(appKey string) (string, error) {

	var resultStr string

	prefix, err2 := getRandomString()
	if err2 != nil {
		return resultStr, err2
	}

	suffix := strconv.FormatUint(uint64(time.Now().UnixMilli()), 36)

	resultStr = prefix + appKey + suffix

	return resultStr, nil

}

// StateToken 생성기
func genStateToken() (string, error) {

	var resultStr string

	prefix, err1 := getRandomString()
	if err1 != nil {
		return resultStr, err1
	}

	suffix, err2 := getRandomString()
	if err2 != nil {
		return resultStr, err2
	}

	resultStr = prefix + suffix

	return resultStr, nil
}

// Float64기준 소수자리 숫자 => Base36 문자화 하는 로직
func getRandomString() (string, error) {
	var resultStr string

	randNum := rand.Float64()
	randStrWithDot := fmt.Sprintf("%v", randNum)
	randStr := randStrWithDot[2:]
	parseUint, err := strconv.ParseUint(randStr, 10, 0)
	if err != nil {
		return resultStr, err
	}
	resultStr = strconv.FormatUint(parseUint, 36)

	return resultStr, nil
}

//func var tranId = Math.random().toString(36).slice(2) + getAppKey$1() + Date.now().toString(36);
//return tranId.slice(0, 60);
