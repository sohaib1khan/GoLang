#!/bin/bash
# Add the working directory to the PATH
export PATH=$PATH:$(pwd)

# Set and export the PS1 environment variable for the custom prompt
export PS1="[\u@\h ${ENV_NAME}]$ "

echo "Welcome to the custom CLI container!"
echo "Your persistent data directory is mounted at: $HOST_DIR"
echo "Type 'list-tools' to see the installed tools."
echo "Type 'exit' to leave the container."
echo

# Enter an interactive shell
exec bash
