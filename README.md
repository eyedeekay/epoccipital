# epoccipital


epoccipital is is an I2P-Hosted Tailscale control server based on the Open-Source Headscale server

Variant created by
eyedeekay
https
//github.com/eyedeekay/epoccipital

Usage

        epoccipital [command]

Available Commands

        apikeys                 Handle the Api keys in epoccipital
        completion        Generate the autocompletion script for the specified shell
        debug                         debug and testing commands
        generate                Generate commands
        help                                Help about any command
        hiddenserve Launches the epoccipital server as a hidden(I2P) service
        mockoidc                Runs a mock OIDC server for testing
        namespaces        Manage the namespaces of epoccipital
        nodes                         Manage the nodes of epoccipital
        preauthkeys Handle the preauthkeys in epoccipital
        routes                        Manage the routes of epoccipital
        version                 Print the version.

Flags

        -c, --config string                config file (default is /etc/headscale/config.yaml)
                        --force                                                Disable prompts and forces the execution
        -h, --help                                                 help for epoccipital
        -o, --output string                Output format. Empty for human-readable, 'json', 'json-line' or 'yaml'
        -s, --samaddr string         Address of the SAMv3 API for I2P (default "127.0.0.1
7656")
        -n, --tunname string         Name to use for the I2P tunnel (default "epoccipital")

Use "epoccipital [command] --help" for more information about a command.

## apikeys

Handle the Api keys in epoccipital

Usage

        epoccipital apikeys [command]

Aliases

        apikeys, apikey, api

Available Commands

        create                        Creates a new Api key
        expire                        Expire an ApiKey
        list                                List the Api keys for epoccipital

Flags

        -h, --help         help for apikeys

Global Flags

        -c, --config string                config file (default is /etc/headscale/config.yaml)
                        --force                                                Disable prompts and forces the execution
        -o, --output string                Output format. Empty for human-readable, 'json', 'json-line' or 'yaml'
        -s, --samaddr string         Address of the SAMv3 API for I2P (default "127.0.0.1
7656")
        -n, --tunname string         Name to use for the I2P tunnel (default "epoccipital")

Use "epoccipital apikeys [command] --help" for more information about a command.

## completion

Generate the autocompletion script for epoccipital for the specified shell.
See each sub-command's help for details on how to use the generated script.

Usage

        epoccipital completion [command]

Available Commands

        bash                                Generate the autocompletion script for bash
        fish                                Generate the autocompletion script for fish
        powershell        Generate the autocompletion script for powershell
        zsh                                 Generate the autocompletion script for zsh

Flags

        -h, --help         help for completion

Use "epoccipital completion [command] --help" for more information about a command.

## debug

debug contains extra commands used for debugging and testing epoccipital

Usage

        epoccipital debug [command]

Available Commands

        create-node Create a node (machine) that can be registered with `nodes register <>` command

Flags

        -h, --help         help for debug

Global Flags

        -c, --config string                config file (default is /etc/headscale/config.yaml)
                        --force                                                Disable prompts and forces the execution
        -o, --output string                Output format. Empty for human-readable, 'json', 'json-line' or 'yaml'
        -s, --samaddr string         Address of the SAMv3 API for I2P (default "127.0.0.1
7656")
        -n, --tunname string         Name to use for the I2P tunnel (default "epoccipital")

Use "epoccipital debug [command] --help" for more information about a command.

## generate

Generate commands

Usage

        epoccipital generate [command]

Aliases

        generate, gen

Available Commands

        private-key Generate a private key for the epoccipital server

Flags

        -h, --help         help for generate

Global Flags

        -c, --config string                config file (default is /etc/headscale/config.yaml)
                        --force                                                Disable prompts and forces the execution
        -o, --output string                Output format. Empty for human-readable, 'json', 'json-line' or 'yaml'
        -s, --samaddr string         Address of the SAMv3 API for I2P (default "127.0.0.1
7656")
        -n, --tunname string         Name to use for the I2P tunnel (default "epoccipital")

Use "epoccipital generate [command] --help" for more information about a command.

## hiddenserve

Launches the epoccipital server as a hidden(I2P) service

Usage

        epoccipital hiddenserve [flags]

Flags

        -h, --help         help for hiddenserve

Global Flags

        -c, --config string                config file (default is /etc/headscale/config.yaml)
                        --force                                                Disable prompts and forces the execution
        -o, --output string                Output format. Empty for human-readable, 'json', 'json-line' or 'yaml'
        -s, --samaddr string         Address of the SAMv3 API for I2P (default "127.0.0.1
7656")
        -n, --tunname string         Name to use for the I2P tunnel (default "epoccipital")

## mockoidc

This internal command runs a OpenID Connect for testing purposes

Usage

        epoccipital mockoidc [flags]

Flags

        -h, --help         help for mockoidc

## namespaces

Manage the namespaces of epoccipital

Usage

        epoccipital namespaces [command]

Aliases

        namespaces, namespace, ns, user, users

Available Commands

        create                        Creates a new namespace
        destroy                 Destroys a namespace
        list                                List all the namespaces
        rename                        Renames a namespace

Flags

        -h, --help         help for namespaces

Global Flags

        -c, --config string                config file (default is /etc/headscale/config.yaml)
                        --force                                                Disable prompts and forces the execution
        -o, --output string                Output format. Empty for human-readable, 'json', 'json-line' or 'yaml'
        -s, --samaddr string         Address of the SAMv3 API for I2P (default "127.0.0.1
7656")
        -n, --tunname string         Name to use for the I2P tunnel (default "epoccipital")

Use "epoccipital namespaces [command] --help" for more information about a command.

## nodes

Manage the nodes of epoccipital

Usage

        epoccipital nodes [command]

Aliases

        nodes, node, machine, machines

Available Commands

        delete                        Delete a node
        expire                        Expire (log out) a machine in your network
        list                                List nodes
        move                                Move node to another namespace
        register                Registers a machine to your network
        rename                        Renames a machine in your network
        tag                                 Manage the tags of a node

Flags

        -h, --help         help for nodes

Global Flags

        -c, --config string                config file (default is /etc/headscale/config.yaml)
                        --force                                                Disable prompts and forces the execution
        -o, --output string                Output format. Empty for human-readable, 'json', 'json-line' or 'yaml'
        -s, --samaddr string         Address of the SAMv3 API for I2P (default "127.0.0.1
7656")
        -n, --tunname string         Name to use for the I2P tunnel (default "epoccipital")

Use "epoccipital nodes [command] --help" for more information about a command.

## preauthkeys


## routes

Manage the routes of epoccipital

Usage

        epoccipital routes [command]

Aliases

        routes, r, route

Available Commands

        disable                 Set as disabled a given route
        enable                        Set a route as enabled
        list                                List all routes

Flags

        -h, --help         help for routes

Global Flags

        -c, --config string                config file (default is /etc/headscale/config.yaml)
                        --force                                                Disable prompts and forces the execution
        -o, --output string                Output format. Empty for human-readable, 'json', 'json-line' or 'yaml'
        -s, --samaddr string         Address of the SAMv3 API for I2P (default "127.0.0.1
7656")
        -n, --tunname string         Name to use for the I2P tunnel (default "epoccipital")

Use "epoccipital routes [command] --help" for more information about a command.

## version

The version of epoccipital.

Usage

        epoccipital version [flags]

Flags

        -h, --help         help for version

