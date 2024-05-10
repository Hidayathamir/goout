package app

import (
	"context"
	"net"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/repo/cache"
	"github.com/Hidayathamir/goout/internal/repo/db"
	"github.com/Hidayathamir/goout/internal/repo/db/migration/migrate"
	transportgrpc "github.com/Hidayathamir/goout/internal/transport/grpc"
	"github.com/Hidayathamir/goout/internal/transport/grpc/grpcmiddleware"
	transporthttp "github.com/Hidayathamir/goout/internal/transport/http"
	"github.com/Hidayathamir/goout/internal/transport/http/httpmiddleware"
	"github.com/Hidayathamir/goout/internal/usecase/injection"
	"github.com/Hidayathamir/goout/pkg/errutil"
	pbgocheck "github.com/Hidayathamir/protobuf/gocheck"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Run runs application.
func Run() { //nolint:funlen
	logrus.SetFormatter(&logrus.JSONFormatter{})

	cfg, err := config.Load("./config.yml")
	fatalIfErr(err)

	pg, err := db.NewPostgres(cfg)
	fatalIfErr(err)

	if cfg.GetMigrationAuto() {
		err := migrate.Up(pg.DB)
		if cfg.GetMigrationRequired() {
			fatalIfErr(err)
		} else {
			warnIfErr(err)
		}
	}

	redis, err := cache.NewRedis(cfg)
	fatalIfErr(err)

	gocheckGRPCClientConn, err := grpc.Dial(net.JoinHostPort(cfg.GetExtAPIGocheckGRPCHost(), cfg.GetExtAPIGocheckGRPCPort()), grpc.WithTransportCredentials(insecure.NewCredentials()))
	fatalIfErr(err)
	defer func() {
		err := gocheckGRPCClientConn.Close()
		warnIfErr(err)
	}()

	gocheckgrpcDigitalWalletClient := pbgocheck.NewDigitalWalletClient(gocheckGRPCClientConn)

	usecaseErajolBike := injection.InitUsecaseErajolBike(cfg, pg, redis, gocheckgrpcDigitalWalletClient)
	transportgrpcErajolBike := transportgrpc.NewErajolBike(cfg, usecaseErajolBike)
	transporthttpErajolBike := transporthttp.NewErajolBike(cfg, usecaseErajolBike)

	logrus.Info("initializing grpc server in a goroutine so that it won't block the graceful shutdown handling below")
	var grpcServer *grpc.Server
	go func() {
		grpcServer = grpc.NewServer(grpc.UnaryInterceptor(grpcmiddleware.TraceID))

		registerGRPCServer(
			grpcServer,
			transportgrpcErajolBike,
		)

		addr := net.JoinHostPort(cfg.GetGRPCHost(), cfg.GetGRPCPort())
		lis, err := net.Listen("tcp", addr)
		fatalIfErr(err)

		logrus.WithField("address", addr).Info("grpc server running 🟢")
		err = grpcServer.Serve(lis)
		fatalIfErr(err)
	}()

	logrus.Info("initializing http server in a goroutine so that it won't block the graceful shutdown handling below")
	var httpServer *http.Server
	go func() {
		ginEngine := gin.Default()
		ginEngine.Use(httpmiddleware.TraceID)

		registerHTTPRouter(
			ginEngine,
			transporthttpErajolBike,
		)

		addr := net.JoinHostPort(cfg.GetHTTPHost(), cfg.GetHTTPPort())
		httpServer = &http.Server{ //nolint:gosec
			Addr:    addr,
			Handler: ginEngine,
		}

		logrus.WithField("address", addr).Info("http server running 🟢")
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logrus.Fatal(errutil.Wrap(err))
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logrus.Info("listens for the interrupt signal from the OS")
	<-ctx.Done()

	stop()
	logrus.Info("shutting down gracefully, press Ctrl+C again to force")

	var gracefulShutdownWG sync.WaitGroup

	gracefulShutdownWG.Add(1)
	go func() {
		logrus.Info("shutting down gracefully grpc server")
		defer gracefulShutdownWG.Done()

		grpcServer.GracefulStop()

		logrus.Info("shutting down gracefully grpc server done 🟢")
	}()

	gracefulShutdownWG.Add(1)
	go func() {
		logrus.Info("shutting down gracefully http server")
		defer gracefulShutdownWG.Done()

		logrus.Info("inform http server it has 10 seconds to finish the request it is currently handling")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //nolint:gomnd
		defer cancel()
		err = httpServer.Shutdown(ctx)
		fatalIfErr(err)

		logrus.Info("shutting down gracefully http server done 🟢")
	}()

	logrus.Info("wait graceful shutdown finish")
	gracefulShutdownWG.Wait()
	logrus.Info("graceful shutdown done 🟢")
}

func fatalIfErr(err error) {
	if err != nil {
		logrus.Fatal(errutil.Wrap(err, errutil.WithSkip(1)))
	}
}

func warnIfErr(err error) {
	if err != nil {
		logrus.Warn(errutil.Wrap(err, errutil.WithSkip(1)))
	}
}
