#!/bin/bash
set -ex
make
./Config --config=config.conf
