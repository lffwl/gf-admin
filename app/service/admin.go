package service

import (
	"bieshu-oa/app/dao"
	"bieshu-oa/app/define"
	"bieshu-oa/app/model"
	"bieshu-oa/library/helper"
	"bieshu-oa/library/service"
	"context"
	"errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

var AdminService = adminService{}

type adminService struct {
}

// Index 列表
func (s *adminService) Index(ctx context.Context, input model.AdminIndexInput) (output *model.AdminIndexOutput, err error) {

	var m = dao.Admin.Ctx(ctx)

	output = &model.AdminIndexOutput{
		Page:  input.Page,
		Limit: input.Limit,
	}

	// 条件检索
	if input.UserName != "" {
		m = m.Where("user_name like ?", input.UserName+"%")
	}
	if input.RealName != "" {
		m = m.Where("real_name like ?", input.RealName+"%")
	}
	if input.Status != "" {
		m = m.Where("status=?", input.Status)
	}
	if input.Mobile != "" {
		m = m.Where("mobile like ?", input.Mobile+"%")
	}
	if input.Email != "" {
		m = m.Where("email like ?", input.Email+"%")
	}

	// 数据权限处理
	admins := gconv.SliceStr(ctx.Value(define.CurrentAdminData)) // 获取可以管理的数据权限
	// 不存在全部数据权限需要处理数据
	if helper.Array.In(define.AllData, admins) == false {
		m = m.WhereIn("id", admins)
	}

	listModel := m.Fields(model.AdminList{}).Page(input.Page, input.Limit)

	if err := listModel.Scan(&output.List); err != nil {
		return output, err
	}

	// 总条数
	if output.Total, err = m.Count(); err != nil {
		return output, err
	}

	// 获取角色列表
	if err := dao.Role.Ctx(ctx).Fields(model.AdminRoleList{}).Scan(&output.Roles); err != nil {
		return output, err
	}

	// 获取组织架构列表
	if err := dao.Group.Ctx(ctx).Fields(model.AdminGroupList{}).Scan(&output.Groups); err != nil {
		return output, err
	}

	// 配置
	output.Config = g.Map{
		"admin_status": g.Cfg("sys").GetMap("adminStatus"),
	}

	return
}

// Show 管理员详情
func (s *adminService) Show(ctx context.Context, id int) (output *model.AdminShowOutput, err error) {

	var m = dao.Admin.Ctx(ctx)

	output = &model.AdminShowOutput{}

	// 数据权限处理
	admins := gconv.SliceStr(ctx.Value(define.CurrentAdminData)) // 获取可以管理的数据权限
	// 不存在全部数据权限需要处理数据
	if helper.Array.In(define.AllData, admins) == false {
		m = m.WhereIn("id", admins)
	}

	// 验证ID是否存在
	if err := service.PublicService.CheckIdExist(m, id); err != nil {
		return output, err
	}

	// 获取详情
	if err := m.Fields(model.AdminList{}).Where("id", id).Scan(&output); err != nil {
		return output, err
	}

	// 获取管理员角色
	if output.Roles, err = s.GetAdminRole(ctx, id); err != nil {
		return output, err
	}

	return
}

// Store 新增
func (s *adminService) Store(ctx context.Context, input model.AdminStoreInput) error {

	var m = dao.Admin.Ctx(ctx)

	// 验证管理员名称是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.Admin.Columns.UserName, input.UserName, 0); err != nil {
		return err
	}

	// 保存管理员权限
	if err := s.SaveAndSetRole(ctx, input, 0); err != nil {
		return err
	}

	return nil

}

// Update 更新
func (s *adminService) Update(ctx context.Context, input model.AdminUpdateInput) error {

	var m = dao.Admin.Ctx(ctx)

	// 数据权限处理
	admins := gconv.SliceStr(ctx.Value(define.CurrentAdminData)) // 获取可以管理的数据权限
	// 不存在全部数据权限需要处理数据
	if helper.Array.In(define.AllData, admins) == false {
		m = m.WhereIn("id", admins)
	}

	// 验证ID是否存在
	if err := service.PublicService.CheckIdExist(m, input.Id); err != nil {
		return err
	}

	// 验证管理员名称是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.Admin.Columns.UserName, input.UserName, input.Id); err != nil {
		return err
	}

	// 保存管理员权限
	if err := s.SaveAndSetRole(ctx, model.AdminStoreInput{
		UserName: input.UserName,
		Password: input.Password,
		Avatar:   input.Avatar,
		RealName: input.RealName,
		Mobile:   input.Mobile,
		Email:    input.Email,
		Status:   input.Status,
		Roles:    input.Roles,
		GroupId:  input.GroupId,
	}, input.Id); err != nil {
		return err
	}

	return nil
}

