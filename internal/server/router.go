package server

import (
	"comet/middlewares"
	"gateway/config"
	"gateway/internal/client"
	"gateway/internal/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

func NewRouter(trace trace.Tracer, cfg *config.Config, p *client.ServiceProducer, responseHash *client.MessageHash) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-type"},
	}))

	v1 := router.Group("V0.0.0")
	{
		v1.Use(middlewares.JSONMiddleware())
		v1.Use(middlewares.ResponseMiddleware(cfg.PrivateKeyLocation))

		controller := controllers.NewControllers(trace, cfg, p, responseHash, client.NewAccountClient(cfg))

		/***************** ACCOUNT SERVICE PATH`S **********************/
		accountsGroup := v1.Group("accounts")
		{
			accountsGroup.GET("/health", controller.Status)
			accountsGroup.POST("/refresh-token", controller.RefreshToken)
		}
		authGroup := accountsGroup.Group("auth")
		{
			authGroup.POST("login", controller.LoginUser)
		}
		registerGroup := accountsGroup.Group("register")
		{
			registerGroup.POST("/user", controller.RegisterUser)
		}
		profileGroup := accountsGroup.Group("profile")
		{
			profileGroup.GET("", controller.GetAccountInfo)
			profileGroup.GET("/all", controller.GetAccountsInfo)
		}
		departmentGroup := accountsGroup.Group("department")
		{
			departmentGroup.POST("", controller.AddUserDepartment)
			departmentGroup.GET("/id", controller.GetUserDepartments)
			departmentGroup.DELETE("/id", controller.RemoveUserDepartment)
			departmentGroup.GET("/all", controller.GetUserDepartment)
		}
		roleGroup := accountsGroup.Group("role")
		{
			roleGroup.POST("", controller.AddUserRole)
			roleGroup.GET("/id", controller.GetUserRoles)
			roleGroup.DELETE("/id", controller.RemoveUserRole)
			roleGroup.GET("/all", controller.GetUserRole)
		}

		/********************** MAPS SERVICE PATH`S ****************************/
		mapsGroup := v1.Group("maps")
		{
			mapsGroup.GET("/health", controller.Status)
		}
		layerGroup := mapsGroup.Group("layer")
		{
			layerGroup.POST("/", controller.AddLayer)
			layerGroup.GET("/", controller.Layer)
			layerGroup.PUT("/", controller.EditLayer)
			layerGroup.DELETE("/", controller.DeleteLayer)
			layerGroup.GET("/list", controller.Layers)
		}
		groupGroup := mapsGroup.Group("group")
		{
			groupGroup.GET("/", controller.Group)
			groupGroup.GET("/list", controller.Groups)
			groupGroup.POST("/", controller.AddGroup)
			groupGroup.PUT("/", controller.EditGroup)
			groupGroup.DELETE("/", controller.DeleteGroup)
		}
		mapGroup := mapsGroup.Group("map")
		{
			mapGroup.GET("/", controller.Map)
			mapGroup.GET("/list", controller.Maps)
			mapGroup.POST("/", controller.AddMap)
			mapGroup.PUT("/", controller.EditMap)
			mapGroup.DELETE("/", controller.DeleteMap)
		}
		styleGroup := mapsGroup.Group("style")
		{
			styleGroup.GET("/", controller.Style)
			styleGroup.GET("/list/pagination", controller.StylesPagination)
			styleGroup.GET("/list", controller.Styles)
			styleGroup.POST("/", controller.AddStyle)
			styleGroup.PUT("/", controller.EditStyle)
			styleGroup.DELETE("/", controller.DeleteStyle)
		}
		glrGroup := mapsGroup.Group("glr")
		{
			glrGroup.GET("/list", controller.GroupLayerRelations)
			glrGroup.POST("/", controller.AddGroupLayerRelation)
			glrGroup.DELETE("/", controller.DeleteGroupLayerRelation)
			glrGroup.GET("/grl", controller.GroupRelationLayers)
			glrGroup.GET("/lrg", controller.LayerRelationGroups)
			glrGroup.POST("/up", controller.GroupLayerOrderUp)
			glrGroup.POST("/down", controller.GroupLayerOrderDown)
		}
		mgrGroup := mapsGroup.Group("mgr")
		{
			mgrGroup.GET("/list", controller.MapGroupRelations)
			mgrGroup.POST("/", controller.AddMapGroupRelation)
			mgrGroup.DELETE("/", controller.DeleteMapGroupRelation)
			mgrGroup.GET("/mrg", controller.MapRelationGroups)
			mgrGroup.GET("/grm", controller.GroupRelationMaps)
			mgrGroup.POST("/up", controller.MapGroupOrderUp)
			mgrGroup.POST("/down", controller.MapGroupOrderDown)
		}
		lsrGroup := mapsGroup.Group("lsr")
		{
			lsrGroup.GET("/list", controller.LayerStyleRelations)
			lsrGroup.POST("/", controller.AddLayerStyleRelation)
			lsrGroup.DELETE("/", controller.DeleteLayerStyleRelation)
			lsrGroup.GET("/lrs", controller.LayerRelationStyles)
		}
		styledMapGroup := mapsGroup.Group("styled")
		{
			styledMapGroup.GET("/", controller.StyledMap)
		}
		patternGroup := mapsGroup.Group("pattern")
		{
			patternGroup.GET("/", controller.Pattern)
			patternGroup.GET("/sprite:sep", controller.Sprite)
		}
		tableGroup := mapsGroup.Group("table")
		{
			tableGroup.GET("/list", controller.Tables)
			tableGroup.GET("/", controller.Table)
			tableGroup.PUT("/", controller.EditTable)
			tableGroup.POST("/", controller.AddTable)
			tableGroup.DELETE("/", controller.DeleteTable)
			tableGroup.GET("/columns", controller.TableColumns)
			tableGroup.GET("/column/unique", controller.TableColumnUniqueValues)
			tableGroup.GET("/feature", controller.GetFeatures)
		}
	}

	return router
}
