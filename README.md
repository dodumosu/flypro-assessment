# flypro-assessment
This application is for the FlyPro assessment.

The application is configured using environment variables. The file
`.env-template` lists the environment variables required for
setting up the application.

One setting of note is SERVER_ALLOWED_ORIGINS. It is *required*. The app
will throw a fit and panic if it is not set.
