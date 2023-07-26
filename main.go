package main

import (
	"fmt"
	"k8s_test/global"
	"k8s_test/internal/model"
	"k8s_test/pkg/logger"
	"k8s_test/pkg/setting"
	"log"
	"time"

	"k8s_test/internal/controller"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./pkg/static")
	router.LoadHTMLGlob("pkg/templates/*")
	// todo 参数校验
	mainC := controller.NewMainController()
	router.GET("/index", mainC.GetIndex)
	deployC := controller.NewDeploymentController()
	router.POST("/deployment", deployC.CreateDeployment)
	router.GET("/deployment", deployC.ListDeployment)
	router.DELETE("/deployment", deployC.DeleteDeployment)
	router.Run(":" + global.ServerSetting.HttpPort)
}

func init() {
	err := setupSettings()
	if err != nil {
		log.Fatalf("init.setupSetting failed.err:%v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger failed.err: %v", err)
	}
	err = setupKubernetes()
	if err != nil {
		log.Fatalf("init setupKubernetes failed.err:%v", err)
	}
}

func setupSettings() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Kubernete", &global.KuberneteSetting)
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	fmt.Println(global.ServerSetting, global.AppSetting, global.KuberneteSetting)
	return nil
}

func setupLogger() error {
	// todo if LogSavePath not exist should create
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupKubernetes() error {
	var err error
	global.KubeConfig, err = clientcmd.BuildConfigFromFlags("", global.KuberneteSetting.KubeConfig)
	if err != nil {
		return fmt.Errorf("client cmd build config from flags failed.err:%v", err)
	}
	return nil
}
