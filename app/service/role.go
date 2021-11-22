package service

import (
	"bieshu-oa/app/dao"
	"bieshu-oa/app/model"
	"bieshu-oa/library/service"
	"context"
	"errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

var RoleService = roleService{}

type roleService struct{}

// Index 列表
func (s *roleService) Index(ctx context.Context, input model.RoleIndexInput) (output *model.RoleIndexOutput, err error) {

	var m = dao.Role.Ctx(ctx)

	output = &model.RoleIndexOutput{
		Page:  input.Page,
		Limit: input.Limit,
	}

	listModel := m.Fields(model.RoleList{}).Page(input.Page, input.Limit)

	if err := listModel.Scan(&output.List); err != nil {
		return output, err
	}

	// 总条数
	if output.Total, err = m.Count(); err != nil {
		return output, err
	}

	// 配置
	output.Config = g.Map{
		"dps": g.Cfg("sys").GetMap("dps"),
	}

	return
}

// Store 新增
func (s *roleService) Store(ctx context.Context, input model.RoleStoreInput) error {

	var m = dao.Role.Ctx(ctx)

	// 验证角色名称是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.Role.Columns.Name, input.Name, 0); err != nil {
		return err
	}

	// 保存角色权限
	if err := s.SaveAndSetMenu(ctx, input, 0); err != nil {
		return err
	}

	return nil

}

// Update 更新
func (s *roleService) Update(ctx context.Context, input model.RoleUpdateInput) error {

	var m = dao.Role.Ctx(ctx)

	// 验证ID是否存在
	if err := service.PublicService.CheckIdExist(m, input.Id); err != nil {
		return err
	}

	// 验证角色名称是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.Role.Columns.Name, input.Name, input.Id); err != nil {
		return err
	}

	// 保存角色权限
	if err := s.SaveAndSetMenu(ctx, model.RoleStoreInput{
		Name:  input.Name,
		Menus: input.Menus,
		Dp:    input.Dp,
	}, input.Id); err != nil {
		return err
	}

	return nil
}

// Delete 删除
func (s *roleService) Delete(ctx context.Context, id int) error {

	var m = dao.Role.Ctx(ctx)

	// 验证ID是否存在
	if err := service.PublicService.CheckIdExist(m, id); err != nil {
		return err
	}

	// 验证是否存在管路员
	if num := s.GetRoleAdminNum(ctx, id); num > 0 {
		return errors.New(g.I18n().Tf(ctx, `id.exist.admin`, id, num))
	}

	if err := service.PublicService.Delete(m, id); err != nil {
		return err
	}

	return nil
}

// SaveAndSetMenu 保存并设置角色权限
func (s *roleService) SaveAndSetMenu(ctx context.Context, input model.RoleStoreInput, id int) error {

	arrMenu := strings.Split(input.Menus, `,`)

	// 验证是否存在 menu id
	if num, _ := dao.Menu.Ctx(ctx).WhereIn("id", arrMenu).Count(); num != len(arrMenu) {
		return errors.New(g.I18n().Tf(ctx, `set.role.menus.not.exist`))
	}

	// 开启事务
	if err := dao.Role.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		// 判断是新增还是更新
		if id == 0 {
			// 新增
			// 保存并且获取id
			if lastInsertId, err := service.PublicService.StoreAndGetId(dao.Role.Ctx(ctx), gconv.Map(input)); err != nil {
				return err
			} else {
				id = int(lastInsertId)
			}
		} else {
			// 更新
			if err := service.PublicService.Update(dao.Role.Ctx(ctx), gconv.Map(model.RoleUpdateInput{
				Id:   id,
				Name: input.Name,
				Dp:   input.Dp,
			})); err != nil {
				return err
			}
		}

		// 处理多个数据插入
		var data g.List
		for _, val := range arrMenu {
			data = append(data, g.Map{
				dao.RoleMenu.Columns.RoleId: id,
				dao.RoleMenu.Columns.MenuId: val,
			})
		}

		// 先删除
		if _, err := dao.RoleMenu.Ctx(ctx).Where("role_id = ?", id).Delete(); err != nil {
			return err
		}

		// 批量写入
		if _, err := dao.RoleMenu.Ctx(ctx).Data(data).Insert(); err != nil {
			return err
		}

		return nil

	}); err != nil {
		return err
	}

	return nil

}

// Show 详情
func (s *roleService) Show(ctx context.Context, id int) (output *model.RoleShowOutput, err error) {

	var m = dao.Role.Ctx(ctx)

	output = &model.RoleShowOutput{}

	// 验证ID是否存在
	if err := service.PublicService.CheckIdExist(m, id); err != nil {
		return output, err
	}

	// 获取详情
	if err := m.Where("id=?", id).Scan(output); err != nil {
		return output, err
	}

	// 获取角色权限
	if output.Menus, err = s.GetRoleMenu(ctx, id); err != nil {
		return output, err
	}

	return
}

// GetRoleMenu 获取角色权限
func (s *roleService) GetRoleMenu(ctx context.Context, id int) ([]string, error) {

	menu, err := dao.RoleMenu.Ctx(ctx).Fields(dao.RoleMenu.Columns.MenuId).Where("role_id", id).Array()
	if err != nil {
		return nil, err
	}

	return gconv.Strings(menu), nil
}

// GetRoleAdminNum 获取绑定角色的管理员人数
func (s *roleService) GetRoleAdminNum(ctx context.Context, id int) int {

	if count, _ := dao.AdminRole.Ctx(ctx).Where("role_id", id).Count(); count > 0 {
		return count
	}

	return 0
}
