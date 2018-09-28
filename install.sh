#!/bin/bash

sudo mkdir -p /etc/rofi-snippet/
sudo cp ./config.toml /etc/rofi-snippet/
go install .
