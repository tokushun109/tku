resource "aws_vpc" "tku-vpc" {
  # (resource arguments)
}

resource "aws_subnet" "tku-public-subnet-a" {
  # (resource arguments)
}

resource "aws_subnet" "tku-public-subnet-c" {
  # (resource arguments)
}

resource "aws_route_table" "public-rt" {
  # (resource arguments)
}

resource "aws_internet_gateway" "tku-igw" {
  # (resource arguments)
}


resource "aws_security_group" "public-ecs-sg" {
  # (resource arguments)
}

resource "aws_security_group" "db-sg" {
  # (resource arguments)
}

resource "aws_ecs_cluster" "tku-ecs-api-cluster" {
  # (resource arguments)
}

resource "aws_ecs_service" "tku-api-service" {

}

resource "aws_ecs_task_definition" "tku-api-task-definition" {

}

resource "aws_secretsmanager_secret" "tku-secret-manager" {

}

