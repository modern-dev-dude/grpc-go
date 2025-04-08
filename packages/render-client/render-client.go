package renderclient

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"rendering-engine/packages/renderer"
)

func StartClient() {
	e := echo.New()
	e.Use(middleware.RequestID())

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	staticAssetPath := filepath.Join(cwd, "packages", "render-client", "static")
	e.Static("/", staticAssetPath)
	e.GET("/", renderPageHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

func renderPageHandler(c echo.Context) error {
	c.Logger().Info("renderPageHandler")
	grpcConn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		errorHandler(c, err)
		return err
	}
	defer grpcConn.Close()

	grpcClient := renderer.NewRenderingEngineClient(grpcConn)
	msg := renderer.ReqMessage{
		Data: "hello world",
		Metadata: &renderer.Metadata{
			ReqId: c.Request().Header.Get(echo.HeaderXRequestID),
		},
	}
	res, err := grpcClient.RenderPage(context.Background(), &msg)
	if err != nil {
		errorHandler(c, err)
	}

	if err := c.HTML(http.StatusOK, res.Markup); err != nil {
		errorHandler(c, err)
	}

	return nil
}

func errorHandler(c echo.Context, err error) {
	c.Logger().Errorf("rendering err: %v", err)
	c.HTML(http.StatusInternalServerError, "Error handling request")
}
