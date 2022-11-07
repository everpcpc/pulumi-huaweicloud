// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sfs

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "huaweicloud:Sfs/accessRule:AccessRule":
		r = &AccessRule{}
	case "huaweicloud:Sfs/fileSystem:FileSystem":
		r = &FileSystem{}
	case "huaweicloud:Sfs/turbo:Turbo":
		r = &Turbo{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := huaweicloud.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Sfs/accessRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Sfs/fileSystem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"huaweicloud",
		"Sfs/turbo",
		&module{version},
	)
}