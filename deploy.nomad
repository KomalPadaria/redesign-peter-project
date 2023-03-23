job "rdb" {
  datacenters = ["us-west-2"]
  type = "service"

  # Run the job on the "service" cluster.
  constraint {
    attribute = "${meta.cluster}"
    operator  = "="
    value     = "service"
  }

  group "api" {
    count = 1
    network {
      port "http" {
        to = 8080
      }
      port "grpc" {
        to = 9002
      }
    }
    
    service {
      name = "redesign-api"
      tags = ["urlprefix-/"]
      port = "http"
      check {
        name     = "alive"
        type     = "http"
        path     = "/health"
        interval = "10s"
        timeout  = "2s"
      }
    }

    restart {
      attempts = 2
      interval = "30m"
      delay = "15s"
      mode = "fail"
    }

    task "api" {
      driver = "docker"
      config {
        image = "ghcr.io/nurdsoft/redesign_api:main"
        force_pull = true
        ports = ["http"]
        logging {
          type = "awslogs"
          config {
            awslogs-endpoint  = "logs.us-west-2.amazonaws.com"
            awslogs-group     = "rdb"
          }
        }
      }
      env {
        GITHUB_SHA                        = "${{ github.sha }}"
        REDESIGN_DB_POSTGRES_HOST         = "${{ secrets.REDESIGN_DB_POSTGRES_HOST}}"
        REDESIGN_DB_POSTGRES_USER         = "${{ secrets.REDESIGN_DB_POSTGRES_USER}}"
        REDESIGN_DB_POSTGRES_PASSWORD     = "${{ secrets.REDESIGN_DB_POSTGRES_PASSWORD}}"
        REDESIGN_DB_POSTGRES_PORT         = "5432"
        REDESIGN_SALESFORCE_APIHOST       = "https://redesign-group.my.salesforce.com"
        REDESIGN_SALESFORCE_APIVERSION    = "v55.0"
        REDESIGN_SALESFORCE_CLIENTID      = "${{ secrets.REDESIGN_SALESFORCE_CLIENTID}}"
        REDESIGN_SALESFORCE_CLIENTSECRET  = "${{ secrets.REDESIGN_SALESFORCE_CLIENTSECRET}}"
        REDESIGN_SALESFORCE_USERNAME      = "${{ secrets.REDESIGN_SALESFORCE_USERNAME}}"
        REDESIGN_SALESFORCE_PASSWORD      = "${{ secrets.REDESIGN_SALESFORCE_PASSWORD}}"
      }
    }
  }
}
