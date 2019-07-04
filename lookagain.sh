#! /bin/bash

echo "$(find -name '*.sh' -exec basename \{} .sh \; )"

