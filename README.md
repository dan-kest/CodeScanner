# Code Scanner

A code scanner application example.

[API Documentation](document/APIDOC.md)

### What does it do?

- You can CRUD git repositories (repository name, URL)
- You can trigger a repository scan and it will do its things
- You can check the scan status if its still scanning, success, or fail
- You can then view the scan result

### How do I setup and run this thing?

See [SETUP.md](document/SETUP.md) and hope it works. 🤞

### Explain these directories to me

- `/cmd` :: Main application
- `/config` :: Application configuration files
- `/database` :: Database connection utility files
- `/docker` :: docker-compose.yaml and stuff along with Docker configuration files
- `/document` :: Documents!
- `/http` :: HTTP entry point and initialization
- `/internal` :: Main application code!
  - `/constants` :: Constants for everything
  - `/handlers` :: Controller
    - `/payloads` :: Controller models
  - `/interfaces` :: Interfaces
  - `/models` :: Business models
  - `/repositories` :: Data layer
    - `/tables` :: Data layer models
  - `/routers` :: Route and route group
  - `/service` :: Business logic
- `/pkg` :: Shared Utilitiy files
- `/queue` :: Message queue entry point and initialization

So what I try to do is to make the structure as flat as possible while maintaining some key point of Clean Architecture™.

I like the idea of layers separation of clean architecture specifically between service and repository, so that I can plug them in or out or easily mock things for testing business logic.

### What's your database schema looks like?

See comments in [schema.sql](document/schema.sql)
