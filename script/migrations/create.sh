#!/bin/bash

PS3="> "

echo 'Select format:'
select format in "unix" "unixNano"
do
    break
done

echo "Enter migration name:"

read name

migrate create -ext sql -dir deployments/migrations/postgres -format $format $name

exit 0