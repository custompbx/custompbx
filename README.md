#CustomPBX
-----------------  
**CustomPBX** (under development) is API server and Web GUI for [FreeSwitch](https://github.com/signalwire/freeswitch).
Providing pure FreeSWITCH experience and can be installed on existing systems with import existing configuration.
Built in the single binary file.

The project is under development, not tested well and can have security issues.  

Backend created with Golang v.1.19 and frontend with Angular v.15.

Requirements:  
OS Linux (amd64)  
FreeSWITCH  
Postgres

To build run:  
``make install``  
To rebuild after resolved dependencies:  
``make build``

Built binary file can be found in ``bin/`` directory and can be used according [Documentation](https://github.com/CustomPBX/doc)