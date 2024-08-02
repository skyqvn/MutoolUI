go build -buildmode=exe -ldflags="-H windowsgui" -o ./app/MutoolUI.exe
copy .\bin\liblcl.dll .\app
copy .\bin\mutool.exe .\app
copy .\docs .\app
mkdir .\app\docs
xcopy .\docs .\app\docs /S /E /Y
