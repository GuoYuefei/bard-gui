#!/bin/sh
cat run.pid | xargs kill > /dev/null 2>&1
cat ./bard/run.pid | xargs kill > /dev/null 2>&1