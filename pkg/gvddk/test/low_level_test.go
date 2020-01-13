package main

import (
	"github.com/vmware-tanzu/astrolabe/pkg/gvddk/gdisklib"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	// Set up
	path := os.Getenv("LIBPATH")
	if path == "" {
		t.Skip("Skipping testing if environment variables are not set.")
	}
	res := gdisklib.Init(6, 7, path)
	if res != nil {
		t.Errorf("Init failed, got error code: %d, error message: %s.", res.VixErrorCode(), res.Error())
	}
	serverName := os.Getenv("IP")
	thumPrint := os.Getenv("THUMBPRINT")
	userName := os.Getenv("USERNAME")
	password := os.Getenv("PWD")
	fcdId := os.Getenv("FCDID")
	ds := os.Getenv("DATASTORE")
	identity := os.Getenv("IDENTITY")
	params := gdisklib.NewConnectParams("", serverName,thumPrint, userName,
		password, fcdId, ds, "", "", identity, "", gdisklib.VIXDISKLIB_FLAG_OPEN_COMPRESSION_SKIPZ,
		false, gdisklib.NBD)
	err1 := gdisklib.PrepareForAccess(params)
	if err1 != nil {
		t.Errorf("Prepare for access failed. Error code: %d. Error message: %s.", err1.VixErrorCode(), err1.Error())
	}
	conn, err2 := gdisklib.ConnectEx(params)
	if err2 != nil {
		gdisklib.EndAccess(params)
		t.Errorf("Connect to vixdisk lib failed. Error code: %d. Error message: %s.", err2.VixErrorCode(), err2.Error())
	}
	gdisklib.Disconnect(conn)
	gdisklib.EndAccess(params)
}
