# Docker ContainerConfig Error - shared

## Symptoms
- Running docker-compose up --build or docker compose up --build fails with:
```
    ERROR: for <service> 'ContainerConfig'
    ERROR: for <service-name> 'ContainerConfig'
```
- The service is not recreated or started.
- Logs provide no specific error details beyond 'ContainerConfig'.
- docker-compose build may appear to work silently or exit with vague errors.

## Diagnostic Steps
1. Check docker logs: `make logs`
2. Deploy the platform: `make deploy`
3. Override CMD to keep container running: `CMD ["sleep", "infinity"]`, then connect with: `docker exec -it <container-name> /bin/bash`

## Recovery Procedures
1. Stop and delete containers: `make down`
2. Clean docker images and volumes: `make docker-clean`
3. Check and correct environment variables in .env for affected containers.
4. Rebuild all services. 
5. Redeploy the platform: `make deploy`