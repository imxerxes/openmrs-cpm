#!/usr/bin/env bash

SOURCE="${BASH_SOURCE[0]}"
DIR="$( dirname "$SOURCE" )"
${DIR}/buildtools/maven/bin/mvn ${*-package}
