// Copyright 2019 go-utils Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package symlink

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/liuchonglin/go-utils/uuidutil"
)

// Replace will do an atomic replacement of a symlink to a new path
func Replace(link, newpath string) error {
	dstDir := filepath.Dir(link)
	uuid, err := uuidutil.NewUUID()
	if err != nil {
		return err
	}
	randStr := uuid.String()
	tmpFile := filepath.Join(dstDir, "tmpfile"+randStr)
	// Create the new symlink before removing the old one. This way, if New()
	// fails, we still have a link to the old tools.
	err = New(newpath, tmpFile)
	if err != nil {
		return fmt.Errorf("cannot create symlink: %s", err)
	}
	// On Windows, symlinks may not be overwritten. We remove it first,
	// and then rename tmpFile
	if _, err := os.Stat(link); err == nil {
		err = os.RemoveAll(link)
		if err != nil {
			return err
		}
	}
	err = os.Rename(tmpFile, link)
	if err != nil {
		return fmt.Errorf("cannot update tools symlink: %v", err)
	}
	return nil
}
