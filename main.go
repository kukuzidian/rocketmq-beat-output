/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

import (
    //"github.com/elastic/beats/v7/filebeat/cmd"
    _ "github.com/kukuzidian/rocketmq-beat-output/rocketmq"
    //inputs "github.com/elastic/beats/v7/filebeat/input/default-inputs"
    "github.com/elastic/beats/x-pack/filebeat/cmd"
    "os"

    "fmt"
)

func main() {
    fmt.Println("-----start")

    if err := cmd.RootCmd.Execute(); err != nil {
       os.Exit(1)
    }
    fmt.Println("-----end")

    /*p,_ := rocketmq.NewProducer(producer.WithNameServer(primitive.NamesrvAddr{"172.16.14.128:9876"}),
        producer.WithRetry(3),
        producer.WithGroupName("testest"))

    p.Start()

    for i:=1;i<=2000;i++ {
        msg := &primitive.Message{
            Topic: "beat_topic",
            Body:  []byte("serializedEvent1230000000000000000--"+string(i)),
        }
        res, err := p.SendSync(context.Background(), msg)

        if err != nil {
            fmt.Printf("send message error: %s\n", err)
        } else {
            fmt.Printf("send message success: result=%s\n", res.String())
        }
    }

    p.Shutdown()*/
}
