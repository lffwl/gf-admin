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
	"strconv"
)

var GroupService = groupService{}

type groupService struct{}

// Index 列表
func (s *groupService) Index(ctx context.Context, input model.GroupIndexInput) (output *model.GroupIndexOutput, err error) {

	var m = dao.Group.Ctx(ctx)

	output = &model.GroupIndexOutput{}

	listModel := m.Fields(model.GroupList{})

	if err := listModel.Scan(&output.List); err != nil {
		return output, err
	}

	return
}

// Store 新增
func (s *groupService) Store(ctx context.Context, input model.GroupStoreInput) error {

	var m = dao.Group.Ctx(ctx)

	// 验证PID
	if input.Pid != 0 {
		// 验证组织架构PID是否存在
		if err := service.PublicService.CheckIdExist(m, input.Pid); err != nil {
			return errors.New(g.I18n().Tf(ctx, `pid.not.exist`, input.Pid))
		}

		// 处理link路径，获取PID的Link路径 + 自己的
		input.Link = s.GetLink(ctx, input.Pid) + gconv.String(input.Pid) + ":"
	} else {
		// 处理link路径
		input.Link = "0:"
	}

	// 验证组织架构名称是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.Group.Columns.Name, input.Name, 0); err != nil {
		return err
	}

	if err := service.PublicService.Store(m, gconv.Map(input)); err != nil {
		return err
	}

	return nil

}

// Update 更新
func (s *groupService) Update(ctx context.Context, input model.GroupUpdateInput) error {

	var m = dao.Group.Ctx(ctx)

	// 验证ID是否存在
	if err := service.PublicService.CheckIdExist(m, input.Id); err != nil {
		return err
	}

	// 验证PID
	if input.Pid != 0 {
		// 验证组织架构PID是否存在
		if err := service.PublicService.CheckIdExist(m, input.Pid); err != nil {
			return errors.New(g.I18n().Tf(ctx, `pid.not.exist`, input.Pid))
		}

		// 处理link路径，获取PID的Link路径 + 自己的
		input.Link = s.GetLink(ctx, input.Pid) + gconv.String(input.Pid) + ":"
	} else {
		// 处理link路径
		input.Link = "0:"
	}

	// 验证组织架构名称是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.Group.Columns.Name, input.Name, input.Id); err != nil {
		return err
	}

	// 获取原来的PID
	info, _ := m.Fields(dao.Group.Columns.Pid, dao.Group.Columns.Link).Where("id=?", input.Id).One()

	// 检查是否修改了PID
	if gconv.Int(info["pid"]) != input.Pid {
		//开启事务
		if err := dao.Group.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
			// 更新
			if err := service.PublicService.Update(dao.Group.Ctx(ctx), gconv.Map(input)); err != nil {
				return err
			}

			// 更新下级的LINK
			str := gconv.String(info["link"]) + gconv.String(input.Id) + ":"
			if _, err := dao.Group.Ctx(ctx).Where("link like ?", str+"%").Data(g.Map{
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
func (s *groupService) Delete(ctx context.Context, id int) error {

	var m = dao.Group.Ctx(ctx)

	// 验证ID是否存在
	if err := service.PublicService.CheckIdExist(m, id); err != nil {
		return err
	}

	// 检查是否存在子集
	if num := s.GetIdSubsetNum(ctx, id); num > 0 {
		return errors.New(g.I18n().Tf(ctx, `id.exist.subset`, id, num))
	}

	// 验证是否存在管路员
	if num := s.GetGroupAdminNum(ctx, id); num > 0 {
		return errors.New(g.I18n().Tf(ctx, `id.exist.admin`, id, num))
	}

	if err := service.PublicService.Delete(m, id); err != nil {
		return err
	}

	return nil
}

// GetLink 获取id的Link地址
func (s *groupService) GetLink(ctx context.Context, id int) string {

	link, _ := dao.Group.Ctx(ctx).Fields(dao.Group.Columns.Link).Where("id = ?", id).Value()

	return gconv.String(link)
}

// GetIdSubsetNum 获取ID的下级子集个数
func (s *groupService) GetIdSubsetNum(ctx context.Context, id int) int {

	if count, _ := dao.Group.Ctx(ctx).Where("link like ?", "%:"+gconv.String(id)+":%").Count(); count > 0 {
		return count
	}

	return 0
}

// GetGroupAdminNum 获取组织架构人数
func (s *groupService) GetGroupAdminNum(ctx context.Context, id int) int {

	if count, _ := dao.Admin.Ctx(ctx).Where("group_id", id).Count(); count > 0 {
		return count
	}

	return 0
}

// GetGroupAndSubsetId 获取自己和下级组织构架构Id
func (s *groupService) GetGroupAndSubsetId(ctx context.Context, id int) []string {

	// 获取id下的组织架构
	ids, _ := dao.Group.Ctx(ctx).Fields("id").Where("link like ?", "%:"+gconv.String(id)+":%").Array()

	sliceId := gconv.SliceStr(ids)

	// 追加值
	sliceId = append(sliceId, strconv.Itoa(id))

	return sliceId
}

// GetGroupIdsAdminId 获取组织架构下的管理员ID
func (s *groupService) GetGroupIdsAdminId(ctx context.Context, ids []string) []string {

	adminIds, _ := dao.Admin.Ctx(ctx).Fields("id").WhereIn(dao.Admin.Columns.GroupId, ids).Array()

	return gconv.SliceStr(adminIds)

}
