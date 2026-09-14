package main

import "embed"

//go:embed migrations/*.up.sql
var migrationFiles embed.FS
