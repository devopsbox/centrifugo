An attempt to move Centrifuge from Python to Go

Why:

* performance
* no need to launch several processes to utilize several cores
* static binary

What's left:

* redis engine API channel
* metrics
* tests
* deploy (Dockerfile, rpm) 
* documentation
* compare with python version (cpu, memory, performance etc)