cargo-relay
===========

A carbon-relay written in go.


Motivation
-----------
First I wanted to play around with go and I needed a subject for this.
The graphite-project with it's carbon-deamons is of vital interest to me and I would like to have a relay for myself (and others).
The relay should be abel to 
 - forward messages with the ability to smooth (maybe Exponentially Weighted Moving Average) the flow of metrics
 - publish metrics with zeroMQ PUB/SUB pattern to allow other relays to subscribe to this relay
 - transform metrics and forward them to other data stores (e.g. MongoDB, InfluxDB)

Currently it's just a proof of concept. Depending on how much time I can squeeze out of my days.
