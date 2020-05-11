/*
 * Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package ionhash

type scalarSerializer struct {
	serializer
}

func newScalarSerializer(hashFunction IonHasher, depth int) Serializer {
	return &scalarSerializer{serializer{hashFunction: hashFunction, depth: depth}}
}

func (scalarSerializer *scalarSerializer) scalar(ionValue interface{}) {
	panic("implement me")
}

func (scalarSerializer *scalarSerializer) stepOut() {
	panic("implement me")
}
