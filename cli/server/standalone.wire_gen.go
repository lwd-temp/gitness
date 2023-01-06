// Code generated by Wire. DO NOT EDIT.

//go:build !wireinject && !harness
// +build !wireinject,!harness

package server

import (
	"context"
	"github.com/harness/gitness/events"
	"github.com/harness/gitness/gitrpc"
	server2 "github.com/harness/gitness/gitrpc/server"
	"github.com/harness/gitness/internal/api/controller/githook"
	"github.com/harness/gitness/internal/api/controller/pullreq"
	"github.com/harness/gitness/internal/api/controller/repo"
	"github.com/harness/gitness/internal/api/controller/serviceaccount"
	"github.com/harness/gitness/internal/api/controller/space"
	"github.com/harness/gitness/internal/api/controller/user"
	webhook2 "github.com/harness/gitness/internal/api/controller/webhook"
	"github.com/harness/gitness/internal/auth/authn"
	"github.com/harness/gitness/internal/auth/authz"
	"github.com/harness/gitness/internal/bootstrap"
	"github.com/harness/gitness/internal/cron"
	events2 "github.com/harness/gitness/internal/events/git"
	"github.com/harness/gitness/internal/router"
	"github.com/harness/gitness/internal/server"
	"github.com/harness/gitness/internal/store"
	"github.com/harness/gitness/internal/store/database"
	"github.com/harness/gitness/internal/url"
	"github.com/harness/gitness/internal/webhook"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/check"
)

// Injectors from standalone.wire.go:

func initSystem(ctx context.Context, config *types.Config) (*system, error) {
	checkUser := check.ProvideUserCheck()
	authorizer := authz.ProvideAuthorizer()
	db, err := database.ProvideDatabase(ctx, config)
	if err != nil {
		return nil, err
	}
	principalUIDTransformation := store.ProvidePrincipalUIDTransformation()
	principalStore := database.ProvidePrincipalStore(db, principalUIDTransformation)
	tokenStore := database.ProvideTokenStore(db)
	controller := user.NewController(checkUser, authorizer, principalStore, tokenStore)
	bootstrapBootstrap := bootstrap.ProvideBootstrap(config, controller)
	authenticator := authn.ProvideAuthenticator(principalStore, tokenStore)
	provider, err := url.ProvideURLProvider(config)
	if err != nil {
		return nil, err
	}
	checkRepo := check.ProvideRepoCheck()
	pathTransformation := store.ProvidePathTransformation()
	spaceStore := database.ProvideSpaceStore(db, pathTransformation)
	repoStore := database.ProvideRepoStore(db, pathTransformation)
	gitrpcConfig := ProvideGitRPCClientConfig(config)
	gitrpcInterface, err := gitrpc.ProvideClient(gitrpcConfig)
	if err != nil {
		return nil, err
	}
	repoController := repo.ProvideController(config, provider, checkRepo, authorizer, spaceStore, repoStore, principalStore, gitrpcInterface)
	checkSpace := check.ProvideSpaceCheck()
	spaceController := space.ProvideController(provider, checkSpace, authorizer, spaceStore, repoStore, principalStore)
	pullReqStore := database.ProvidePullReqStore(db)
	pullReqActivityStore := database.ProvidePullReqActivityStore(db)
	pullReqReviewStore := database.ProvidePullReqReviewStore(db)
	pullReqReviewerStore := database.ProvidePullReqReviewerStore(db)
	pullreqController := pullreq.ProvideController(db, authorizer, pullReqStore, pullReqActivityStore, pullReqReviewStore, pullReqReviewerStore, repoStore, principalStore, gitrpcInterface)
	webhookStore := database.ProvideWebhookStore(db)
	webhookExecutionStore := database.ProvideWebhookExecutionStore(db)
	webhookConfig := ProvideWebhookConfig(config)
	eventsConfig := ProvideEventsConfig(config)
	cmdable, err := ProvideRedis(config)
	if err != nil {
		return nil, err
	}
	eventsSystem, err := events.ProvideSystem(eventsConfig, cmdable)
	if err != nil {
		return nil, err
	}
	readerFactory, err := events2.ProvideReaderFactory(eventsSystem)
	if err != nil {
		return nil, err
	}
	webhookServer, err := webhook.ProvideServer(ctx, webhookConfig, readerFactory, webhookStore, webhookExecutionStore, repoStore, provider, principalStore)
	if err != nil {
		return nil, err
	}
	webhookController := webhook2.ProvideController(config, db, authorizer, webhookStore, webhookExecutionStore, repoStore, webhookServer)
	reporter, err := events2.ProvideReporter(eventsSystem)
	if err != nil {
		return nil, err
	}
	githookController := githook.ProvideController(db, authorizer, principalStore, repoStore, reporter)
	serviceAccount := check.ProvideServiceAccountCheck()
	serviceaccountController := serviceaccount.NewController(serviceAccount, authorizer, principalStore, spaceStore, repoStore, tokenStore)
	apiHandler := router.ProvideAPIHandler(config, authenticator, repoController, spaceController, pullreqController, webhookController, githookController, serviceaccountController, controller)
	gitHandler := router.ProvideGitHandler(config, provider, repoStore, authenticator, authorizer, gitrpcInterface)
	webHandler := router.ProvideWebHandler(config)
	routerRouter := router.ProvideRouter(apiHandler, gitHandler, webHandler)
	serverServer := server.ProvideServer(config, routerRouter)
	serverConfig := ProvideGitRPCServerConfig(config)
	server3, err := server2.ProvideServer(serverConfig)
	if err != nil {
		return nil, err
	}
	nightly := cron.NewNightly()
	serverSystem := newSystem(bootstrapBootstrap, serverServer, server3, webhookServer, nightly)
	return serverSystem, nil
}
