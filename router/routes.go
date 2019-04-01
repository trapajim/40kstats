package router

import (
	"github.com/trapajim/rest/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc handler.HandlerFunc
}

type Routes []Route

var GetRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/dashboard",
		handler.Index,
	},
	Route{
		"Compares two lists",
		"GET",
		"/compare/{ids}",
		handler.Comparison,
	},
	Route{
		"Factions",
		"GET",
		"/factions",
		handler.Factions,
	},
	// lists
	Route{
		"AddNewList",
		"Post",
		"/list",
		handler.AddNewList,
	},
	Route{
		"Update List",
		"Put",
		"/list/{id}",
		handler.UpdateList,
	},
	Route{
		"Delete List",
		"DELETE",
		"/list/{id}",
		handler.DeleteList,
	},
	Route{
		"Show Lists",
		"Get",
		"/show-list",
		handler.ShowLists,
	},
	// battle reports
	Route{
		"Add new battle report",
		"Post",
		"/battlereport",
		handler.AddReport,
	},
	Route{
		"List Battle reports",
		"Get",
		"/battlereport",
		handler.ListReports,
	},
	Route{
		"Delete a battle report",
		"DELETE",
		"/battlereport/{id}",
		handler.DeleteBattleReport,
	},
	Route{
		"Update a battle report",
		"Put",
		"/battlereport/{id}",
		handler.UpdateBattleReport,
	},
	Route{
		"Show a battle report",
		"GET",
		"/battlereport/{id}",
		handler.ShowBattleReport,
	},
}
