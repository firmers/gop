/*
 * Copyright (c) 2019, 2025, firmer.tech and/or its affiliates. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client

//go:generate go install github.com/firmers/mesh/client/golang/cmd/babel@master
//go:generate babel -n gop -l dart       -d godoc -w ./dart       -i mpi,kinds --mesh 0.0.10 -e primitive -c cause.dart,status.dart
//go:generate babel -n gop -l java       -d godoc -w ./java       -i mpi,kinds --mesh 0.0.5  -e primitive -c cause.java,status.java --pkg tech.firmer.gop
//go:generate babel -n gop -l javascript -d godoc -w ./javascript -i mpi,kinds --mesh 0.0.16 -e primitive -c cause.ts,status.ts
//go:generate babel -n gop -l python     -d godoc -w ./python     -i mpi,kinds --mesh 0.0.11 -e primitive -c cause.py,status.py     --pkg gop
//go:generate babel -n gop -l rust       -d godoc -w ./rust       -i mpi,kinds --mesh 0.0.11 -e primitive -c cause.rs,status.rs
//go:generate babel -n gop -l golang     -d godoc -w ./           -i mpi,kinds --mesh 0.0.11 -e primitive -c cause.go,status.go     -m 11
