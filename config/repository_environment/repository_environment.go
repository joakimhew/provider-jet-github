package repositoryenvironment

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_repository_environment", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "github"
		r.ShortGroup = "environment"
		r.Kind = "RepositoryEnvironment"
		r.ExternalName = config.IdentifierFromProvider

		r.References["reviewers.teams"] = config.Reference{
			Type: "github.com/joakimhew/provider-jet-github/apis/team/v1alpha1.Team",
		}
	})
}
