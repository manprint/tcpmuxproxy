#!/bin/sh

if [ -n "$BASIC_AUTH" ]; then
    echo "Starting proxytunnel with Basic Authentication..."
    echo "Command: proxytunnel -p $PROXY_HOST -d $PROXY_REMOTE -a $LOCAL_PORT -P $BASIC_AUTH"
    proxytunnel -p $PROXY_HOST -d $PROXY_REMOTE -a $LOCAL_PORT -P $BASIC_AUTH
else
    echo "Starting proxytunnel..."
    echo "Command: proxytunnel -p $PROXY_HOST -d $PROXY_REMOTE -a $LOCAL_PORT"
    proxytunnel -p $PROXY_HOST -d $PROXY_REMOTE -a $LOCAL_PORT
fi