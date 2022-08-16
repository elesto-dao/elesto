# Technical setup for the docs portal

For a successful experience working with the docs, ensure your local technical setup meets these requirements.

## Codebase

Fork or clone the <https://github.com/elesto-dao/elesto-docs/> repository.

Internal users have different permissions. If you're not sure if you have access, fork the repo.

## Software requirements

Ensure your local environment meets these software requirements.

## Setting up your local environment

<!-- do docs need pre-commit hooks? -->

To install the configured pre-commit hooks, run these commands from the project's root folder:

```sh
pre-commit install
pre-commit install --hook-type commit-msg
```

## Updating docs

Thank you for making relevant documentation updates. We appreciate your help preserving and maintaining accurate and trusted technical content. All code updates are considered complete only after accompanying documentation is submitted.

## Building docs

The documentation portal is generated using the [mkdocs](https://www.mkdocs.org/) static site generator. Documentation source files are written in Markdown and configured with a single [mkdocs.yml]() configuration file.

To build or run the documentation portal locally, the requirements are:

- pip3
- Python v3.8 or higher

To install the required packages, run these commands at the command line:

```sh
pip3 install mkdocs
pip3 install -r requirements.txt
```

To build the documentation portal, execute the following command from the project root:

```sh
mkdocs build
```

The documentation portal is built in the `site` directory.

To run the documentation portal, run the following command from the project root directory:

```sh
mkdocs serve
```

In a web browser, the docs portal is served on 'http://127.0.0.1:8000/'.

## GitHub integration

In the Visual Studio Code sidebar, click the GitHub icon and follow the prompts.
