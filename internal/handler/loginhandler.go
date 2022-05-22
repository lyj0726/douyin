package handler

import (
	"net/http"

	"douyin/internal/logic"
	"douyin/internal/svc"
	"douyin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		//手动抓取特征，parse有问题
		queryMap := r.URL.Query()
		username := queryMap["username"][0]
		password := queryMap["password"][0]
		/**if err := httpx.Parse(r, &req); err != nil {
			fmt.Println(err)
			httpx.Error(w, err)
			return
		}**/
		req.Username = username
		req.Password = password

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
