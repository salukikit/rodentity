// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/salukikit/rodentity/ent/device"
	"github.com/salukikit/rodentity/ent/domain"
	"github.com/salukikit/rodentity/ent/group"
	"github.com/salukikit/rodentity/ent/loot"
	"github.com/salukikit/rodentity/ent/rodent"
	"github.com/salukikit/rodentity/ent/schema"
	"github.com/salukikit/rodentity/ent/task"
	"github.com/salukikit/rodentity/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	deviceFields := schema.Device{}.Fields()
	_ = deviceFields
	// deviceDescOs is the schema descriptor for os field.
	deviceDescOs := deviceFields[1].Descriptor()
	// device.DefaultOs holds the default value on creation for the os field.
	device.DefaultOs = deviceDescOs.Default.(string)
	// deviceDescArch is the schema descriptor for arch field.
	deviceDescArch := deviceFields[2].Descriptor()
	// device.DefaultArch holds the default value on creation for the arch field.
	device.DefaultArch = deviceDescArch.Default.(string)
	// deviceDescVersion is the schema descriptor for version field.
	deviceDescVersion := deviceFields[3].Descriptor()
	// device.DefaultVersion holds the default value on creation for the version field.
	device.DefaultVersion = deviceDescVersion.Default.(string)
	// deviceDescLocaladdress is the schema descriptor for localaddress field.
	deviceDescLocaladdress := deviceFields[4].Descriptor()
	// device.DefaultLocaladdress holds the default value on creation for the localaddress field.
	device.DefaultLocaladdress = deviceDescLocaladdress.Default.(string)
	domainFields := schema.Domain{}.Fields()
	_ = domainFields
	// domainDescAD is the schema descriptor for AD field.
	domainDescAD := domainFields[1].Descriptor()
	// domain.DefaultAD holds the default value on creation for the AD field.
	domain.DefaultAD = domainDescAD.Default.(bool)
	// domainDescOwned is the schema descriptor for owned field.
	domainDescOwned := domainFields[2].Descriptor()
	// domain.DefaultOwned holds the default value on creation for the owned field.
	domain.DefaultOwned = domainDescOwned.Default.(bool)
	// domainDescCloud is the schema descriptor for cloud field.
	domainDescCloud := domainFields[3].Descriptor()
	// domain.DefaultCloud holds the default value on creation for the cloud field.
	domain.DefaultCloud = domainDescCloud.Default.(string)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescDescription is the schema descriptor for description field.
	groupDescDescription := groupFields[1].Descriptor()
	// group.DefaultDescription holds the default value on creation for the description field.
	group.DefaultDescription = groupDescDescription.Default.(string)
	// groupDescPermissions is the schema descriptor for permissions field.
	groupDescPermissions := groupFields[2].Descriptor()
	// group.DefaultPermissions holds the default value on creation for the permissions field.
	group.DefaultPermissions = groupDescPermissions.Default.(string)
	lootFields := schema.Loot{}.Fields()
	_ = lootFields
	// lootDescLocation is the schema descriptor for location field.
	lootDescLocation := lootFields[2].Descriptor()
	// loot.DefaultLocation holds the default value on creation for the location field.
	loot.DefaultLocation = lootDescLocation.Default.(string)
	rodentFields := schema.Rodent{}.Fields()
	_ = rodentFields
	// rodentDescType is the schema descriptor for type field.
	rodentDescType := rodentFields[1].Descriptor()
	// rodent.DefaultType holds the default value on creation for the type field.
	rodent.DefaultType = rodentDescType.Default.(string)
	// rodentDescBurned is the schema descriptor for burned field.
	rodentDescBurned := rodentFields[6].Descriptor()
	// rodent.DefaultBurned holds the default value on creation for the burned field.
	rodent.DefaultBurned = rodentDescBurned.Default.(bool)
	// rodentDescAlive is the schema descriptor for alive field.
	rodentDescAlive := rodentFields[7].Descriptor()
	// rodent.DefaultAlive holds the default value on creation for the alive field.
	rodent.DefaultAlive = rodentDescAlive.Default.(bool)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescType is the schema descriptor for type field.
	taskDescType := taskFields[1].Descriptor()
	// task.DefaultType holds the default value on creation for the type field.
	task.DefaultType = taskDescType.Default.(string)
	// taskDescExecuted is the schema descriptor for Executed field.
	taskDescExecuted := taskFields[5].Descriptor()
	// task.DefaultExecuted holds the default value on creation for the Executed field.
	task.DefaultExecuted = taskDescExecuted.Default.(bool)
	// taskDescLooted is the schema descriptor for looted field.
	taskDescLooted := taskFields[6].Descriptor()
	// task.DefaultLooted holds the default value on creation for the looted field.
	task.DefaultLooted = taskDescLooted.Default.(bool)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescGivenname is the schema descriptor for givenname field.
	userDescGivenname := userFields[1].Descriptor()
	// user.DefaultGivenname holds the default value on creation for the givenname field.
	user.DefaultGivenname = userDescGivenname.Default.(string)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.DefaultEmail holds the default value on creation for the email field.
	user.DefaultEmail = userDescEmail.Default.(string)
	// userDescOwned is the schema descriptor for owned field.
	userDescOwned := userFields[3].Descriptor()
	// user.DefaultOwned holds the default value on creation for the owned field.
	user.DefaultOwned = userDescOwned.Default.(bool)
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[4].Descriptor()
	// user.DefaultAge holds the default value on creation for the age field.
	user.DefaultAge = userDescAge.Default.(string)
	// userDescHomedir is the schema descriptor for homedir field.
	userDescHomedir := userFields[5].Descriptor()
	// user.DefaultHomedir holds the default value on creation for the homedir field.
	user.DefaultHomedir = userDescHomedir.Default.(string)
	// userDescEnabled is the schema descriptor for enabled field.
	userDescEnabled := userFields[6].Descriptor()
	// user.DefaultEnabled holds the default value on creation for the enabled field.
	user.DefaultEnabled = userDescEnabled.Default.(bool)
}
