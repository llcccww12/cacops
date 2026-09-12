// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

func keysInt64(m map[int64]struct{}) []int64 {
	var keys = make([]int64, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
func keysString(m map[string]struct{}) []string {
	var keys = make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func valuesRepository(m map[int64]*Repository) []*Repository {
	var values = make([]*Repository, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func valuesDataset(m map[string]*DatasetRegistry) []*DatasetRegistry {
	var values = make([]*DatasetRegistry, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func valuesAimodel(m map[string]*AiModelManage) []*AiModelManage {
	var values = make([]*AiModelManage, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func valuesUser(m map[int64]*User) []*User {
	var values = make([]*User, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func valuesComment(m map[int64]*Comment) []*Comment {
	var values = make([]*Comment, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
func valuesCloudbrain(m map[int64]*Cloudbrain) []*Cloudbrain {
	var values = make([]*Cloudbrain, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
