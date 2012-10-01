@echo off
set action=build
if "%1" == "clean" set action=clean
for /d %%x in ("src\cmd\*") do cd "%%x" && echo %action%ing %%x ... && go %action% & cd ..
@echo on
