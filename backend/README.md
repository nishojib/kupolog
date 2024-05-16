# FFXIV Dailies API

API consumed by the FFXIV Dailies app.

## Requirements

Make sure the followings are installed in your system:

1. [Go](https://go.dev/)
2. [Goose](https://pressly.github.io/goose/)
3. [sqlc](https://sqlc.dev/)
4. [turso](https://turso.tech/)
5. [air](https://github.com/cosmtrek/air)
   
## Installation

To install the app run make install:
```zsh
make install
```

## Run

To run the app locally follow the steps:

1. Make sure there is a `.env` file in the root of the `backend`. You can copy `.env.example`  and rename to `.env`. Add the values for each item in the file.
2. The app comes with a `.air.toml` file. You can use it to run the app for live reloading.
Run the air command.
  ```zsh
  air
  ```
3. (Optional) While it is recommended to run the app via air for developer experience you can also run the make run/api command to run the app.
```zsh
make run/api
```