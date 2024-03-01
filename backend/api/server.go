package api

import (
	"fmt"

	"backend/ent"
	"backend/middleware"
	"backend/token"
	"backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     utils.Config
	client     *ent.Client
	tokenMaker token.Maker
	engine     *gin.Engine
}

func NewServer(config utils.Config, client *ent.Client) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		client:     client,
		tokenMaker: tokenMaker,
	}

	setUpValidator()

	server.setupRouetr()

	return server, nil
}

func setUpValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validRoleName", validRoleName)
		v.RegisterValidation("validEmail", validEmail)
		v.RegisterValidation("validPwd", validPwd)
		v.RegisterValidation("validPostTitle", validPostTitle)
		v.RegisterValidation("validPostBody", validPostBody)
		v.RegisterValidation("validSectionName", validSectionName)
		v.RegisterValidation("validStatement", validStatment)
	}
}

func (server *Server) setupRouetr() {
	router := gin.Default()

	allowCSOR := router.Use(middleware.CrosHandler(server.config))
	allowCSOR.POST("/api/register", server.registerUser)
	allowCSOR.POST("/api/login", server.loginUser)
	allowCSOR.GET("/api/parseCookie", server.parseCookie)
	allowCSOR.GET("/api/role/basis", server.getRoleBasis)
	allowCSOR.GET("/api/role/posts", server.getRolePosts)
	allowCSOR.GET("/api/section/only", server.getSection)
	allowCSOR.GET("/api/section/suggest", server.suggestSection)

	afterSessionAuth := allowCSOR.Use(middleware.SessionAuth(server.config, server.client))
	afterSessionAuth.POST("/api/post", server.createPost)

	afterAdminAuth := afterSessionAuth.Use(middleware.AdminAuth(server.config, server.client))
	afterAdminAuth.POST("/api/role/type", server.updateRoleType)
	afterAdminAuth.POST("/api/section", server.createSection)

	server.engine = router
}

func (server *Server) Start(address string) error {
	return server.engine.Run(address)
}
