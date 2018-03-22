#!/usr/bin/python

import os
import subprocess
import socket
import sys

from subprocess import Popen, PIPE, STDOUT

OK = '\033[92m'
WARN = '\033[93m'
FAIL = '\033[91m'
RESET = '\033[0m'

hostname = socket.gethostname().rstrip()

def run_command(s):
    process = Popen(args=s, stdout=PIPE, shell=True)
    return process.communicate()[0]


interface_switchport = {}

def get_switchport_details():

    em1_port = run_command("lldpcli show neighbors | grep -A6 em1 | awk '{NF} END {print $4}'")
    em2_port = run_command("lldpcli show neighbors | grep -A6 em2 | awk '{NF} END {print $4}'")


    print ("%sOK   - %s - Port    - %s - %s." % (OK, hostname, "em1", em1_port.rstrip()))
    print ("%sOK   - %s - Port    - %s - %s." % (OK, hostname, "em2", em2_port.rstrip()))

    interface_switchport["em1"] = em1_port.split("/")[2]
    interface_switchport["em2"] = em2_port.split("/")[2]


def cabling_check():
    if interface_switchport["em1"] == interface_switchport["em2"]:
        print ("%sOK   - %s - em1 and em2 are cabled to port %s \
               on switch A and switch B." %
               (OK, hostname, interface_switchport["em1"].rstrip()))
    else:
        print ("%sFAIL - %s - em1 and em2 are cabled to different \
               switchports." % (FAIL, hostname))


def main():
    get_switchport_details()
    cabling_check()

if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        sys.exit()
