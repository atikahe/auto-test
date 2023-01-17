#!/bin/bash

set -e

# base_url=https://github.com/atikahe/auto-test/releases/latest/download

main() {
    # set ${variable-default}
    BIN_DIR=${BIN_DIR-"$HOME/.bin"}
    mkdir -p BIN_DIR

    # Detect shell
    case $SHELL in
    */zsh)
        PROFILE=$HOME/.zshrc
        PREF_SHELL=zsh
        ;;
    */bash)
        PROFILE=$HOME/.bashrc
        PREF_SHELL=bash
        ;;
    */ash)
        PROFILE=$HOME/.PROFILE
        PREF_SHELL=ash
        ;;
    *)
        echo "could not detect shell, manually add ${BIN_DIR} to your path."
        exit 1
    esac

    # Add BIN_DIR to path if not exist
    if [[ ":$PATH:" !=  *":${BIN_DIR}:"* ]]; then
        echo >> $PROFILE && echo "export PATH=\"\$PATH:$BIN_DIR\"" >> $PROFILE
    fi

    # Detect OS & Arch
    PLATFORM="$(uname -s)"
    case $PLATFORM in
    Linux)
        PLATFORM="linux"
        ;;
    Darwin)
        PLATFORM="darwin"
        ;;
    Windows)
        PLATFORM="windows"
        ;;
    *)
        echo "unsupported platform: $PLATFORM"
        exit 1
    esac

    ARCH="$(uname -m)"
    if [ "${ARCH}" == "x86_64"]; then
    if [ "$(sysctl -n sysctl.proc_translated 2>/dev/null)" != "1" ]; then
        ARCH="x86_64"
    else    
        echo "unsupported architecture"
        exit 1
    fi
    elif [ "${ARCH}" = "arm64" ] || [ "${ARCH}" = "aarch64" ]; then
        echo "unsupported architecture: ${ARCH}"
        exit 1
    fi

    # Download url
    BIN_URL="https://github.com/atikahe/auto-test/releases/latest/download/auto-test-${PLATFORM}-${ARCH}"
    echo BIN_URL

    # Start download
    echo "Downloading latest binary"
    catch curl -L "$BIN_URL" -o "$BIN_DIR/auto-test"

    # Make executable
    chmod +x "$BIN_DIR/auto-test"

    echo "Installed - $("$BIN_DIR/auto-test" --version)"
}

# If argument fails, print error message and terminate
catch() {
    if ! "$@"; then
        echo"command failed: $*"
        exit 1
    fi
}

main "$@" || exit 1