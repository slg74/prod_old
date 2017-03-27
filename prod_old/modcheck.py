#!/usr/bin/python

from subprocess import Popen,PIPE,STDOUT
import os

filesystems = [ "opt", "usr", "etc" ]

def run_command(str):
    process = Popen(args=str, stdout=PIPE, shell=True)
    return process.communicate()[0]

def create_sha512sum(fs):
    run_command("find /" + fs + " -type f -exec sha512sum {} > sha512sum_" + fs + ".out +")

if __name__=="__main__":
    for fs in filesystems:
	run_command("cp sha512sum_" + fs + ".out sha512sum_" + fs + ".prev")

    for fs in filesystems:
        print "creating sha512sum for fs: /" + fs
        create_sha512sum(fs)
	print "sha512sum for fs: /" + fs + " COMPLETE."

    for fs in filesystems:
	print("running command: diff sha512sum_" + fs + ".out sha512sum_" + fs + ".prev")
        #run_command("diff sha256sum_" + fs + ".out sha256sum_" + fs + ".prev")
        os.system("diff sha512sum_" + fs + ".out sha512sum_" + fs + ".prev")


