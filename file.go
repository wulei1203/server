package server

import (
	"path"
)

func readFile(conn *Conn, param string) {
	conn.writeMessage(450, "dirver root path :====>>"+conn.driver.GetRootPath())
	fileName := "./" + param
	fileSuffix := path.Ext(fileName)

	if fileSuffix != ".csv" {
		conn.writeMessage(450, fileSuffix+",file type not support!")
	}
}
