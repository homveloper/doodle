#!/bin/bash

# Blog Doodle - Quick Start Script

echo "🔧 Generating Templ templates..."
templ generate

if [ $? -ne 0 ]; then
    echo "❌ Failed to generate templates. Make sure templ is installed:"
    echo "   go install github.com/a-h/templ/cmd/templ@latest"
    exit 1
fi

echo "🚀 Starting blog server..."
go run main.go
