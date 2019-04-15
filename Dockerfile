FROM mcr.microsoft.com/windows/nanoserver:1809

COPY out/client.exe /client.exe
COPY out/hello.exe /hello.exe

ENTRYPOINT client.exe
