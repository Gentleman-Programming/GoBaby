# GoBaby App

GoBaby is a Go application designed to help parents track feeding times for newborn babies. The application utilizes Clean Architecture principles to ensure separation of concerns, maintainability, and testability.

## Main Features

- **Countdown Timer**: The app starts a countdown timer set to 4 hours, simulating the interval between feedings for a newborn baby.
- **Reset Button**: Users can reset the countdown timer with a button click. When the timer is reset, the app logs the reset time along with the date.
- **Logging**: The application logs each feeding time, including the time at which the feeding occurred and the date.
- **Log Viewer**: Users can view a log of all feeding times in a table format.

## Folder Structure

The folder structure follows Clean Architecture principles, separating concerns into layers:

- **cmd**: Contains the main application entry point (`main.go`) and related documents.
- **internal**: Houses the core business logic of the application.
  - **models**: Defines data models such as logs and routes.
  - **utils**: Provides utility functions like API helpers and clock utilities.
- **web**: Implements the web interface and handles HTTP requests.
  - **domain**: Contains domain-specific logic.
    - **error**: Handles error handling functionality, including rendering error templates.
    - **log**: Handles logging functionality.
    - **main**: Manages the main application logic, including the countdown timer.
    - **options**: Handles application options and configurations.
    - **repository**: Implements data access logic using adapters and configurations.
  - **routes**: Defines HTTP routes for different functionalities.
- **ui**: Contains HTML templates and static assets for the user interface.
  - **html**: Houses HTML pages for different parts of the application.
    - **pages**: Includes pages for logs, main functionality, options, and error handling.
  - **static**: Stores CSS stylesheets and other static assets.

```mermaid
graph LR;
A[~/work/GoBaby]

A --> documents
A --> GoBaby
A --> go.mod
A --> go.sum
A --> web

cmd --> web
    web --> domain
        domain --> error
            error --> error.domain.go
        domain --> log
            log --> log.domain.go
        domain --> main
            main --> clock.domain.go
            main --> main.domain.go
        domain --> options
            options --> options.domain.go
        domain --> repository
            repository --> adapters
                adapters --> user.adapter.go
            repository --> config
                config --> db.config.go
            repository --> repository.domain.go
    web --> routes
        routes --> mainRoute
            mainRoute --> clock.route.go
            mainRoute --> main.route.go
        routes --> log.route.go
        routes --> options.route.go
        routes --> starter.route.go
    web --> main.go

cmd --> documents
    documents --> entities.mkd
    documents --> structure.mkd

cmd --> internal
    internal --> models
        models --> log.model.go
        models --> routes.model.go
        models --> user.model.go
        models --> error.model.go
    internal --> utils
        utils --> api.go
        utils --> clockUtils.go
        utils --> template.go

cmd --> ui
    ui --> html
        html --> pages
            pages --> logs
                logs --> logTable.tpml.html
                logs --> logs.tmpl.html
            pages --> main
                main --> clock.tmpl.html
                main --> main.tmpl.html
            pages --> options
                optinos --> options.tmpl.html
            pages --> error
                error --> error.tmpl.html
        html --> base.html
    ui --> static
        static --> css
            css --> main.css
```

## Clean Architecture

GoBaby follows Clean Architecture principles to achieve a modular and maintainable codebase. The application is structured in layers, with dependencies flowing inward toward the core business logic. This architecture allows for easy testing, scalability, and flexibility in adapting to future requirements.

## Installation and Usage

To run the application, ensure you have Go installed on your machine. Clone the repository and navigate to the `cmd/web` directory. Run the `main.go` file to start the application. Access the application through a web browser at the specified port.

## Contributing

Contributions to GoBaby are welcome! Feel free to submit bug reports, feature requests, or pull requests through GitHub.

## Organizing Files Based on Logic

To maintain a structured and organized codebase, it's essential to follow a consistent approach when adding new files. This section outlines guidelines for organizing files based on the type of logic they represent within the application.

### Domain Logic

Domain logic represents the core business rules and operations of the application. When adding files related to domain logic, follow these guidelines:

- **Folder**: Place domain logic files under the `web/domain` directory.
- **Naming**: Use descriptive names for files and directories that reflect the specific domain aspect they address (e.g., `log`, `main`, `options`).
- **Responsibilities**: Each domain module should encapsulate related functionality, such as handling logs, managing main application logic, or dealing with application options.

### Data Access Logic

Data access logic handles interactions with databases, external APIs, or other data sources. Follow these guidelines when adding files related to data access:

- **Folder**: Place data access logic files under the `web/domain/repository` directory.
- **Adapters**: Use the `adapters` directory to store data access adapters, which encapsulate interactions with external data sources.
- **Configurations**: Store database configurations and other data access configurations under the `config` directory.
- **Responsibilities**: Separate data access logic based on the type of data source or functionality it addresses.

### User Interface Logic

User interface logic manages the presentation layer of the application, including HTML templates, CSS stylesheets, and static assets. Here's how to organize UI-related files:

- **Folder**: Place UI-related files under

the `ui` directory.

- **HTML Templates**: Store HTML templates under the `html/pages` directory, organized by different application views (e.g., `logs`, `main`, `options`).
- **Static Assets**: Store CSS stylesheets and other static assets under the `static/css` directory.

### Internal Logic and Utilities

Internal logic and utility functions support the overall operation of the application but are not directly tied to domain-specific or user interface logic. Follow these guidelines for organizing internal logic and utilities:

- **Folder**: Place internal logic and utility files under the `internal` directory.
- **Models**: Define data models and structures under the `internal/models` directory.
- **Utilities**: Store utility functions, such as API helpers and clock utilities, under the `internal/utils` directory.

By following these guidelines, developers can ensure a consistent and structured approach to organizing files within the GoBaby application, promoting clarity, maintainability, and collaboration across the development team.

## Managing Errors

Proper error handling is crucial for maintaining a robust and reliable application. GoBaby adopts a centralized error management approach to ensure consistent error handling throughout the application. Here's how error management is implemented:

### Error Domain

The `errorDomain` package contains functionality related to error handling, including rendering error templates. Error-related logic is encapsulated within this domain to keep the code organized and maintainable.

### Error Rendering

The `ErrorTemplate` function in the `errorDomain` package is responsible for rendering error templates. When an error occurs, this function is called to display the appropriate error message to the user.

### Error Models

The `AppError` struct in the `models` package defines the structure of application errors. Each error includes a message, code, and underlying error, providing comprehensive information for debugging and troubleshooting.

### Error Codes

GoBaby defines custom error codes to categorize different types of errors that may occur within the application. These codes help developers identify the nature of the error and take appropriate action to resolve it.

### Error Handling in Application Logic

Throughout the application logic, errors are handled using the `AppError` struct and associated error codes. When an error occurs, it is wrapped in an `AppError` instance and returned to the caller, ensuring consistent error propagation.

By following these error management practices, GoBaby maintains a reliable and user-friendly error handling system, enhancing the overall robustness and usability of the application.
