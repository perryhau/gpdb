package services

import (
	pb "gp_upgrade/idl"

	gpbackupUtils "github.com/greenplum-db/gpbackup/utils"
	"golang.org/x/net/context"
)

func (s *cliToHubListenerImpl) CheckSeginstall(ctx context.Context,
	in *pb.CheckSeginstallRequest) (*pb.CheckSeginstallReply, error) {

	gpbackupUtils.GetLogger().Info("starting CheckSeginstall()")

	successReply := &pb.CheckSeginstallReply{}
	return successReply, nil
}
