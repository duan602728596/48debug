package pocket48

type UserInfo struct {
  UserId int64 `json: "userId"`
  Nickname string `json: "nickname"`
}

type Content struct {
  UserInfo UserInfo `json: "userInfo"`
  Token string `json: "token"`
}

type LoginResponseData struct {
  Content Content `json: "content"`
}