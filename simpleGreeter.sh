#!/bin/bash

Help()
{
   echo "Simple greeter creates a short greeting"
   echo
   echo "Syntax: simpleGreeter [-n|t]"
   echo 
   echo "n     Name of person to greet"
   echo "t     greeting type [HELLO, GOOD_MORNING, GOOD_NIGHT, HAPPY_BIRTHDAY, MERRY_XMAS]"
   echo
}

while getopts :hn:t: opt
do
    case "${opt}" in
        h) Help
           exit;;
        n) name=${OPTARG};;
        t) greeting_type=${OPTARG};;
    esac
done


if [ -z "$name" ]; then
    echo 'name is required' >&2
    exit 1
fi

if [ -z "$greeting_type" ]; then
    echo 'type is required' >&2
    exit 1
fi

export RECEPIENT=$name
export GREETING_TYPE=$greeting_type

docker-compose up --quiet-pull greeting_client
docker-compose stop > /dev/null 2>&1 