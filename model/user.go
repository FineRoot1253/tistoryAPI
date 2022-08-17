package model

// UserData 유저 데이터 DTO
// ClientId Tistory Open API에 블로그 등록시 발급받는 ClientId
// SecretKey Tistory Open API에 블로그 등록시 발급받는 SecretKey
// RedirectUrl Tistory Open API에 블로그 등록시 발급받는 RedirectUrl
// AuthorizationCode Selenium을 통해 얻어야하는 AuthorizationCode
type UserData struct {
	// ClientId Tistory Open API에 블로그 등록시 발급받는 ClientId
	ClientId string `json:"client_id"`

	// SecretKey Tistory Open API에 블로그 등록시 발급받는 SecretKey
	SecretKey string `json:"secret_key"`

	// RedirectUrl Tistory Open API에 블로그 등록시 발급받는 RedirectUrl
	RedirectUrl string `json:"redirect_url"`

	// AuthorizationCode Selenium을 통해 얻어야하는 AuthorizationCode
	AuthorizationCode string `json:"authorization_code"`
}

// UserAuthData 유저 인증용 데이터 DTO
// UserId 유저 아이디
// UserPwd 유저 패스워드
type UserAuthData struct {
	// UserId 유저 아이디
	UserId string `json:"user_id"`

	// UserPwd 유저 패스워드
	UserPwd string `json:"user_pwd"`
}
