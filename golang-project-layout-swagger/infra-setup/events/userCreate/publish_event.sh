#!/bin/bash
# If kafkacat does not installed, please install with brew --> brew install kafkacat

#local
kcat -P -b localhost:9092 -t folksdev.user-created.0 event.json