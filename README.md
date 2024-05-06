# go-anhgelus

go-anhgelus is a URL Shortener web application.

## Installation

There are two ways to install the bot: docker and build.

### Docker

1. Clone the repository
```bash
$ git clone https://github.com/anhgelus/go-anhgelus.git 
```
2. Go into the repository, rename `.env.example` into `.env` and customize it: add the wanted links to be shortened.
3. Start the compose file
```bash
$ docker compose up -d --build 
```

### Build

1. Clone the repository
```bash
$ git clone https://github.com/anhgelus/go-anhgelus.git 
```
2. Install Go 1.22+
3. Go into the repository and build the program
```bash
$ go build . 
```
4. Run the application through bash (or PowerShell if you are on windows)

#### Usage

- `go-anhgelus` and `go-anhgelus -h` show the help
- `go-anhgelus run` runs the application
- `go-anhgelus config {path} {url}` creates (or edit) a config at the given path for shortening the given url
`go-anhgelus config foo.toml example.org` will create a config at `config/foo.toml` with a shortening link for 
`example`.org

## Config

Every config files are in `config` folder.
They followed this template:
```toml
[[links]]
id = "slug"
link = "https://foo.example.org"
```

- `[[links]].id` is the slug used. You can modify it.
- `[[links]].link` is the shortened link.

You can create as many config files as you want.
It also supports config is subdirectories

## Technologies

- Go 1.22
- gorilla/mux
- anhgelus/human-readable-slug
- pelletier/go-toml
