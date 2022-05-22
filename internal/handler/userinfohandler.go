package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"douyin/internal/logic"
	"douyin/internal/svc"
	"douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRequest
		fmt.Println(r.URL.Query())
		queryMap := r.URL.Query()
		userId, err := strconv.Atoi(queryMap["user_id"][0])
		token := queryMap["token"][0]
		/**if err := httpx.Parse(r, &req); err != nil {
			fmt.Println(err)
			httpx.Error(w, err)
			return
		}**/
		req.User_id = int64(userId)
		req.Token = token

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
