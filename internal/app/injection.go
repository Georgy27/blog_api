package app

import (
	"context"
	"github.com/Georgy27/blogger_api/internal/blog/handler"
	blogRepository "github.com/Georgy27/blogger_api/internal/blog/repository"
	blogService "github.com/Georgy27/blogger_api/internal/blog/service"
	"github.com/Georgy27/blogger_api/internal/config"
	"github.com/Georgy27/blogger_api/pkg/closer"
	"github.com/Georgy27/blogger_api/pkg/db"
	"github.com/Georgy27/blogger_api/pkg/db/pg"
	"github.com/Georgy27/blogger_api/pkg/db/transaction"
	"log"
)

type injector struct {
	db             db.Client
	txManager      db.TxManager
	blogRepository blogRepository.BlogRepository
	blogService    blogService.BlogService
	blogHandler    *handler.BlogHandler

	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig
}

func newInjector() *injector {
	return &injector{}
}

func (i *injector) GetPGConfig() config.PGConfig {
	if i.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		i.pgConfig = cfg
	}

	return i.pgConfig
}

func (i *injector) GRPCConfig() config.GRPCConfig {
	if i.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		i.grpcConfig = cfg
	}

	return i.grpcConfig
}

func (i *injector) HTTPConfig() config.HTTPConfig {
	if i.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		i.httpConfig = cfg
	}

	return i.httpConfig
}

func (i *injector) SwaggerConfig() config.SwaggerConfig {
	if i.swaggerConfig == nil {
		cfg, err := config.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		i.swaggerConfig = cfg
	}

	return i.swaggerConfig
}

func (i *injector) GetDB(ctx context.Context) db.Client {
	if i.db == nil {
		cl, err := pg.New(ctx, i.GetPGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %s", err.Error())
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		i.db = cl
	}

	return i.db
}

func (i *injector) GetTxManager(ctx context.Context) db.TxManager {
	if i.txManager == nil {
		i.txManager = transaction.NewTransactionManager(i.GetDB(ctx).DB())
	}

	return i.txManager
}

func (i *injector) GetNoteRepository(ctx context.Context) blogRepository.BlogRepository {
	if i.blogRepository == nil {
		i.blogRepository = blogRepository.NewBlogRepository(i.GetDB(ctx))
	}

	return i.blogRepository
}

func (i *injector) GetBlogService(ctx context.Context) blogService.BlogService {
	if i.blogService == nil {
		i.blogService = blogService.NewBlogService(
			i.GetNoteRepository(ctx),
		)
	}

	return i.blogService
}

func (i *injector) GetBlogHandler(ctx context.Context) *handler.BlogHandler {
	if i.blogHandler == nil {
		i.blogHandler = handler.NewBlogHandler(
			i.GetBlogService(ctx),
		)
	}

	return i.blogHandler
}
