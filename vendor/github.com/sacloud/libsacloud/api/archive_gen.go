// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

/************************************************
  generated by IDE. for [ArchiveAPI]
************************************************/

import (
	"github.com/sacloud/libsacloud/sacloud"
)

/************************************************
   To support fluent interface for Find()
************************************************/

// Reset 検索条件リセット
func (api *ArchiveAPI) Reset() *ArchiveAPI {
	api.reset()
	return api
}

// Offset オフセット
func (api *ArchiveAPI) Offset(offset int) *ArchiveAPI {
	api.offset(offset)
	return api
}

// Limit リミット
func (api *ArchiveAPI) Limit(limit int) *ArchiveAPI {
	api.limit(limit)
	return api
}

// Include 取得する項目
func (api *ArchiveAPI) Include(key string) *ArchiveAPI {
	api.include(key)
	return api
}

// Exclude 除外する項目
func (api *ArchiveAPI) Exclude(key string) *ArchiveAPI {
	api.exclude(key)
	return api
}

// FilterBy 任意項目でのフィルタ(部分一致)
func (api *ArchiveAPI) FilterBy(key string, value interface{}) *ArchiveAPI {
	api.filterBy(key, value, false)
	return api
}

// FilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *ArchiveAPI) FilterMultiBy(key string, value interface{}) *ArchiveAPI {
	api.filterBy(key, value, true)
	return api
}

// WithNameLike 名称条件
func (api *ArchiveAPI) WithNameLike(name string) *ArchiveAPI {
	return api.FilterBy("Name", name)
}

// WithTag タグ条件
func (api *ArchiveAPI) WithTag(tag string) *ArchiveAPI {
	return api.FilterBy("Tags.Name", tag)
}

// WithTags タグ(複数)条件
func (api *ArchiveAPI) WithTags(tags []string) *ArchiveAPI {
	return api.FilterBy("Tags.Name", []interface{}{tags})
}

// WithSizeGib アーカイブサイズ条件
func (api *ArchiveAPI) WithSizeGib(size int) *ArchiveAPI {
	api.FilterBy("SizeMB", size*1024)
	return api
}

// WithSharedScope 共有スコープ条件
func (api *ArchiveAPI) WithSharedScope() *ArchiveAPI {
	api.FilterBy("Scope", "shared")
	return api
}

// WithUserScope ユーザースコープ条件
func (api *ArchiveAPI) WithUserScope() *ArchiveAPI {
	api.FilterBy("Scope", "user")
	return api
}

// SortBy 任意項目でのソート指定
func (api *ArchiveAPI) SortBy(key string, reverse bool) *ArchiveAPI {
	api.sortBy(key, reverse)
	return api
}

// SortByName 名称でのソート
func (api *ArchiveAPI) SortByName(reverse bool) *ArchiveAPI {
	api.sortByName(reverse)
	return api
}

// SortBySize サイズでのソート
func (api *ArchiveAPI) SortBySize(reverse bool) *ArchiveAPI {
	api.sortBy("SizeMB", reverse)
	return api
}

/************************************************
   To support Setxxx interfaces for Find()
************************************************/

// SetEmpty 検索条件リセット
func (api *ArchiveAPI) SetEmpty() {
	api.reset()
}

// SetOffset オフセット
func (api *ArchiveAPI) SetOffset(offset int) {
	api.offset(offset)
}

// SetLimit リミット
func (api *ArchiveAPI) SetLimit(limit int) {
	api.limit(limit)
}

// SetInclude 取得する項目
func (api *ArchiveAPI) SetInclude(key string) {
	api.include(key)
}

// SetExclude 除外する項目
func (api *ArchiveAPI) SetExclude(key string) {
	api.exclude(key)
}

// SetFilterBy 任意項目でのフィルタ(部分一致)
func (api *ArchiveAPI) SetFilterBy(key string, value interface{}) {
	api.filterBy(key, value, false)
}

// SetFilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *ArchiveAPI) SetFilterMultiBy(key string, value interface{}) {
	api.filterBy(key, value, true)
}

// SetNameLike 名称条件
func (api *ArchiveAPI) SetNameLike(name string) {
	api.FilterBy("Name", name)
}

// SetTag タグ条件
func (api *ArchiveAPI) SetTag(tag string) {
	api.FilterBy("Tags.Name", tag)
}

// SetTags タグ(複数)条件
func (api *ArchiveAPI) SetTags(tags []string) {
	api.FilterBy("Tags.Name", []interface{}{tags})
}

// SetSizeGib アーカイブサイズ条件
func (api *ArchiveAPI) SetSizeGib(size int) {
	api.FilterBy("SizeMB", size*1024)
}

// SetSharedScope 共有スコープ条件
func (api *ArchiveAPI) SetSharedScope() {
	api.FilterBy("Scope", "shared")
}

// SetUserScope ユーザースコープ条件
func (api *ArchiveAPI) SetUserScope() {
	api.FilterBy("Scope", "user")
}

// SetSortBy 任意項目でのソート指定
func (api *ArchiveAPI) SetSortBy(key string, reverse bool) {
	api.sortBy(key, reverse)
}

// SetSortByName 名称でのソート
func (api *ArchiveAPI) SetSortByName(reverse bool) {
	api.sortByName(reverse)
}

// SetSortBySize サイズでのソート
func (api *ArchiveAPI) SetSortBySize(reverse bool) {
	api.sortBy("SizeMB", reverse)
}

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

// Create 新規作成
func (api *ArchiveAPI) Create(value *sacloud.Archive) (*sacloud.Archive, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.create(api.createRequest(value), res)
	})
}

// Read 読み取り
func (api *ArchiveAPI) Read(id sacloud.ID) (*sacloud.Archive, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.read(id, nil, res)
	})
}

// Update 更新
func (api *ArchiveAPI) Update(id sacloud.ID, value *sacloud.Archive) (*sacloud.Archive, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.update(id, api.createRequest(value), res)
	})
}

// Delete 削除
func (api *ArchiveAPI) Delete(id sacloud.ID) (*sacloud.Archive, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.delete(id, nil, res)
	})
}

// New 作成用パラメータ作成
func (api *ArchiveAPI) New() *sacloud.Archive {
	return &sacloud.Archive{}
}

/************************************************
  Inner functions
************************************************/

func (api *ArchiveAPI) setStateValue(setFunc func(*sacloud.Request)) *ArchiveAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

func (api *ArchiveAPI) request(f func(*sacloud.Response) error) (*sacloud.Archive, error) {
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.Archive, nil
}

func (api *ArchiveAPI) createRequest(value *sacloud.Archive) *sacloud.Request {
	req := &sacloud.Request{}
	req.Archive = value
	return req
}