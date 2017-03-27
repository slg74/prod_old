#!/usr/bin/perl

use strict;
use warnings;

sub get_frag_levels {

        my @device;
        my @fraglevel;

        open my $cmd, "mount|", or die "cannot issue command: $!";
        while (<$cmd>) {
                my @line = split;
                if ($_ =~ /storage0/) {
                        my @data = split(' ', `xfs_db -c frag -r $line[0]`);
			$data[6] =~ s/\s*\d+%$//;  # remove percent sign
                        push @device, $line[0];
                        push @fraglevel, sprintf("%.2f", $data[6]); 
                }
        }
        close $cmd;

        my $frag_threshhold = 50.0;

        for (my $i = 0; $i < @device; $i++) {
                if ($fraglevel[$i] > $frag_threshhold) {
                        print "$device[$i] $fraglevel[$i] above threshhold\n";
                } else {
                        print "$device[$i] $fraglevel[$i] below threshhold\n";
                }
        }
}

sub check_mounted_volumes {
        my $type = shift;
        my $count = 0;
        open my $cmd, "mount|" or die "cannot issue mount command: $!";
        while (<$cmd>) {
                my $line = $_;
                $count++, if $line =~ /storage0/;
        }
        close $cmd;

        # check production gen2 & gen3 mounted filesystems.
        if ($type eq "production") {
                if (($count == 6) or ($count == 9)) {
                        print "OK\n";
                        return "OK";
                } else {
                        #print_warning("check_mounted_volumes", "incorrect number of mounted filesystems");
                        print "check_mounted_volumes :: incorrect number of mounted filesystems\n";
                        print "FAIL\n";
                        return "FAIL";
                }
        }
}

get_frag_levels();
check_mounted_volumes("production"); 

