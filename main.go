/*
 * Copyright (C) 2018 The dna Authors
 * This file is part of The dna library.
 *
 * The dna is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The dna is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The dna.  If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	"flag"
	"math/rand"
	"strings"
	"time"

	log4 "github.com/alecthomas/log4go"
	"github.com/dnaproject2/dna-tool/config"
	"github.com/dnaproject2/dna-tool/core"
	_ "github.com/dnaproject2/dna-tool/methods"
)

var (
	Config    string //config file
	LogConfig string //Log config file
	Methods   string //Methods list in cmdline
)

func init() {
	flag.StringVar(&Config, "cfg", "./config.json", "Config of dna-tool")
	flag.StringVar(&LogConfig, "lfg", "./log4go.xml", "Log config of dna-tool")
	flag.StringVar(&Methods, "t", "", "methods to run. use ',' to split methods")
	flag.Parse()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log4.LoadConfiguration(LogConfig)
	defer time.Sleep(time.Second)

	err := config.DefConfig.Init(Config)
	if err != nil {
		log4.Error("DefConfig.Init error:%s", err)
		return
	}

	methods := make([]string, 0)
	if Methods != "" {
		methods = strings.Split(Methods, ",")
	}

	core.OntTool.Start(methods)
}
