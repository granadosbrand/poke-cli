# Poke-CLI

Poke-CLI is a command-line interface (CLI) for interacting with the PokeAPI. It allows users to search for and capture Pokémon.

## Commands

Here are the available commands in Poke-CLI:

- `help`: Displays a help message with the list of available commands.
- `exit`: Exits Poke-CLI.
- `inspect <pokemon_name>`: Allows you to view the details of a Pokémon you have already captured.
- `catch <pokemon_name>`: Attempts to capture a Pokémon. The higher the base experience of the Pokémon, the lower the chance of capture.
- `pokedex`: List of the Pokémons you have already caught.

## Usage

To use Poke-CLI, simply run the program and then enter commands. For example, this will attempt to capture Pikachu.

```
poke-cli > catch pikachu

```


## Installation

To install Poke-CLI, you can clone this repository and then compile the program with Go. This will create a `poke-cli` executable that you can run to start the program.

```
git clone https://github.com/yourusername/poke-cli.git
cd poke-cli
go build
```
