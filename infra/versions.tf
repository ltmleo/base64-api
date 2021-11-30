terraform {
  backend "remote" {
    organization = "aws-testes"
    workspaces {
      name = "base64-api"
    }
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.25.0"
    }
  }
}
provider "aws" {
  region = "us-east-2"
}
