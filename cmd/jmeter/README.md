JMeter CLI

- Run in Windows Powershell with CMD Arguments

  `go run .\cmd\jmeter\main.go {api-category}`

- Run in Mac Terminal

  `go run ./cmd/jmeter/main.go {api-category}`

- Example

For mac

`go run ./cmd/jmeter/main.go 5`

For windows

`go run .\cmd\jmeter\main.go 5`

- Run jmeter command CLI

`jmeter -n -t <path-to-jmx-file> -l <path-to-log-file> -e -o <path-to-report-folder>`

example command

`jmeter -n -t ./rcie/script/jmx/rcie-api-no5.jmx -l ./rcie/log/12130002072023-rcie-api-no5.jtl -e -o ./rcie/report/12130002072023-rcie-api-no5`

Definition of file name

- 12 = hours
- 13 = minutes
- 00 = seconds
- 02 = month
- 07 = day
- 2023 = year

Way to increase jmeter heap size

1. Open jmeter.bat file in bin folder (location: `D:\jmeter\apache-jmeter-5.5\bin`)
2. Allocate size at -Xmx1g -> -Xmx4g (4GB) or more

Temporary command to paste in terminal
`jmeter -n -t ./rcie/script/jmx/rcie-api-no5.jmx -l ./rcie/log/01500002142023-rcie-api-no5.jtl -e -o ./rcie/report/01500002142023-rcie-api-no5`
