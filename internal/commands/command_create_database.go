package commands

import "go-zdb-api/pkg/file"

func commandCreateDatabase(commands []string) {

	databaseName := commands[2]

	setPathDatabase(databaseName)

	file.CreateDir(getPathDatabase())
	file.CreateDir(getPathTables())
	file.CreateDir(getPathSequences())
	file.CreateDir(getPathIndexes())

}