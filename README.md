# Unreliable

Testing replication between two data stores over an unreliable link.

Top represents a read only data source
Bottom a data sink with a fixed schema (so arbitray data cannot be written to it)
Middle represents an unreliable transport mechanism which can discard, delay or duplicate the data it is supposed to be delivering
Home and Away represent eh process which need to synchronise Bottom with Top each has access to one data store and can comunicate with the other over Middle
