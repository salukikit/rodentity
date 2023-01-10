# rodentity

**WIP PACKAGE, YOU HAVE BEEN WARNED**

rodentity is the [ent](https://entgo.io) schema package for rodents. This package can be used as a template schema for rodent implants and databases.
This is intended to be used as either a drop in package for the upcoming [RodentFramework](https://github.com/salukikit/rodentframework), your own custom C2 even as a starting point for your own C2, using the schema under `ent/schema` and modifying it to suit your needs.

## Why entity/ORM for a C2?

Simply put Ent is a really neat package that I wanted to use. It supports multiple databases, allowing for operator choice between sqlite (in memory only), mysql and postgres. You can also use this schema with other DBs such as graphing dbs! I also wanted to ensure rodent had the flexibility, and structure, to map relationships between many of the objects you may encounter during an engagement. Ideally giving a detailed dataset which can easily trace back the origins of loot, the task TTPs and the details around a user compromise, e.g. their group memberships, which devices we've seen them on and which tasks relate to them.
