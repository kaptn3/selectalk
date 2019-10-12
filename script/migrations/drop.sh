#!/bin/bash

migrate -source file://deployments/migrations/postgres -database postgres://selectel:selectel@localhost:5432/selectel?sslmode=disable drop