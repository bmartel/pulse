#!/bin/bash

pkgName="$(go list -e -f '{{.ImportComment}}' 2>/dev/null || true)"

docker build -t "$pkgName" .
