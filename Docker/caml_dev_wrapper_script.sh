#!/bin/bash

# this is a quick and dirty way to have the caml_dev demo up and running as the default state. 

mkdir ./pom

# a simple server exposing metrics on port 9100
./node_exporter/node_exporter --collector.disable-defaults --collector.textfile --collector.textfile.directory ./pom &

# forever generate metrics data
# NOTE: every 20 seconds is an order of magnitude faster than the current pull interval
cd CAML/src/main/
while true; do ./camldataservice | sed 's/-/_/g' > ../../../pom/caml.metrics.prom ; echo -n "updated at : "; date; sleep 20; done

# some say the script is still running to this day
exit $?