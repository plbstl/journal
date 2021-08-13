/*
Copyright © 2021 Paul Ebose <paulebose@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package internal

import (
	"bufio"
	"errors"
	"io"
	"os"
)

// FileExists checks if a file already exists in the directory.
func FileExists(item, filename, dir string) (bool, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return false, errors.New("cannot read " + item + "s directory: " + dir)
	}
	for _, file := range files {
		if !file.IsDir() {
			if file.Name() == filename {
				return true, nil
			}
		}
	}
	return false, nil
}

// parseLines parses a NOTE file line by line.
// It returns a slice of three strings.
//
// The first string is the author's name,
// which is gotten from the first line.
//
// The second string is the note's title,
// which is gotten from the second line.
//
// The third string is the note's body,
// which is gotten from the remaining lines.
func parseLines(fp string) ([]string, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var results []string
	var readLength int64
	currentLine := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		currentLine++
		if currentLine > 2 {
			break
		}
		line := scanner.Bytes()
		readLength += int64(len(line))
		results = append(results, string(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	maxSize := stat.Size() - readLength
	var chunkSize int64 = 320
	noteBody := make([]byte, chunkSize)
	const lastChar, newline = 1, 1

	for maxSize >= 0 {
		_, err = f.ReadAt(noteBody, readLength+lastChar+newline)
		if err != nil && err != io.EOF {
			return nil, err
		}
		maxSize -= chunkSize
	}

	results = append(results, string(noteBody))
	return results, nil
}
