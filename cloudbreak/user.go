package cloudbreak

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/cloudbreak/v2/client_cloudbreak/v1users"
	"github.com/hortonworks/cb-cli/cloudbreak/v2/utils"
	"github.com/hortonworks/cb-cli/cloudbreak/v2/models_cloudbreak"
	"github.com/urfave/cli"
)

var userListHeader = []string{"User ID", "Name"}

type userListOut struct {
	User *models_cloudbreak.UserResponseJSON
}

func (u *userListOut) DataAsStringArray() []string {
	return []string{u.User.UserID, u.User.Username}
}

func ListUsers(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "list users")
	log.Infof("[ListUsers] List all users in a tenant")
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}

	cbClient := NewCloudbreakHTTPClientFromContext(c)

	resp, err := cbClient.Cloudbreak.V1users.GetAllUsers(v1users.NewGetAllUsersParams())
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	tableRows := []utils.Row{}
	for _, user := range resp.Payload {
		tableRows = append(tableRows, &userListOut{user})
	}
	output.WriteList(userListHeader, tableRows)
}
