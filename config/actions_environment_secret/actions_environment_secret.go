package actionsenvironmentsecret

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_actions_environment_secret", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "github"
		r.ShortGroup = "secrets"
		r.Kind = "EnvironmentSecret"
		r.ExternalName = config.IdentifierFromProvider

		r.References["repository"] = config.Reference{
			Type: "github.com/joakimhew/provider-jet-github/apis/repository/v1alpha1.Repository",
		}
	})
}
