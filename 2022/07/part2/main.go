package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	DiskSize = int64(70000000)
	Goal     = int64(30000000)
)

func main() {
	top := process(strings.Split(Input, "\n"))
	totalSize := getSizes(top)

	fmt.Println(fmt.Sprintf("Total size: %d", totalSize))

	available := DiskSize - totalSize
	fmt.Println(fmt.Sprintf("Available: %d", available))

	neededToDelete := Goal - available

	fmt.Println(fmt.Sprintf("Needed to delete: %d", neededToDelete))

	smallestDeletable := DiskSize

	for _, size := range allSizes {
		if size >= neededToDelete {
			if smallestDeletable > size {
				smallestDeletable = size
			}
		}
	}

	fmt.Println(fmt.Sprintf("Size of folder to delete: %d", smallestDeletable))
}

var allSizes = make(map[string]int64)

type dir struct {
	Name    string
	Parent  *dir
	SubDirs map[string]*dir
	Files   map[string]int64
}

func newDir(name string) *dir {
	return &dir{
		Name:    name,
		SubDirs: make(map[string]*dir),
		Files:   make(map[string]int64),
	}
}

func (d *dir) addSubDir(s *dir) {
	_, ok := d.SubDirs[s.Name]
	if ok {
		return
	}
	s.Parent = d
	d.SubDirs[s.Name] = s
}

func (d *dir) getPath() string {
	if d == nil {
		return ""
	}

	ret := d.Name

	current := d
	for true {
		if current.Parent == nil {
			break
		}
		ret = fmt.Sprintf("%s/%s", current.Parent.Name, ret)
		current = current.Parent
	}
	return ret
}

func getSizes(d *dir) int64 {
	totalSize := int64(0)
	for _, size := range d.Files {
		totalSize += size
	}

	for _, sd := range d.SubDirs {
		totalSize += getSizes(sd)
	}

	allSizes[d.getPath()] = totalSize

	return totalSize
}

func process(in []string) *dir {
	top := newDir("")
	current := top

	for i := 0; i < len(in); i++ {
		// fmt.Println(fmt.Sprintf("%d: %s", i, in[i]))
		parts := strings.Split(in[i], " ")

		switch parts[0] {
		case "$":
			switch parts[1] {
			case "cd":
				switch parts[2] {
				case "..":
					current = current.Parent
					// fmt.Println(fmt.Sprintf("CDing to %s", current.name))
				default:
					c, ok := current.SubDirs[parts[2]]
					if !ok {
						c = newDir(parts[2])
						current.addSubDir(c)
					}
					current = c
					// fmt.Println(fmt.Sprintf("CDing to %s", current.name))
				}
			case "ls":
				// Don't need to do anything
			}
			continue
		case "dir":
			current.addSubDir(newDir(parts[1]))
			// fmt.Println(fmt.Sprintf("Adding subdir %s", parts[1]))
		default:
			size, _ := strconv.ParseInt(parts[0], 10, 64)
			current.Files[parts[1]] = size
			// fmt.Println(fmt.Sprintf("Adding file %s", parts[1]))
		}
	}

	return top
}
