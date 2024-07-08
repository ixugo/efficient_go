/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
)

const (
	Topic    = "test"
	Endpoint = "192.168.31.39:9876"
)

func main() {
	// os.Setenv("mq.consoleAppender.enabled", "true")
	golang.ResetLogger()
	// new producer instance
	producer, err := golang.NewProducer(&golang.Config{
		Endpoint: Endpoint,
		// Group:    GroupName,
		// Region:   Region,
		Credentials: &credentials.SessionCredentials{},
	},
		golang.WithTopics(Topic),
	)
	if err != nil {
		log.Fatal(err)
	}
	if producer == nil {
		panic("producer is nil")
	}
	// start producer
	err = producer.Start()
	if err != nil {
		panic(err)
	}
	// gracefule stop producer
	defer producer.GracefulStop()

	for i := 0; i < 10; i++ {
		// new a message
		msg := &golang.Message{
			Topic: Topic,
			Body:  []byte("this is a message : " + strconv.Itoa(i)),
		}
		// set keys and tag
		msg.SetKeys("a", "b")
		msg.SetTag("ab")
		// send message in sync
		resp, err := producer.Send(context.TODO(), msg)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(resp); i++ {
			fmt.Printf("%#v\n", resp[i])
		}
		// wait a moment
		time.Sleep(time.Second * 1)
	}
}
