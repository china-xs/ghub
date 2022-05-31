package service

import (
	"github.com/china-xs/gin-tpl/pkg/db"
	"github.com/china-xs/gin-tpl/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"

	pb "ghub/api/v1/helloword"
)

const _logKey = "hello-log"

type GreeterService struct {
	pb.UnimplementedGreeterServer
	log *zap.Logger
	rdb *redis.Client  // redis
	db  *gorm.DB       // gorm db
	tx  db.Transaction // gorm 挎包事物 tx.
}

func NewGreeterService(
	log *zap.Logger,
	rdb *redis.Client,
	db *gorm.DB,
	tx db.Transaction,
) *GreeterService {
	return &GreeterService{
		log: log,
		rdb: rdb,
		db:  db,
		tx:  tx,
	}
}

func (s *GreeterService) CreateGreeter(c *gin.Context, req *pb.CreateGreeterRequest) (*pb.CreateGreeterReply, error) {
	ctx := c.Request.Context()
	log.WithCtx(ctx, s.log).Info(_logKey, zap.Any("req", req))

	return &pb.CreateGreeterReply{}, nil
}
func (s *GreeterService) UpdateGreeter(c *gin.Context, req *pb.UpdateGreeterRequest) (*pb.UpdateGreeterReply, error) {
	return &pb.UpdateGreeterReply{}, nil
}
func (s *GreeterService) DeleteGreeter(c *gin.Context, req *pb.DeleteGreeterRequest) (*pb.DeleteGreeterReply, error) {
	ctx := c.Request.Context()
	log.WithCtx(ctx, s.log).Info(_logKey, zap.Any("req", req))
	return &pb.DeleteGreeterReply{}, nil
}
func (s *GreeterService) GetGreeter(c *gin.Context, req *pb.GetGreeterRequest) (*pb.GetGreeterReply, error) {
	return &pb.GetGreeterReply{}, nil
}
func (s *GreeterService) ListGreeter(c *gin.Context, req *pb.ListGreeterRequest) (*pb.ListGreeterReply, error) {
	return &pb.ListGreeterReply{}, nil
}
