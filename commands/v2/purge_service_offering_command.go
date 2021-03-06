package v2

import (
	"os"

	"code.cloudfoundry.org/cli/cf/cmd"
	"code.cloudfoundry.org/cli/commands"
	"code.cloudfoundry.org/cli/commands/flags"
)

type PurgeServiceOfferingCommand struct {
	RequiredArgs    flags.Service `positional-args:"yes"`
	Force           bool          `short:"f" description:"Force deletion without confirmation"`
	Provider        string        `short:"p" description:"Provider"`
	usage           interface{}   `usage:"CF_NAME purge-service-offering SERVICE [-f] [-p PROVIDER]\n\nWARNING: This operation assumes that the service broker responsible for this service offering is no longer available, and all service instances have been deleted, leaving orphan records in Cloud Foundry's database. All knowledge of the service will be removed from Cloud Foundry, including service instances and service bindings. No attempt will be made to contact the service broker; running this command without destroying the service broker will cause orphan service instances. After running this command you may want to run either delete-service-auth-token or delete-service-broker to complete the cleanup."`
	relatedCommands interface{}   `related_commands:"marketplace, purge-service-instance, service-brokers"`
}

func (_ PurgeServiceOfferingCommand) Setup(config commands.Config, ui commands.UI) error {
	return nil
}

func (_ PurgeServiceOfferingCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
