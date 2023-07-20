package main

import (
	"github/adekang/gin-demo/controller"
	"github/adekang/gin-demo/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	r.POST("/api/img", controller.ImgPost)
	r.GET("/RawMaterial/screenMat", controller.ScreenMaterialAttr)

	// AlphaBeta 管理
	alphaBeta := r.Group("/AlphaBeta")
	{
		alphaBeta.GET("/findAll", controller.FindAllAlphaBeta)
		alphaBeta.POST("/update", controller.UpdateAlphaBeta)
	}

	//指标管理
	targetStatic := r.Group("/TargetStatic")
	{
		targetStatic.GET("/findAll", controller.FindAllTargetStatic)
		targetStatic.POST("/update", controller.UpdateTarget)
	}

	//指标管理
	formSystem := r.Group("/FormSystem")
	{
		formSystem.GET("/findAll", controller.FindAllFormSystem)
		formSystem.POST("/update", controller.UpdateTarget)
	}
	//指标管理
	materialAttr := r.Group("/MaterialAttr")
	{
		materialAttr.GET("/findAll", controller.FindAllMaterialAttr)
		materialAttr.POST("/update", controller.UpdateMaterialAttr)
		materialAttr.POST("/delete", controller.DeleteMaterialAttr)
	}

	//指标管理
	expression := r.Group("/Expression")
	{
		expression.GET("/findAllByUse", controller.FindAllByUse)
	}

	// 指标
	target := r.Group("/Target")
	{
		target.GET("/findStaticTarget", controller.FindStaticByType)
	}

	return r
}
