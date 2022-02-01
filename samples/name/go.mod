module github.com/rwxrob/cmdbox/samples/name

go 1.17

replace github.com/rwxrob/cmdbox/samples/name/chinese => ./chinese

replace github.com/rwxrob/cmdbox/samples/name/russian => ./russian

require (
	github.com/rwxrob/cmdbox v0.5.0
	github.com/rwxrob/cmdbox/samples/name/chinese v0.0.0-00010101000000-000000000000
	github.com/rwxrob/cmdbox/samples/name/russian v0.0.0-00010101000000-000000000000
)

require gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
