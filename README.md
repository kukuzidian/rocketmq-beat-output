## Beat Output rocketmq

This is a output implementation of [elastic beats](https://github.com/elastic/beats) for support [Filebeat](https://github.com/elastic/beats/tree/master/filebeat), [Metricbeat](https://github.com/elastic/beats/tree/master/metricbeat), [Functionbeat](https://github.com/elastic/beats/tree/master/x-pack/functionbeat), [Winlogbeat](https://github.com/elastic/beats/tree/master/winlogbeat), [Journalbeat](https://github.com/elastic/beats/tree/master/journalbeat), [Auditbeat](https://github.com/elastic/beats/tree/master/auditbeat) to [Apache rocketmq](https://github.com/apache/rocketmq)

### Compatibility
This output is developed and tested using Apache rocketmq-client-go  2.1.0 and Beats 7.9.1

### Download rocketmq-beat-output

```
mkdir -p $GOPATH/src/github.com/kukuzidian/
cd $GOPATH/src/github.com/kukuzidian/
git clone https://github.com/kukuzidian/rocketmq-beat-output
cd rocketmq-beat-output
```

### Build

#### Build Filebeat

Edit main.go file

```
package main

import (
    "os"
    _ "github.com/kukuzidian/rocketmq-beat-output/rocketmq"
    "github.com/elastic/beats/x-pack/filebeat/cmd"
)

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
```

```
go build -o filebeat main.go
```

### Usage

#### Add following configuration to filebeat.yml
```
output.rocketmq:
  namesrvaddr: "172.16.14.128:9876"
  topic: beat_topic
  enabled: true
  codec.json:
    pretty: false
    escape_html: false
```
#### Start filebeat
```
./filebeat modules enable system
./filebeat modules list
./filebeat -c filebeat.yml -e
```

#### Build other beat

```
go build -o metricbeat metricbeat.go
go build -o filebeat filebeat.go
go build -o functionbeat functionbeat.go
go build -o journalbeat journalbeat.go
go build -o auditbeat auditbeat.go
go build -o winlogbeat winlogbeat.go
go build -o packetbeat packetbeat.go
```

### Configurations



### FAQ

