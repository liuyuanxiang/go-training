package interfaces

import (
	"context"
	pb "section03-project/api/biz"
	"section03-project/internal/application"
	"section03-project/internal/interfaces/dto"
	"git.mysre.cn/yunlian-golang/go-hulk/boot"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedUserServiceServer

	ctx     context.Context
	appServ *application.ApplicationService
}

func NewGRPCServer(ctx context.Context, appServ *application.ApplicationService) *GRPCServer {
	return &GRPCServer{ctx: ctx, appServ: appServ}
}

// RegisterServer 注册 GRPC 服务端的服务接口处理实现
// protobuf 生成的 Server 端实现方法，是业务操作的产物，基础框架无法提前感知或准备好
// 因此框架只编排对 RegisterServer 方法的调用，方法下具体的业务实现编排，需要业务开发过程中实现
func RegisterServer(app *boot.GRPCApplication, serve *GRPCServer) error {
	app.RegisterGRPCServer = func(s *grpc.Server) {
		pb.RegisterUserServiceServer(s, serve)
	}
	app.RegisterGateway = func(c context.Context, mux *runtime.ServeMux) error {
		if err := pb.RegisterUserServiceHandlerServer(c, mux, serve); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (s *GRPCServer) CreateUser(c context.Context, req *pb.CreateUserReq) (*pb.CreateUserRsp, error) {
	id, err := s.appServ.CreateUser(&dto.UserDTO{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserRsp{
		Id:       id,
		Username: req.Username,
	}, err
}

func (s *GRPCServer) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserRsp, error) {
	user, err := s.appServ.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserRsp{
		Id:       user.ID,
		Username: user.Username,
	}, err

}

func (s *GRPCServer) ListUser(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserRsp, error) {
	userList, err := s.appServ.ListUser(int(req.PageNum), int(req.PageSize))
	response := &pb.ListUserRsp{}
	mapstructure.Decode(userList, response)
	return response, err
}
