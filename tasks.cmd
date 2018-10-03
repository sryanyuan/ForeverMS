@rem param1 GOPATH
@rem param2 build path
@rem param3 action

@SET GOPATH=%1
@CD %2

@if "%4"=="" (
@go %3
) else (
@go %3 -o %4
)

@rem succeed or failed
@if %errorlevel%==0 (echo %3 success) else (echo %3 failed)