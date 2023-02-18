// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/salukikit/rodentity/ent/schema","Package":"github.com/salukikit/rodentity/ent","Schemas":[{"name":"Device","config":{"Table":""},"edges":[{"name":"users","type":"User"},{"name":"rodents","type":"Rodent"},{"name":"groups","type":"Group","ref_name":"devices","inverse":true},{"name":"domain","type":"Domain","ref_name":"devices","unique":true,"inverse":true},{"name":"subnets","type":"Subnet","ref_name":"hosts","inverse":true},{"name":"services","type":"Services","ref_name":"devices","inverse":true}],"fields":[{"name":"hostname","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"os","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"unknown","default_kind":24,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"arch","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"unknown","default_kind":24,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"version","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"unknown","default_kind":24,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"net_interfaces","type":{"Type":3,"Ident":"[]string","PkgPath":"","PkgName":"","Nillable":true,"RType":{"Name":"","Ident":"[]string","Kind":23,"PkgPath":"","Methods":{}}},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"machinepass","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"certificates","type":{"Type":3,"Ident":"[]string","PkgPath":"","PkgName":"","Nillable":true,"RType":{"Name":"","Ident":"[]string","Kind":23,"PkgPath":"","Methods":{}}},"optional":true,"position":{"Index":6,"MixedIn":false,"MixinIndex":0}}]},{"name":"Domain","config":{"Table":""},"edges":[{"name":"devices","type":"Device"},{"name":"users","type":"User"},{"name":"groups","type":"Group"},{"name":"childdomains","type":"Domain"},{"name":"parentdomain","type":"Domain","ref_name":"childdomains","unique":true,"inverse":true}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"AD","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":true,"default_kind":1,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"owned","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"cloud","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"unknown","default_kind":24,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}]},{"name":"Group","config":{"Table":""},"edges":[{"name":"devices","type":"Device"},{"name":"users","type":"User"},{"name":"domain","type":"Domain","ref_name":"groups","unique":true,"inverse":true}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"unknown","default_kind":24,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"permissions","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"unknown","default_kind":24,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}]},{"name":"Loot","config":{"Table":""},"edges":[{"name":"rodent","type":"Rodent","ref_name":"loot","unique":true,"inverse":true},{"name":"task","type":"Task","ref_name":"loot","unique":true,"inverse":true}],"fields":[{"name":"xid","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"type","type":{"Type":6,"Ident":"loot.Type","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"cred","V":"cred"},{"N":"key","V":"key"},{"N":"cert","V":"cert"},{"N":"enum","V":"enum"},{"N":"objective","V":"objective"},{"N":"pii","V":"pii"},{"N":"other","V":"other"}],"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"location","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"unknown","default_kind":24,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"data","type":{"Type":5,"Ident":"","PkgPath":"","PkgName":"","Nillable":true,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"collectedon","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"position":{"Index":4,"MixedIn":false,"MixinIndex":0}}]},{"name":"Operator","config":{"Table":""},"edges":[{"name":"projects","type":"Project","ref_name":"operators","inverse":true},{"name":"tasks","type":"Task"}],"fields":[{"name":"username","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"privkey","type":{"Type":5,"Ident":"","PkgPath":"","PkgName":"","Nillable":true,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"cert","type":{"Type":5,"Ident":"","PkgPath":"","PkgName":"","Nillable":true,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}]},{"name":"Project","config":{"Table":""},"edges":[{"name":"operators","type":"Operator"},{"name":"rodents","type":"Rodent"},{"name":"routers","type":"Router"}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"objective","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"end_date","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"start_date","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}]},{"name":"Rodent","config":{"Table":""},"edges":[{"name":"device","type":"Device","ref_name":"rodents","unique":true,"inverse":true},{"name":"user","type":"User","ref_name":"rodents","unique":true,"inverse":true},{"name":"project","type":"Project","ref_name":"rodents","unique":true,"inverse":true},{"name":"router","type":"Router","ref_name":"rodents","inverse":true},{"name":"tasks","type":"Task"},{"name":"loot","type":"Loot"}],"fields":[{"name":"xid","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"type","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"FancyRat","default_kind":24,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"codename","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"key","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"usercontext","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"comms","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"comms_inspected","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":6,"MixedIn":false,"MixinIndex":0}},{"name":"beacontime","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":7,"MixedIn":false,"MixinIndex":0}},{"name":"burned","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":8,"MixedIn":false,"MixinIndex":0}},{"name":"alive","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":true,"default_kind":1,"position":{"Index":9,"MixedIn":false,"MixinIndex":0}},{"name":"joined","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"position":{"Index":10,"MixedIn":false,"MixinIndex":0}},{"name":"lastseen","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"position":{"Index":11,"MixedIn":false,"MixinIndex":0}}]},{"name":"Router","config":{"Table":""},"edges":[{"name":"rodents","type":"Rodent"},{"name":"project","type":"Project","ref_name":"routers","unique":true,"inverse":true}],"fields":[{"name":"username","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"privkey","type":{"Type":5,"Ident":"","PkgPath":"","PkgName":"","Nillable":true,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"cert","type":{"Type":5,"Ident":"","PkgPath":"","PkgName":"","Nillable":true,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"commands","type":{"Type":3,"Ident":"[]string","PkgPath":"","PkgName":"","Nillable":true,"RType":{"Name":"","Ident":"[]string","Kind":23,"PkgPath":"","Methods":{}}},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}]},{"name":"Services","config":{"Table":""},"edges":[{"name":"devices","type":"Device"},{"name":"subnet","type":"Subnet"}],"fields":[{"name":"service_name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"port","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}}]},{"name":"Subnet","config":{"Table":""},"edges":[{"name":"hosts","type":"Device"}],"fields":[{"name":"cidr","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"mask","type":{"Type":5,"Ident":"","PkgPath":"","PkgName":"","Nillable":true,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"outbound_tcpports","type":{"Type":3,"Ident":"[]string","PkgPath":"","PkgName":"","Nillable":true,"RType":{"Name":"","Ident":"[]string","Kind":23,"PkgPath":"","Methods":{}}},"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"outbound_udpports","type":{"Type":3,"Ident":"[]string","PkgPath":"","PkgName":"","Nillable":true,"RType":{"Name":"","Ident":"[]string","Kind":23,"PkgPath":"","Methods":{}}},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"proxy","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}}]},{"name":"Task","config":{"Table":""},"edges":[{"name":"rodent","type":"Rodent","ref_name":"tasks","unique":true,"inverse":true},{"name":"operator","type":"Operator","ref_name":"tasks","unique":true,"inverse":true},{"name":"loot","type":"Loot"}],"fields":[{"name":"xid","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"type","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"cmd","default_kind":24,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"args","type":{"Type":3,"Ident":"[]string","PkgPath":"","PkgName":"","Nillable":true,"RType":{"Name":"","Ident":"[]string","Kind":23,"PkgPath":"","Methods":{}}},"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"data","type":{"Type":5,"Ident":"","PkgPath":"","PkgName":"","Nillable":true,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"result","type":{"Type":5,"Ident":"","PkgPath":"","PkgName":"","Nillable":true,"RType":null},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"Executed","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"looted","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":6,"MixedIn":false,"MixinIndex":0}},{"name":"requestedat","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"position":{"Index":7,"MixedIn":false,"MixinIndex":0}},{"name":"completedat","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":8,"MixedIn":false,"MixinIndex":0}},{"name":"TTPs","type":{"Type":3,"Ident":"[]string","PkgPath":"","PkgName":"","Nillable":true,"RType":{"Name":"","Ident":"[]string","Kind":23,"PkgPath":"","Methods":{}}},"optional":true,"position":{"Index":9,"MixedIn":false,"MixinIndex":0}}]},{"name":"User","config":{"Table":""},"edges":[{"name":"devices","type":"Device","ref_name":"users","inverse":true},{"name":"rodents","type":"Rodent"},{"name":"groups","type":"Group","ref_name":"users","inverse":true},{"name":"domain","type":"Domain","ref_name":"users","unique":true,"inverse":true}],"fields":[{"name":"username","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"givenname","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"unknown","default_kind":24,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"email","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"unknown","default_kind":24,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"owned","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"age","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":"unknown","default_kind":24,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"homedir","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"unknown","default_kind":24,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"enabled","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":true,"default_kind":1,"position":{"Index":6,"MixedIn":false,"MixinIndex":0}}]}],"Features":["schema/snapshot"]}`
