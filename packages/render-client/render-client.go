package renderclient

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"path/filepath"
	rn "rendering-engine/packages/random-number"
	"rendering-engine/packages/renderer"
)

func StartClient() {
	e := echo.New()
	e.Use(addRequestIdMiddleware)
	e.Use(middleware.RequestID())

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	staticAssetPath := filepath.Join(cwd, "packages", "render-client", "static")
	e.Static("/", staticAssetPath)
	e.GET("/", renderPageHandler)
	e.GET("/random-number", getRandomNumberHandler)

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

	metadata := &renderer.Metadata{
		ReqId: c.Request().Header.Get(echo.HeaderXRequestID),
	}

	msg := renderer.ReqMessage{
		Data:     "hello world",
		Metadata: metadata,
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
	if err := c.HTML(http.StatusInternalServerError, "Error handling request"); err != nil {
		log.Fatalf("error handling request %v\n", err)
	}
}

func addRequestIdMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	guid, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}

	return func(c echo.Context) error {
		c.Request().Header.Set(echo.HeaderXRequestID, guid.String())
		return next(c)
	}
}

func getRandomNumberHandler(c echo.Context) error {
	c.Logger().Info("getRandomNumberHandler")
	grpcConn, err := grpc.NewClient(":9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		errorHandler(c, err)
		return err
	}
	defer grpcConn.Close()

	rnClient := rn.NewRandomNumberClient(grpcConn)
	msg := &rn.ReqMessage{
		ReqId: c.Request().Header.Get(echo.HeaderXRequestID),
	}

	data, err := rnClient.GetRandomNumber(context.Background(), msg)
	if err != nil {
		errorHandler(c, err)
	}

	if err := c.JSON(http.StatusOK, data); err != nil {
		errorHandler(c, err)
	}

	return nil
}
