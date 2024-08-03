go build -buildmode=exe -ldflags="-H windowsgui -s -w" -o ./app/MutoolUI.exe
copy .\bin\liblcl.dll .\app
copy .\bin\mutool.exe .\app
mkdir .\app\docs
xcopy .\docs .\app\docs /S /E /Y
