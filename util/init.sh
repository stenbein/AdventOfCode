#!/bin/bash

#Quick script to gen the project directory.

year=$1

mkdir "${year}"
pushd "${year}" > /dev/null

mkdir -p day{01..25}
for D in *;
    do touch "${D}"/{input,README.md};
done

popd > /dev/null
