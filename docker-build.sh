#!/bin/bash
for chal in *; do
    if [ -d "${chal}" ] && [ -f "${chal}/Dockerfile" ]; then
        docker build -t "${chal}" "${chal}"
    fi
done

for chal in WorkExperience; do
if [ -d "WorkExperience/${chal}" ] && [ -f "WorkExperience/${chal}/Dockerfile" ]; then
    docker build -t "${chal}" "WorkExperience/${chal}"
fi
done
