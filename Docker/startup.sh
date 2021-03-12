#!/bin/bash

useradd -m $USERNAME
echo "$USERNAME:pass123" | chpasswd
wetty --base=/"$USERNAME"
