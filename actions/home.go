package actions

import "github.com/gobuffalo/buffalo"

/// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	c.Set("task-title", "Dashboard")
	c.Set("task-icon", "fa fa-dashboard")
	c.Set("task-desc", "A free and open source Bootstrap 4 admin template")
	return c.Render(200, r.HTML("index.html"))
}

// ChartsHandler testing dynamic content handling of vali-admin theme
func ChartsHandler(c buffalo.Context) error {
	c.Set("task-title", "Charts")
	c.Set("task-icon", "fa fa-pie-chart")
	c.Set("task-desc", "Various type of charts for your project")
	return c.Render(200, r.HTML("charts.html"))
}

func ComponentsHandler(c buffalo.Context) error {
	c.Set("task-title", "Bootstrap Elements")
	c.Set("task-icon", "fa fa-laptop")
	c.Set("task-desc", "Bootstrap Elements")
	return c.Render(200, r.HTML("bootstrap-components.html"))
}

func CardsHandler(c buffalo.Context) error {
	c.Set("task-title", " Cards")
	c.Set("task-icon", "fa fa-laptop")
	c.Set("task-desc", "Material design inspired cards")
	return c.Render(200, r.HTML("ui-cards.html"))
}

func WidgetsHandler(c buffalo.Context) error {
	c.Set("task-title", " Widgets")
	c.Set("task-icon", "fa fa-archive")
	c.Set("task-desc", "Widgets to interactively display data")
	return c.Render(200, r.HTML("widgets.html"))
}

func FormComponentsHandler(c buffalo.Context) error {
	c.Set("task-title", " Form Components")
	c.Set("task-icon", "fa fa-edit")
	c.Set("task-desc", "Bootstrap default form components")
	return c.Render(200, r.HTML("form-components.html"))
}

func FormCustomHandler(c buffalo.Context) error {
	c.Set("task-title", " Custom Form Elements")
	c.Set("task-icon", "fa fa-edit")
	c.Set("task-desc", "Customized form elements")
	return c.Render(200, r.HTML("form-custom.html"))
}

func FormSamplesHandler(c buffalo.Context) error {
	c.Set("task-title", " Form Samples")
	c.Set("task-icon", "fa fa-edit")
	c.Set("task-desc", "Sample forms")
	return c.Render(200, r.HTML("form-samples.html"))
}

func FormNotificationsHandler(c buffalo.Context) error {
	c.Set("task-title", " Form Notifications")
	c.Set("task-icon", "fa fa-bell-o")
	c.Set("task-desc", "Jquery Plugins to notify user status about some action.")
	return c.Render(200, r.HTML("form-notifications.html"))
}

func TableBasicHandler(c buffalo.Context) error {
	c.Set("task-title", " Basic Tables")
	c.Set("task-icon", "fa fa-th-list")
	c.Set("task-desc", "Basic bootstrap tables")
	return c.Render(200, r.HTML("table-basic.html"))
}

func DataTableHandler(c buffalo.Context) error {
	c.Set("task-title", " Data Table")
	c.Set("task-icon", "fa fa-th-list")
	c.Set("task-desc", "Table to display analytical data effectively")
	return c.Render(200, r.HTML("table-data-table.html"))
}

func BlankPageHandler(c buffalo.Context) error {
	c.Set("task-title", " Blank Page")
	c.Set("task-icon", "fa fa-dashboard")
	c.Set("task-desc", "Start a beautiful journey here")
	return c.Render(200, r.HTML("blank-page.html"))
}

func PageLoginHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("page-login.html", "layout-no-nav.html"))
}

func LockscreenHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("page-lockscreen.html", "layout-no-nav.html"))
}

func PageUserHandler(c buffalo.Context) error {
	c.Set("task-title", " User Profile")
	c.Set("task-icon", "fa fa-user")
	c.Set("task-desc", "User Settings")
	return c.Render(200, r.HTML("page-user.html"))
}

func PageInvoiceHandler(c buffalo.Context) error {
	c.Set("task-title", " Invoice")
	c.Set("task-icon", "fa fa-file-text-o")
	c.Set("task-desc", "A Printable invoice Format")
	return c.Render(200, r.HTML("page-invoice.html"))
}

func PageCalendarHandler(c buffalo.Context) error {
	c.Set("task-title", " Calendar")
	c.Set("task-icon", "fa fa-calendar")
	c.Set("task-desc", "Full Calander page for managing events")
	return c.Render(200, r.HTML("page-calendar.html"))
}

func PageMailboxHandler(c buffalo.Context) error {
	c.Set("task-title", " Mailbox")
	c.Set("task-icon", "fa fa-envelope-o")
	c.Set("task-desc", "A Mailbox page sample")
	return c.Render(200, r.HTML("page-mailbox.html"))
}

func PageErrorHandler(c buffalo.Context) error {
	c.Set("task-title", " Error")
	c.Set("task-icon", "fa ")
	c.Set("task-desc", "")

	return c.Render(200, r.HTML("page-error.html"))
}
