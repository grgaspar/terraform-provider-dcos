module github.com/mesosphere/terraform-provider-dcos

go 1.13

require (
	github.com/antihax/optional v0.0.0-20180407024304-ca021399b1a6
	github.com/beevik/etree v1.1.0
	github.com/dcos/client-go v0.0.0-20190910161559-e3e16c6d1484
	github.com/fatz/go-marathon v0.7.1 // indirect
	github.com/gambol99/go-marathon v0.7.2-0.20191203055606-2d3f62a40d37
	github.com/go-chi/render v1.0.1
	github.com/hashicorp/terraform v0.12.9
	github.com/imdario/mergo v0.3.7
	github.com/itchyny/gojq v0.10.3
	github.com/magiconair/properties v1.8.1
	github.com/mesos/mesos-go v0.0.10 // indirect
	github.com/mesosphere-incubator/cosmos-repo-go v0.0.0-20190919140530-1bfc03a5c181
	github.com/pquerna/ffjson v0.0.0-20190930134022-aa0246cd15f7 // indirect
	github.com/savaki/jq v0.0.0-20161209013833-0e6baecebbf8
)

replace github.com/gambol99/go-marathon => github.com/fatz/go-marathon v0.7.2-0.20191224115431-b677ec57fc07
