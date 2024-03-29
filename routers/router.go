// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"compgen-api-docs/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/document",
			beego.NSInclude(
				&controllers.DocumentController{},
			),
		),
		beego.NSNamespace("/permissions",
			beego.NSInclude(
				&controllers.PermissionsController{},
			),
		),
		beego.NSNamespace("/serviceorder",
			beego.NSInclude(
				&controllers.ServiceOrderController{},
			),
		),
		beego.NSNamespace("/services",
			beego.NSInclude(
				&controllers.ServicesController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
