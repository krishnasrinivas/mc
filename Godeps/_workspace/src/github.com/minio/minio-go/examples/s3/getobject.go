// +build ignore

/*
 * Minimal object storage library (C) 2015 Minio, Inc.
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
 */

package main

import (
	"io"
	"log"
	"os"

	"github.com/minio/minio-go"
)

func main() {
	config := minio.Config{
		AccessKeyID:     "YOUR-ACCESS-KEY-HERE",
		SecretAccessKey: "YOUR-PASSWORD-HERE",
		Endpoint:        "https://s3.amazonaws.com",
	}
	s3Client, err := minio.New(config)
	if err != nil {
		log.Fatalln(err)
	}
	reader, stat, err := s3Client.GetObject("mybucket", "myobject")
	if err != nil {
		log.Fatalln(err)
	}

	localfile, err := os.Create("testfile")
	if err != nil {
		log.Fatalln(err)
	}
	defer localfile.Close()

	if _, err = io.CopyN(localfile, reader, stat.Size); err != nil {
		log.Fatalln(err)
	}
}
