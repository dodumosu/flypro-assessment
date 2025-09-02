# flypro-assessment
This application is for the FlyPro assessment.

### configuration
The application is configured using environment variables. The file
`.env-template` lists the environment variables required for
setting up the application. This file should be used as a guide to
setting up the required environment variables. If there is a `.env`
file present, the application will attempt to read from it.

One setting of note is SERVER_ALLOWED_ORIGINS. It is *required*. The app
will throw a fit and panic if it is not set.

### dependency services
The application requires the following to operate:
- a PostgreSQL server and database
- a Redis (or Redis-compatible equivalent like Dragonfly/Valkey) server
- an [APILayer exchange rate](https://apilayer.com/marketplace/exchangerates_data-api) key

Each of the above services/resources has associated settings.

### database migrations
Migrations use the [goose](https://github.com/pressly/goose) migration tool. The included makefile contains targets for managing migrations. There is a separate section of `.env-template` for migrations, because goose can read from the environment as well.
