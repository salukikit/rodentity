# rodentity

**WIP PACKAGE, YOU HAVE BEEN WARNED**

rodentity is the [ent](https://entgo.io) schema package for rodents. This package can be used as a template schema for rodent implants and databases.
This is intended to be used as either a drop in package for the not quite released - [RodentFramework](https://github.com/salukikit/rodentframework).

However, this can also be used for your own custom C2, or just as a schema for a red team style tool. The schema can be taken from `ent/schema` and modified it to suit your needs.

## Why entity/ORM for a C2?

1. Simply put Ent is a really neat package that I wanted to use. It supports multiple databases, allowing for operator choice between sqlite (inc. an In-Memory DB), mysql and postgres. You can also use this schema with other DBs such as graphing dbs! 

2. I wanted to ensure rodent had the flexibility, and structure, to map relationships between many of the objects you may encounter during an engagement. Ideally giving a detailed dataset which can easily trace back the origins of loot, the task TTPs and the details around a user compromise, e.g. their group memberships, which devices we've seen them on and which tasks relate to them.

3. Rodentity is just the beginning, and is only intended to be a single component between numerous tools I am making. However, it's the core that backs much of my ideas, which range from enumeration all the way to reporting.