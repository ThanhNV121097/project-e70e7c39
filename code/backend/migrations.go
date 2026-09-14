package backend

import "embed"

//go:embed migrations/*.up.sql
var Migrations embed.FS
