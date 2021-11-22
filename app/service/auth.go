package service

import (
	"bieshu-oa/app/dao"
	"bieshu-oa/app/define"
	"bieshu-oa/app/model"
	"bieshu-oa/library/helper"
	"context"
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"golang.org/x/crypto/bcrypt"
)

var AuthService = authService{}

type authService struct{}

// Login 登录
func (s *authService) Login(ctx context.Context, input model.AuthLoginInput) (g.Map, error) {

	// 验证管理员是否存在
	id := s.CheckUserNameExistAndGetId(ctx, input.UserName)
	if id == 0 {
		return nil, errors.New(g.I18n().Tf(ctx, `auth.user-name.error`, input.UserName))
	}

	// 验证管理员密码是否正确
	if s.CheckPassword(ctx, id, input.Password) == false {
		return nil, errors.New(g.I18n().Tf(ctx, `auth.password.error`))
	}

	return g.Map{
		"username": input.UserName,
		"id":       id,
	}, nil

}

// CheckUserNameExistAndGetId 验证管理员是否存在并且获取管理员id
func (s *authService) CheckUserNameExistAndGetId(ctx context.Context, userName string) int {

	id, err := dao.Admin.Ctx(ctx).Fields(dao.Admin.Columns.Id).Where(dao.Admin.Columns.UserName, userName).Value()
	if err != nil {
		return 0
	}

	resId := gconv.Int(id)

	return resId
}

// CheckPassword 验证管理员密码是否正确
func (s *authService) CheckPassword(ctx context.Context, id int, loginPassword string) bool {

	// 获取密码
	password, err := dao.Admin.Ctx(ctx).Fields(dao.Admin.Columns.Password).Where("id", id).Value()
	if err != nil {
		return false
	}

	//验证密码是否正确
	if err := bcrypt.CompareHashAndPassword(gconv.Bytes(password), gconv.Bytes(loginPassword)); err != nil {
		return false
	}

	return true
}

// Info 获取详情
func (s *authService) Info(ctx context.Context, id int) (g.Map, error) {

	// 获取管理员详情
	info, err := AdminService.Show(ctx, id)
	if err != nil {
		return nil, err
	}

	data := gconv.Map(info)

	// 检查是否超级管理员
	if helper.Array.In(gconv.String(define.SuperAdminId), info.Roles) {

		// 获取全部的菜单
		menus, err := MenuService.Index(ctx, model.MenuIndexInput{}, false)
		if err != nil {
			return nil, err
		}

		data["auth"] = menus.List

	} else {

		// 是否存在roles
		if info.Roles != nil {
			if data["auth"], err = g.DB().Ctx(ctx).GetAll("SELECT `id`,`name`,`key`,`pid` FROM menu WHERE id IN ( SELECT DISTINCT menu_id FROM role_menu WHERE role_id IN ( ? ) ) and deleted_at  IS NULL", info.Roles); err != nil {
				return nil, err
			}
		}
	}

	return data, nil

}

// CheckAdminRouter 检查是否有权限访问当前地址
func (s *authService) CheckAdminRouter(ctx context.Context, router string, method string) bool {

	// 检查是否不需要权限
	notCheckAdminRouters := g.SliceStr{"/auth/info"}
	if helper.Array.In(router, notCheckAdminRouters) {
		return true
	}

	// 获取当前登录的ID
	id := gconv.Int(ctx.Value(define.CurrentAdminId))

	// 获取管理员角色
	roles, err := AdminService.GetAdminRole(ctx, id)
	if err != nil {
		return false
	}

	// 检查是否超级管理员
	if helper.Array.In(gconv.String(define.SuperAdminId), roles) {
		return true
	}

	// 返回是否存在权限
	return MenuService.CheckRoleRouter(ctx, roles, router, method)
}
