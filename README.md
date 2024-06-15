# Service Registry

Implementation of a simple service registry in Go. 🪵

# Introduction

A Service Registry is a critical component in a microservices architecture. It serves as a central directory that keeps track of service instances and their locations, enabling service discovery. The Service Registry will keep track of all service instances in a microservices architecture. It will provide mechanisms for services to register, deregister, and discover other services.

# Features

- **Service Registration:** Services must be able to register themselves with the Service Registry.
- **Service Deregistration:** Services must be able to deregister themselves.
- **Service Discovery:** Services must be able to discover other services.
- **Health Checks:** The registry should support health checks to ensure that only healthy services are discoverable.
- **Load Balancing:** The registry should provide information that can be used for load balancing requests among multiple service instances.
- **TTL (Time to Live) Management:** Services should periodically renew their registration to indicate they are alive. The registry should deregister services that do not renew within a specified TTL.
