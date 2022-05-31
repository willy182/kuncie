// Code generated by candi v1.11.5.

package configs

import (
	"github.com/golangid/candi/codebase/factory"
	"github.com/golangid/candi/codebase/factory/appfactory"
	"github.com/golangid/candi/config/env"
)

/*
InitAppFromEnvironmentConfig constructor

Construct server/worker for running application from environment value

## Server
USE_REST=[bool]
USE_GRPC=[bool]
USE_GRAPHQL=[bool]

## Worker
USE_KAFKA_CONSUMER=[bool] # event driven handler
USE_CRON_SCHEDULER=[bool] # static scheduler
USE_REDIS_SUBSCRIBER=[bool] # dynamic scheduler
USE_TASK_QUEUE_WORKER=[bool]
USE_POSTGRES_LISTENER_WORKER=[bool]
USE_RABBITMQ_CONSUMER=[bool] # event driven handler and dynamic scheduler
*/
func InitAppFromEnvironmentConfig(service factory.ServiceFactory) (apps []factory.AppServerFactory) {

	if env.BaseEnv().UseKafkaConsumer {
		apps = append(apps, appfactory.SetupKafkaWorker(service))
	}
	if env.BaseEnv().UseCronScheduler {
		apps = append(apps, appfactory.SetupCronWorker(service))
	}
	if env.BaseEnv().UseTaskQueueWorker {
		apps = append(apps, appfactory.SetupTaskQueueWorker(service))
	}
	if env.BaseEnv().UseRedisSubscriber {
		apps = append(apps, appfactory.SetupRedisWorker(service))
	}
	if env.BaseEnv().UsePostgresListenerWorker {
		apps = append(apps, appfactory.SetupPostgresWorker(service))
	}
	if env.BaseEnv().UseRabbitMQWorker {
		apps = append(apps, appfactory.SetupRabbitMQWorker(service))
	}

	if env.BaseEnv().UseREST {
		apps = append(apps, appfactory.SetupRESTServer(service))
	}
	if env.BaseEnv().UseGRPC {
		apps = append(apps, appfactory.SetupGRPCServer(service))
	}
	if !env.BaseEnv().UseREST && env.BaseEnv().UseGraphQL {
		apps = append(apps, appfactory.SetupGraphQLServer(service))
	}

	return
}
