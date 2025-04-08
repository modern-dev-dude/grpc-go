package renderclient

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"rendering-engine/packages/renderer"
)

func StartClient() {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.GET("/", renderPageHandler)

	s := http.Server{
		Addr:    ":8080",
		Handler: e,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func renderPageHandler(c echo.Context) error {
	c.Logger().Info("renderPageHandler")
	grpcConn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.Logger().Errorf("grpc client init err: %v", err)
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
		c.Logger().Errorf("grpc client render err: %v", err)
	}

	if err := c.HTML(http.StatusOK, res.Markup); err != nil {
		c.Logger().Errorf("rendering err: %v", err)
	}

	return nil
}
