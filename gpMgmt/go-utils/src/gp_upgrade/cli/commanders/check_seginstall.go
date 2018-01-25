package commanders

import (
	"context"
	pb "gp_upgrade/idl"

	gpbackupUtils "github.com/greenplum-db/gpbackup/utils"
)

type SeginstallChecker struct {
	client pb.CliToHubClient
}

func NewSeginstallChecker(client pb.CliToHubClient) SeginstallChecker {
	return SeginstallChecker{
		client: client,
	}
}

func (req SeginstallChecker) Execute() error {
	logger := gpbackupUtils.GetLogger()
	_, err := req.client.CheckSeginstall(context.Background(),
		&pb.CheckSeginstallRequest{})
	if err != nil {
		logger.Error("ERROR - gRPC call 'check seginstall' to hub failed")
		return err
	}
	logger.Info("Check seginstall request is being processed.")
	return nil
}
