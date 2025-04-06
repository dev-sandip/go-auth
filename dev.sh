#!/bin/bash

# This script will help run the development environment
# with Air for live reloading and Templ integration

# Make sure we have all dependencies
echo "Installing/updating dependencies..."
go mod tidy

# Generate initial Templ files
echo "Generating initial Templ files..."
templ generate

# Run Air for live reloading
echo "Starting Air for live reloading..."
air