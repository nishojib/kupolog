# FFXIV Dailies

FFXIV dailies is an app to help you track your dailies, weeklies and custom tasks.

## Installation

You can install the app two ways. The easiest way to run the app is using docker as the other option is a bit more involved.

### Using Docker

Copy and create an .env file following the .env.example in both the client and backend projects. 

We use [Turso](https://turso.tech/) for our database. Please follow the [guide](https://docs.turso.tech/quickstart) to setup your own turso database.

Use the docker compose command to run locally.

```zsh
docker-compose -f docker-compose.dev.yml up 
```

### Using local environment

Follow the steps mentioned in the following sections:
1. [Frontend](/client/README.md)
2. [Backend](/backend/README.md)


## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)