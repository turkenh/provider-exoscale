package repository

import "github.com/crossplane-contrib/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Customize(p *config.Provider) {
    p.AddResourceConfigurator("exoscale_sks_cluster", func(r *config.Resource) {

        // we need to override the default group that terrajet generated for
        // this resource, which would be "github"  
        r.ShortGroup = "sks_cluster"
    })
}
