#!/bin/bash
for chal in *; do
    if [ -d "${chal}" ] && [ -f "${chal}/Dockerfile" ]; then
        docker build -t "${chal}" "${chal}"
    fi
done
