package server

import (
	"context"
	"log"

	"gomod.pri/helloGrpc/pb/xcommon"
	authipb "gomod.pri/helloGrpc/pb/xinternal/auth"
)

func newAuthServer() authipb.AuthServiceServer {
	return new(authServiceServer)
}

type authServiceServer struct {
	authipb.UnimplementedAuthServiceServer
}

// ListUserGroupByRoleID 实现了 AuthServiceServer 接口中的 ListUserGroupByRoleID 方法
func (s *authServiceServer) ListUserGroupByRoleID(ctx context.Context, in *authipb.ListUserGroupByRoleIDReq) (*authipb.ListUserGroupByRoleIDResp, error) {
	log.Printf("用户 %d, 你的角色ID是 %d", in.GetCommonIn().UserId, in.RoleId)
	return &authipb.ListUserGroupByRoleIDResp{UserGroupInfos: []*xcommon.MarkMeta{&xcommon.MarkMeta{
		Id:   10086,
		Name: "tom",
		Mark: xcommon.Mark_MARK_ON,
	}}}, nil
}
