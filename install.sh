#!/bin/bash

currentDir="$(pwd)"

importPath="${currentDir#*/src/}"

pkgName="$(go list -e -f '{{.ImportComment}}' 2>/dev/null || true)"

sed -i "" "s#${pkgName}#${importPath}#g" main.go
