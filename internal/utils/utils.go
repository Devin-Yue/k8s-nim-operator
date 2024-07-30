/**
# Copyright (c) NVIDIA CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"os"
	"path/filepath"
	"sort"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/rand"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// NvidiaAnnotationHashKey indicates annotation name for last applied hash by the operator
	NvidiaAnnotationHashKey = "nvidia.com/last-applied-hash"
)

// GetFilesWithSuffix returns all files under a given base directory that have a specific suffix
// The operation is performed recursively on subdirectories as well
func GetFilesWithSuffix(baseDir string, suffixes ...string) ([]string, error) {
	var files []string
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		// Error during traversal
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// Skip non suffix files
		base := info.Name()
		for _, s := range suffixes {
			if strings.HasSuffix(base, s) {
				files = append(files, path)
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error traversing directory tree: %w", err)
	}
	return files, nil
}

// BoolPtr returns a pointer to the bool value passed in.
func BoolPtr(v bool) *bool {
	return &v
}

func GetStringHash(s string) string {
	hasher := fnv.New32a()
	if _, err := hasher.Write([]byte(s)); err != nil {
		panic(err)
	}
	return rand.SafeEncodeString(fmt.Sprint(hasher.Sum32()))
}

// MergeMaps merges two maps and ensures no duplicate key-value pairs
func MergeMaps(m1, m2 map[string]string) map[string]string {
	merged := make(map[string]string)

	for k, v := range m1 {
		merged[k] = v
	}

	for k, v := range m2 {
		if _, exists := merged[k]; !exists {
			merged[k] = v
		}
	}

	return merged
}

// MergeEnvVars merges two slices of environment variables, giving precedence to the second slice in case of duplicates
func MergeEnvVars(env1, env2 []corev1.EnvVar) []corev1.EnvVar {
	envMap := make(map[string]corev1.EnvVar)
	for _, env := range env1 {
		envMap[env.Name] = env
	}

	for _, env := range env2 {
		envMap[env.Name] = env
	}

	var mergedEnv []corev1.EnvVar
	for _, env := range envMap {
		mergedEnv = append(mergedEnv, env)
	}

	return mergedEnv
}

// SortKeys recursively sorts the keys of a map to ensure consistent serialization
func SortKeys(obj interface{}) interface{} {
	switch obj := obj.(type) {
	case map[string]interface{}:
		sortedMap := make(map[string]interface{})
		keys := make([]string, 0, len(obj))
		for k := range obj {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			sortedMap[k] = SortKeys(obj[k])
		}
		return sortedMap
	case []interface{}:
		// Check if the slice contains maps and sort them if so
		if len(obj) > 0 {
			if _, ok := obj[0].(map[string]interface{}); ok {
				sort.Slice(obj, func(i, j int) bool {
					iName := obj[i].(map[string]interface{})["name"].(string)
					jName := obj[j].(map[string]interface{})["name"].(string)
					return iName < jName
				})
			}
		}
		for i, v := range obj {
			obj[i] = SortKeys(v)
		}
	}
	return obj
}

// GetResourceHash returns a consistent hash for the given object spec
func GetResourceHash(obj client.Object) string {
	// Convert obj to a map[string]interface{}
	objMap, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	var objData map[string]interface{}
	if err := json.Unmarshal(objMap, &objData); err != nil {
		panic(err)
	}

	// Sort keys to ensure consistent serialization
	sortedObjData := SortKeys(objData)

	// Serialize to JSON
	serialized, err := json.Marshal(sortedObjData)
	if err != nil {
		panic(err)
	}

	// Compute the hash
	hasher := sha256.New()
	hasher.Write(serialized)
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

// IsSpecChanged returns true if the spec has changed between the existing one
// and the new resource spec compared by hash.
func IsSpecChanged(current client.Object, desired client.Object) bool {
	if current == nil && desired != nil {
		return true
	}

	hashStr := GetResourceHash(desired)
	foundHashAnnotation := false

	currentAnnotations := current.GetAnnotations()
	desiredAnnotations := desired.GetAnnotations()

	if currentAnnotations == nil {
		currentAnnotations = map[string]string{}
	}
	if desiredAnnotations == nil {
		desiredAnnotations = map[string]string{}
	}

	for annotation, value := range currentAnnotations {
		if annotation == NvidiaAnnotationHashKey {
			if value != hashStr {
				// Update annotation to be added to resource as per new spec and indicate spec update is required
				desiredAnnotations[NvidiaAnnotationHashKey] = hashStr
				desired.SetAnnotations(desiredAnnotations)
				return true
			}
			foundHashAnnotation = true
			break
		}
	}

	if !foundHashAnnotation {
		// Update annotation to be added to resource as per new spec and indicate spec update is required
		desiredAnnotations[NvidiaAnnotationHashKey] = hashStr
		desired.SetAnnotations(desiredAnnotations)
		return true
	}

	return false
}