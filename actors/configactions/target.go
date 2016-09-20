package configactions

func (actor Actor) SetTarget() (Warnings, error) {
	warnings, err := actor.CloudControllerClient.TargetCF()
	if err != nil {
		return Warnings(warnings), err
	}

	actor.Config.SetTargetInformation(
		actor.CloudControllerClient.API(),
		actor.CloudControllerClient.APIVersion(),
		actor.CloudControllerClient.AuthorizationEndpoint(),
		actor.CloudControllerClient.LoggregatorEndpoint(),
		actor.CloudControllerClient.DopplerEndpoint(),
		actor.CloudControllerClient.TokenEndpoint(),
		actor.CloudControllerClient.SkipSSLValidation(),
	)

	return Warnings(warnings), nil
}

func (actor Actor) ClearTarget() {
	actor.Config.SetTargetInformation("", "", "", "", "", "", false)
}