// Delete 删除
func (s *adminService) Delete(ctx context.Context, id int) error {

	var m = dao.Admin.Ctx(ctx)

	// 数据权限处理
	admins := gconv.SliceStr(ctx.Value(define.CurrentAdminData)) // 获取可以管理的数据权限
	// 不存在全部数据权限需要处理数据
	if helper.Array.In(define.AllData, admins) == false {
		m = m.WhereIn("id", admins)
	}

	// 验证ID是否存在
	if err := service.PublicService.CheckIdExist(m, id); err != nil {
		return err
	}

	// 开启事务
	if err := dao.Admin.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		// 删除管理员
		if err := service.PublicService.Delete(dao.Admin.Ctx(ctx), id); err != nil {
			return err
		}

		// 删除管理员角色关联
		if _, err := dao.AdminRole.Ctx(ctx).Where("admin_id", id).Delete(); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// SaveAndSetRole 保存并设置管理员权限
func (s *adminService) SaveAndSetRole(ctx context.Context, input model.AdminStoreInput, id int) error {

	// 密码的处理
	if input.Password != "" {
		// 密码加密
		if hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost); err != nil {
			return err
		} else {
			input.Password = string(hash)
		}
	}

	arrRole := strings.Split(input.Roles, `,`)

	// 验证是否存在 role id
	if num, _ := dao.Role.Ctx(ctx).WhereIn("id", arrRole).Count(); num != len(arrRole) {
		return errors.New(g.I18n().Tf(ctx, `set-admin.role.not.exist`))
	}

	// 开启事务
	if err := dao.Admin.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		// 判断是新增还是更新
		if id == 0 {
			// 新增
			// 保存并且获取id
			if lastInsertId, err := service.PublicService.StoreAndGetId(dao.Admin.Ctx(ctx), gconv.Map(input)); err != nil {
				return err
			} else {
				id = int(lastInsertId)
			}
		} else {
			updateData := g.Map{
				"Id":       id,
				"UserName": input.UserName,
				"Avatar":   input.Avatar,
				"RealName": input.RealName,
				"Mobile":   input.Mobile,
				"Email":    input.Email,
				"Status":   input.Status,
				"Roles":    input.Roles,
				"GroupId":  input.GroupId,
			}

			// 密码不为空就需要更新密码
			if input.Password != "" {
				updateData["Password"] = input.Password
			}

			// 更新
			if err := service.PublicService.Update(dao.Admin.Ctx(ctx), updateData); err != nil {
				return err
			}
		}

		// 处理多个数据插入
		var data g.List
		for _, val := range arrRole {
			data = append(data, g.Map{
				dao.AdminRole.Columns.AdminId: id,
				dao.AdminRole.Columns.RoleId:  val,
			})
		}

		// 先删除
		if _, err := dao.AdminRole.Ctx(ctx).Where("admin_id = ?", id).Delete(); err != nil {
			return err
		}

		// 批量写入
		if _, err := dao.AdminRole.Ctx(ctx).Data(data).Insert(); err != nil {
			return err
		}

		return nil

	}); err != nil {
		return err
	}

	return nil

}

// GetAdminRole 获取管理员角色
func (s *adminService) GetAdminRole(ctx context.Context, id int) ([]string, error) {

	role, err := dao.AdminRole.Ctx(ctx).Fields(dao.AdminRole.Columns.RoleId).Where("admin_id", id).Array()
	if err != nil {
		return nil, err
	}

	return gconv.Strings(role), nil
}

// GetAdminDataAdminId 获取管理员有数据权限的管理员ID
func (s *adminService) GetAdminDataAdminId(ctx context.Context, id int) ([]string, error) {

	var ids []string

	// 获取管理员角色
	roles, err := s.GetAdminRole(ctx, id)
	if err != nil {
		return nil, err
	}

	// 检查是否超级管理员
	if helper.Array.In(gconv.String(define.SuperAdminId), roles) {
		ids = append(ids, define.AllData)
		return ids, nil
	}

	// 数据权限集合
	var dpInt []int

	// 角色数据权限处理
	for _, role := range roles {

		// 获取角色的数据权限
		dp, err := dao.Role.Ctx(ctx).Where("id", role).Fields(dao.Role.Columns.Dp).Value()
		if err != nil {
			return nil, err
		}

		// to int
		dpInt = append(dpInt, gconv.Int(dp))

	}

	// 获取最大的数据权限
	dp := helper.Array.MaxSliceInt(dpInt)

	// 验证是否拥有所有人数据权限
	if dp == define.DpAllDataVal {
		return g.SliceStr{define.AllData}, nil
	} else if dp == define.DpPartDataVal {
		// 拥有所在组织架构和下级组织架构，数据权限

		// 获取管理员组织架构ID
		groupId, err := dao.Admin.Ctx(ctx).Where("id", id).Fields(dao.Admin.Columns.GroupId).Value()
		if err != nil {
			return nil, err
		}

		// 获取组织架构
		groups := GroupService.GetGroupAndSubsetId(ctx, gconv.Int(groupId))

		// 获取组织架构下的全部管理员
		admins := GroupService.GetGroupIdsAdminId(ctx, groups)
		ids = append(ids, admins...)
	} else if dp == define.DpThisDataVal {
		// 仅自己,数据权限
		ids = append(ids, gconv.String(id))
	}

	return ids, err

}
