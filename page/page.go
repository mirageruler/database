// Copyright 2025 macbook16
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

package page

const PageSize = 8192 // 8KB

type PageHeader struct {
	PageID    uint32 // Unique page identifier
	LSN       uint64 // Log Sequence Number (for WAL)
	FreeSpace uint16 // Free space available in the page
}

type Page struct {
	Header PageHeader
	Data   [PageSize - 16]byte // Space for storing row data
}
