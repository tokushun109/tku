export const API_CLUSTER_NAME = 'tku-api-cluster'
export const DB_CLUSTER_NAME = 'tku-db-cluster'

export const ECS_TASK_ASSUME_ROLE_POLICY = {
    Version: '2012-10-17',
    Statement: [
        {
            Effect: 'Allow',
            Principal: {
                Service: 'ecs-tasks.amazonaws.com',
            },
            Action: 'sts:AssumeRole',
        },
    ],
}
