package router

import (
	"bieshu-oa/app/api"
	"bieshu-oa/app/define"
	"bieshu-oa/app/middleware"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func init() {
	s := g.Server()

	//设置全局中间件
	s.BindMiddlewareDefault(
		// 跨域请求
		middleware.BaseMiddleware.CORS,
		// 设置国际化语言
		middleware.BaseMiddleware.SetLanguage,
	)

	// 认证接口
	loginFunc := api.AuthApi.Login

	// gToken 配置
	gfToken := &gtoken.GfToken{
		LoginPath:        "/auth/login",
		LoginBeforeFunc:  loginFunc,
		LogoutPath:       "/auth/logout",
		AuthExcludePaths: g.SliceStr{"/auth/login"}, // 不拦截路径
		MultiLogin:       g.Cfg("gToken").GetBool("MultiLogin"),
		EncryptKey:       g.Cfg("gToken").GetBytes("EncryptKey"),
		AuthFailMsg:      g.Cfg("gToken").GetString("AuthFailMsg"),
	}

	s.Group("/", func(manage *ghttp.RouterGroup) {

		// 调用GOToken
		gfToken.Middleware(manage)

		// 设置登录ID
		manage.Middleware(func(r *ghttp.Request) {
			getTokenData := gfToken.GetTokenData(r)
			r.SetCtxVar(define.CurrentAdminId, gconv.Map(getTokenData.Data)["data"])
			r.Middleware.Next()
		})

		// 权限验证
		manage.Middleware(
			// 权限验证
			middleware.BaseMiddleware.CheckAdminRouter,
			// 数据权限处理
			middleware.DataMiddleware.GetDataAdminId,
		)

		// 演示站点
		manage.Middleware(middleware.BaseMiddleware.Dome)

		// 个人中心
		manage.Group("/auth", func(group *ghttp.RouterGroup) {
			group.GET("/info", api.AuthApi.Info)
		})

		// 组织架构
		manage.Group("/group", func(group *ghttp.RouterGroup) {
			group.GET("/", api.GroupApi.Index)
			group.POST("/", api.GroupApi.Store)
			group.PUT("/:id", api.GroupApi.Update)
			group.DELETE("/:id", api.GroupApi.Delete)
		})

		// 菜单
		manage.Group("/menu", func(group *ghttp.RouterGroup) {
			group.GET("/", api.MenuApi.Index)
			group.POST("/", api.MenuApi.Store)
			group.PUT("/:id", api.MenuApi.Update)
			group.DELETE("/:id", api.MenuApi.Delete)
		})

		// 角色
		manage.Group("/role", func(group *ghttp.RouterGroup) {
			group.GET("/", api.RoleApi.Index)
			group.POST("/", api.RoleApi.Store)
			group.PUT("/:id", api.RoleApi.Update)
			group.DELETE("/:id", api.RoleApi.Delete)
			group.GET("/:id", api.RoleApi.Show)
		})

		// 管理员
		manage.Group("/admin", func(group *ghttp.RouterGroup) {
			group.GET("/", api.AdminApi.Index)
			group.POST("/", api.AdminApi.Store)
			group.PUT("/:id", api.AdminApi.Update)
			group.DELETE("/:id", api.AdminApi.Delete)
			group.GET("/:id", api.AdminApi.Show)
		})

		// 文件上传
		manage.Group("/upload", func(group *ghttp.RouterGroup) {
			group.POST("/", api.UploadApi.File)
		})

	})

}
