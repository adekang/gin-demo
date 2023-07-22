package main

import (
	"github/adekang/gin-demo/controller"
	"github/adekang/gin-demo/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	v1 := r.Group("/api/v1")
	{

		v1.POST("/auth/register", controller.Register)
		v1.POST("/auth/login", controller.Login)
		v1.POST("/auth/info", middleware.AuthMiddleware(), controller.Info)
		v1.POST("/img", controller.ImgPost)
		v1.GET("/RawMaterial/screenMat", controller.ScreenMaterialAttr)

		// AlphaBeta 管理
		alphaBeta := v1.Group("/AlphaBeta")
		{
			alphaBeta.GET("/findAll", controller.FindAllAlphaBeta)
			alphaBeta.POST("/update", controller.UpdateAlphaBeta)
		}

		//指标管理
		targetStatic := v1.Group("/TargetStatic")
		{
			targetStatic.GET("/findAll", controller.FindAllTargetStatic)
			targetStatic.POST("/update", controller.UpdateTarget)
		}

		//指标管理
		formSystem := v1.Group("/FormSystem")
		{
			formSystem.GET("/findAll", controller.FindAllFormSystem)
			formSystem.POST("/update", controller.UpdateTarget)
		}
		//指标管理
		materialAttr := v1.Group("/MaterialAttr")
		{
			materialAttr.GET("/findAll", controller.FindAllMaterialAttr)
			materialAttr.POST("/update", controller.UpdateMaterialAttr)
			materialAttr.POST("/delete", controller.DeleteMaterialAttr)
		}

		//指标管理
		expression := v1.Group("/Expression")
		{
			expression.GET("/findAllByUse", controller.FindAllByUse)
		}

		// 指标
		target := v1.Group("/Target")
		{
			target.GET("/findStaticTarget", controller.FindStaticByType)
		}
	}

	return r
}
