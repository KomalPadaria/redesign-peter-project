# trust-portal-api

The Redesign Group Trust Portal API application.


## Setting up the environment
### Prerequisites
- Install Golang (if you haven't already). https://go.dev/doc/install

    *Note: Golang version 1.19*
- Instal docker (If you haven't already). https://docs.docker.com/engine/install/
### Project setup
- Checkout the project locally
  ```bash
  git clone https://github.com/nurdsoft/redesign-grp-trust-portal-api
  ```
- Change directory to the project folder
  ```bash
  cd redesign-grp-trust-portal-api
  ```
- To install the necessary libraries run the below command
  ```bash
  make setup
  ```
  
## Application Configuration
The application configuration is in the file `./config.yaml`. Please change the values appropriately based on the environment.

## Environment Variables
You can override the configuration values by setting as environment values. 
Below are the environment variables that can override the config.yaml values. 

> Copy these values in a `.env` file and run `source .env` to export them on your system.

```
# Common envs
export REDESIGN_COMMON_NAME="redesign-v1"
export REDESIGN_COMMON_ENV="local"
export REDESIGN_COMMON_VERSION="1.0"
export REDESIGN_COMMON_USERAGENT="redesign/1.0"
export REDESIGN_COMMON_COMPONENT="API"
export REDESIGN_COMMON_FRONTENDDOMAIN="rdf.nurdsoft.co"

export REDESIGN_TRANSPORT_HTTP_PORT=8080
export REDESIGN_TRANSPORT_GRPC_PORT=9002

export REDESIGN_DB_POSTGRES_HOST="localhost"
export REDESIGN_DB_POSTGRES_PORT="5442"
export REDESIGN_DB_POSTGRES_DATABASE="redesign"
export REDESIGN_DB_POSTGRES_USER="db"
export REDESIGN_DB_POSTGRES_PASSWORD=123

#Salesforce environment variables
export REDESIGN_SALESFORCE_APIHOST=""
export REDESIGN_SALESFORCE_APIVERSION="v55.0"
export REDESIGN_SALESFORCE_CLIENTID=""
export REDESIGN_SALESFORCE_CLIENTSECRET=""
export REDESIGN_SALESFORCE_USERNAME=""
export REDESIGN_SALESFORCE_PASSWORD=""

# AWS SES
export REDESIGN_SES_AWSREGION="us-west-1"
export REDESIGN_SES_INVITESENDEREMAIL=""

#Calendly environment variables
export REDESIGN_CALENDLY_ACCESSTOKEN=xxxx
export REDESIGN_CALENDLY_APIHOST="https://api.calendly.com"

#Rapid7
export REDESIGN_RAPID7_SITENAME=""
export REDESIGN_RAPID7_DATAWAREHOUSE_POSTGRES_HOST=""
export REDESIGN_RAPID7_DATAWAREHOUSE_POSTGRES_PORT="5432"
export REDESIGN_RAPID7_DATAWAREHOUSE_POSTGRES_DATABASE=""
export REDESIGN_RAPID7_DATAWAREHOUSE_POSTGRES_USER=""
export REDESIGN_RAPID7_DATAWAREHOUSE_POSTGRES_PASSWORD=""

#JIRA
export REDESIGN_JIRA_USERNAME=""
export REDESIGN_JIRA_APITOKEN=""
```


## Running the application
## Local Env
### Start Environment
- To start the services that is required for the application(ex: database) run the below command
  ```bash
  make start-env
  ```
### Build application
- Build the application using the below command
  ```bash
  make build-dev
  ```
### Running application 
- Run the application locally using the below command
  ```bash
  make run-dev
  ```
  > Once you run the application, the swagger ui docs are available at [http://127.0.0.1:8080/swaggerui/docs/swagger/](http://127.0.0.1:8080/swaggerui/docs/swagger/).

### Running application in docker container
- Run the application in the docker container using the below command
  ```bash
  make start-app
  ```
### Stop Environment
- To stop all the services run the below command
  ```bash
  make stop-env
  ```
  
## Kick-start running the whole application 
- To run the all the services including the application run the below commands
  ```bash
  git clone https://github.com/nurdsoft/redesign-grp-trust-portal-api
  cd redesign-grp-trust-portal-api
  make start-all
  ```

## Database migrations
We use [sql-migrate](https://github.com/rubenv/sql-migrate) for database migrations
- To create new migration
  ```bash
  make name=migration_script_file_name new-migration
  ```
- To apply outstanding migrations
  ```bash
  make migrate
  ```

## CI/CD

The project includes a GitHub Action to automatically **build from all branches** and **deploy from the main** branch. See the [GitHub Workflow file](https://github.com/nurdsoft/redesign-grp-trust-portal-api/blob/main/.github/workflows/build_and_deploy.yml) for details.

Upon a successful deployment, the service can be accessed at: **https://proxy.pacenthink.co** and the **service name as the host header**. E.g.:
```sh
$ curl -H "Host: ghcr-io-nurdsoft-redesign-api-main" https://proxy.pacenthink.co/health
```

### CI/CD - Environment Variables

All **sensitive** environment variables like usernames, passwords, keys, etc. required to run the service must be configured as a [GitHub Secret](https://github.com/nurdsoft/redesign-grp-trust-portal-api/settings/secrets/actions) and the [Deploy GitHub Action](https://github.com/nurdsoft/redesign-grp-trust-portal-api/blob/main/.github/workflows/build_and_deploy.yml#L57-L67).
  - **Important Note**: If you need access to create/update a GitHub Secret, reach out to a team member for assistance.

## Help
Run `make help` command to list all the available make commands to make the development life easier.
