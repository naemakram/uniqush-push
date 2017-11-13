/*
 * Copyright 2011 Nan Deng
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"io"    
	"github.com/naemakram/uniqush-push/srv"
)

var uniqushPushConfFlags = flag.String("config", "/etc/uniqush/uniqush-push.conf", "Config file path")
var uniqushPushShowVersionFlag = flag.Bool("version", false, "Version info")

var uniqushPushVersion = "uniqush-push 2.4.0"

func installPushServices() {

	// open output file
	fo, err := os.Create("/tmp/output.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
	if err := fo.Close(); err != nil {
			panic(err)
	}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
	// read a chunk
	n, err := fi.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}
	if n == 0 {
		break
	}

	// write a chunk
	if _, err := fo.Write(buf[:n]); err != nil {
		panic(err)
	}
	}
	
	srv.InstallGCM()
	srv.InstallFCM()
	srv.InstallAPNS()
	srv.InstallADM()
}

func main() {
	flag.Parse()
	if *uniqushPushShowVersionFlag {
		fmt.Printf("%v\n", uniqushPushVersion)
		return
	}
	runtime.GOMAXPROCS(runtime.NumCPU() + 1)
	installPushServices()

	err := Run(*uniqushPushConfFlags, uniqushPushVersion)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot start: %v\n", err)
	}
	
	// open output file
	fo, err := os.Create("/tmp/output.txt")
	if err != nil {
	panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
	if err := fo.Close(); err != nil {
		panic(err)
	}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
	// read a chunk
	n, err := fi.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}
	if n == 0 {
		break
	}

	// write a chunk
	if _, err := fo.Write(buf[:n]); err != nil {
		panic(err)
	}
	}
}
