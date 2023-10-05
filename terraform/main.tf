terraform {
  backend "s3" {
    bucket         = "lol-counter-bot-state"
    key            = "global/s3/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "lol-counter-bot-state-lock"
    encrypt        = true
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.4"
    }
  }
  required_version = "~> 1.5"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE VPC FOR FARGATE SERVICE
# ---------------------------------------------------------------------------------------------------------------------

module "vpc" {
  source = "./vpc"
  name   = "lol-counter-bot-${terraform.workspace}"
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE ECS/FARGATE CLUSTER/TASK
# ---------------------------------------------------------------------------------------------------------------------

module "fargate" {
  source = "./fargate"

  name           = "lol-counter-bot-${terraform.workspace}"
  subnet_id      = "${module.vpc.subnet_id}"
  image          = "austin1237/lol-counter-bot@${var.DOCKER_IMAGE_SHA}"
  cpu            = 256
  memory         = 512
  desired_count  = 1

  env_vars = {
    COUNTER_API_URL = var.COUNTER_API_URL
    DISCORD_BOT_TOKEN = var.DISCORD_BOT_TOKEN
  }
}
