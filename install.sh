#!/bin/bash

currentDir="$(pwd)"

importPath="${currentDir#*/src/}"

pkgName="$(go list -e -f '{{.ImportComment}}' 2>/dev/null || true)"

appKey="$(date | md5)"

echo $appKey
sed -i "" "s#APPKEY#${appKey}#g" ./config/local.yml
sed -i "" "s#${pkgName}#${importPath}#g" main.go


