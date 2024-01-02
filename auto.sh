#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <nama module>"
    exit 1
fi

module_name=$1

mkdir -p "./domain/$module_name/dto"
mkdir -p "./domain/$module_name/handler"
mkdir -p "./domain/$module_name/repository"
mkdir -p "./domain/$module_name/service"

# Membuat file-filenya
touch "./domain/$module_name/dto/req.go"
touch "./domain/$module_name/dto/res.go"
touch "./domain/$module_name/handler/index.go"
touch "./domain/$module_name/index.go"
touch "./domain/$module_name/repository/index.go"
touch "./domain/$module_name/service/index.go"

echo "domain '$module_name' berhasil dibuat!"