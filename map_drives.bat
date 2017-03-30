@echo off

:: unmap currently mapped network drives
if exist Z:\ (
	net use Z: /delete
)

if exist Y:\ (
	net use Y: /delete
)

if exist X:\ (
	net use X: /delete
)

::if exist W:\ (
::	net use W: /delete
::)

net use Z: \\10.139.118.12\i acc33$$ /user:accessdx\adx.lab /persistent:yes
net use Y: \\10.139.118.12\g acc33$$ /user:accessdx\adx.lab /persistent:yes
net use X: \\10.139.118.12\e acc33$$ /user:accessdx\adx.lab /persistent:yes
::net use W: \\10.139.118.12\h acc33$$ /user:accessdx\adx.lab /persistent:yes
