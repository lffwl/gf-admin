package service

import (
	"bieshu-oa/app/dao"
	"bieshu-oa/app/model"
	"bieshu-oa/library/helper"
	"bieshu-oa/library/service"
	"context"
	"errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var MenuService = menuService{}

type menuService struct{}

// Index 列表
func (s *menuService) Index(ctx context.Context, input model.MenuIndexInput, config bool) (output *model.MenuIndexOutput, err error) {

	var m = dao.Menu.Ctx(ctx)

	output = &model.MenuIndexOutput{}

	listModel := m.Fields(model.MenuList{})

	if err := listModel.Scan(&output.List); err != nil {
		return output, err
	}

	// 配置输出
	if config {
		output.Config = g.Map{
			"menu_methods": g.Cfg("sys").GetMap("menuMethods"),
			"menu_types":   g.Cfg("sys").GetMap("menuTypes"),
		}
	}

	return
}

// Store 新增
func (s *menuService) Store(ctx context.Context, input model.MenuStoreInput) error {

	var m = dao.Menu.Ctx(ctx)

	// 验证PID
	if input.Pid != 0 {
		// 验证菜单PID是否存在
		if err := service.PublicService.CheckIdExist(m, input.Pid); err != nil {
			return errors.New(g.I18n().Tf(ctx, `pid.not.exist`, input.Pid))
		}

		// 处理link路径，获取PID的Link路径 + 自己的
		input.Link = s.GetLink(ctx, input.Pid) + gconv.String(input.Pid) + ":"
	} else {
		// 处理link路径
		input.Link = "0:"
	}

	// 验证菜单名称是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.Menu.Columns.Name, input.Name, 0); err != nil {
		return err
	}

	// 验证是否存在相同的地址和请求方式
	if input.Router != "" && input.Method != "" {
		if num, _ := m.Where("router = ? AND method = ?", input.Router, input.Method).Count(); num > 0 {
			menuMethods := g.Cfg("sys").GetMap("menuMethods")
			return errors.New(g.I18n().Tf(ctx, `menu.router.method.exist`, input.Router, menuMethods[input.Method]))
		}
	}

	if err := service.PublicService.Store(m, gconv.Map(input)); err != nil {
		return err
	}

	return nil

}

// Update 更新
func (s *menuService) Update(ctx context.Context, input model.MenuUpdateInput) error {

	var m = dao.Menu.Ctx(ctx)

	// 验证PID
	if input.Pid != 0 {
		// 验证菜单PID是否存在
		if err := service.PublicService.CheckIdExist(m, input.Pid); err != nil {
			return errors.New(g.I18n().Tf(ctx, `pid.not.exist`, input.Pid))
		}

		// 处理link路径，获取PID的Link路径 + 自己的
		input.Link = s.GetLink(ctx, input.Pid) + gconv.String(input.Pid) + ":"
	} else {
		// 处理link路径
		input.Link = "0:"
	}

	// 验证菜单名称是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.Menu.Columns.Name, input.Name, input.Id); err != nil {
		return err
	}

	// 验证是否存在相同的地址和请求方式
	if input.Router != "" && input.Method != "" {
		if num, _ := m.Where("router = ? AND method = ?", input.Router, input.Method).Where("id <> ?", input.Id).Count(); num > 0 {
			menuMethods := g.Cfg("sys").GetMap("menuMethods")
			return errors.New(g.I18n().Tf(ctx, `menu.router.method.exist`, input.Router, menuMethods[input.Method]))
		}
	}

	// 获取原来的PID
	info, _ := m.Fields(dao.Group.Columns.Pid, dao.Group.Columns.Link).Where("id=?", input.Id).One()

	// 检查是否修改了PID
	if gconv.Int(info["pid"]) != input.Pid {
		//开启事务
		if err := dao.Menu.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
			// 更新
			if err := service.PublicService.Update(dao.Menu.Ctx(ctx), gconv.Map(input)); err != nil {
				return err
			}

			// 更新下级的LINK
			str := gconv.String(info["link"]) + gconv.String(input.Id) + ":"
			if _, err := dao.Menu.Ctx(ctx).Where("link like ?", str+"%").Data(g.Map{
				"link": gdb.Raw("replace(link,'" + str + "', '" + input.Link + gconv.String(input.Id) + ":')"),
			}).Unscoped().Update(); err != nil {
				return err
			}

			return nil
		}); err != nil {
			return err
		}
	} else {
		// 没有修改PID直接更新即可
		if err := service.PublicService.Update(m, gconv.Map(input)); err != nil {
			return err
		}
	}

	return nil
}

// Delete 删除
func (s *menuService) Delete(ctx context.Context, id int) error {

	var m = dao.Menu.Ctx(ctx)

	// 验证ID是否存在
	if err := service.PublicService.CheckIdExist(m, id); err != nil {
		return err
	}

	// 检查是否存在子集
	if num := s.GetIdSubsetNum(ctx, id); num > 0 {
		return errors.New(g.I18n().Tf(ctx, `id.exist.subset`, id, num))
	}

	if err := service.PublicService.Delete(m, id); err != nil {
		return err
	}

	return nil
}

// GetLink 获取id的Link地址
func (s *menuService) GetLink(ctx context.Context, id int) string {

	link, _ := dao.Menu.Ctx(ctx).Fields(dao.Menu.Columns.Link).Where("id = ?", id).Value()

	return gconv.String(link)
}

// GetIdSubsetNum 获取ID的下级子集个数
func (s *menuService) GetIdSubsetNum(ctx context.Context, id int) int {

	if count, _ := dao.Menu.Ctx(ctx).Where("link like ?", "%:"+gconv.String(id)+":%").Count(); count > 0 {
		return count
	}

	return 0
}

// CheckRoleRouter 检查角色是否存在访问地址权限
func (s *menuService) CheckRoleRouter(ctx context.Context, roleIds []string, router string, method string) bool {
	// 配置
	menuMethods := g.Cfg("sys").GetMap("menuMethods")
	menuMethods = helper.Array.MapKeyToValue(menuMethods)
	// 检查角色是否存在访问地址权限
	if num, _ := dao.Menu.Ctx(ctx).Where("id in ?",
		dao.RoleMenu.Ctx(ctx).Fields(dao.RoleMenu.Columns.MenuId).WhereIn("role_id", roleIds),
	).Where(dao.Menu.Columns.Router, router).Where(dao.Menu.Columns.Method, menuMethods[method]).Count(); num > 0 {
		return true
	}

	return false
}
