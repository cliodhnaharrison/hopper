#!/bin/bash

useradd -m $USERNAME
echo "$USERNAME:$USERNAME" | chpasswd
wetty --base=/"$USERNAME"
