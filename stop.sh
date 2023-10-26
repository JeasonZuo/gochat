#!/bin/sh
docker compose -p tt-chat-backend down
docker rmi $(docker images "tt-chat-backend*" -q)