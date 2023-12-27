#!/bin/sh

cd ../db/migration

goose mysql "user:password@tcp(mysql)/db?charset=utf8mb4&parseTime=true" up
