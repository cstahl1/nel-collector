// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// nel-collector runs a NEL collector on port 8080, printing out a summary of
// each report that it receives.
package main

import (
	"log"
	"net/http"

	"github.com/google/nel-collector/pkg/collector"
	_ "github.com/google/nel-collector/pkg/core"
)

var defaultConfig = []byte(`
[[processor]]
type = "KeepNelReports"

[[processor]]
type = "DumpReportsAsCLF"
dest = "stdout"
`)

var rootBody = []byte(`
<html>
  <head>
    <title>Network Error Logging collector</title>
  </head>
  <body>
    <h1>Network Error Logging</h1>
    <p>
      This is a collector that can receive
      <a href="https://wicg.github.io/network-error-logging/">Network Error
      Logging</a> reports.
    </p>
  </body>
</html>
`)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write(rootBody)
}

func main() {
	pipeline := &collector.Pipeline{}
	err := pipeline.LoadFromConfig(defaultConfig)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", handleRoot)
	http.Handle("/upload/", pipeline)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
