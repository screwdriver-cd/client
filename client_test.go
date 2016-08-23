package main

import (
	"testing"
)

func TestCreateApp(t *testing.T) {
	app := CreateApp()
	if app == nil {
		t.Error("Expected app to be created")
	}
	appName := "Screwdriver Client"
	if app.Name != appName {
		t.Errorf("app.Name = %s, want %s", app.Name, appName)
	}
	appUsage := "Continuous Delivery With Screwdriver"
	if app.Usage != appUsage {
		t.Errorf("app.Name = %s, want %s", app.Usage, appUsage)
	}

}
