// ユーザーデータを取得する
export const getUserData = (clusterName: string) => `#!/bin/bash
echo ECS_CLUSTER=${clusterName} >> /etc/ecs/ecs.config;`
