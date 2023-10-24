@echo off

set "taskname=Waiting tasks scheduler"
set "workingdir=%cd%"
set "scriptname=todo-txt-waiting.bat"

schtasks /create /tn "%taskname%" /tr "cmd /c cd /d %workingdir% && %scriptname%" /sc daily /st 09:00

echo Scheduled task created successfully.
