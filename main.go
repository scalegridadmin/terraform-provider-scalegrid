// Command terraform-provider-scalegrid serves the scalegridadmin/scalegrid
// Terraform provider. The provider implementation lives under
// internal/provider/; data source and resource schemas are regenerated
// from the SG API OpenAPI spec by tfplugingen-framework.
package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/scalegridadmin/terraform-provider-scalegrid/internal/provider"
)

// version and commit are stamped at link time by the `Build Terraform
// Provider` pipeline step via
// -ldflags "-X main.version=$VERSION -X main.commit=$BITBUCKET_COMMIT".
// They keep their defaults for local `go run` / `go build` invocations.
var (
	version = "dev"
	commit  = ""
)

func main() {
	v := version
	if commit != "" {
		v = version + "+" + commit
	}

	err := providerserver.Serve(context.Background(), provider.New(v), providerserver.ServeOpts{
		Address: "registry.terraform.io/scalegridadmin/scalegrid",
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}
