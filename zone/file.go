package zone

import (
	"log"
	"os"
	"regexp"
)

type Files []os.FileInfo

func NewZoneFiles(directory string, filter ...string) (Files, error) {
	var fileFilterRegex *regexp.Regexp

	// Regex definition to parse SOA record
	zoneNameRegex := regexp.MustCompile(`(?m)^([^$].*?)\.\s+(?:\d*\s+)?IN\s+SOA\s+`)

	if len(filter) == 1 {
		fileFilterRegex = regexp.MustCompile(filter[0])
	}

	// zoneList initialization
	zoneFiles := Files{}

	// Getting files list from directory
	fileList, err := os.ReadDir(directory)
	if err != nil {
		return Files{}, err
	}

	// Checking all files in the directory
	for _, file := range fileList {
		// Directories are ignored
		if file.IsDir() {
			continue
		}

		if fileFilterRegex != nil {
			if !fileFilterRegex.MatchString(file.Name()) {
				continue
			}
		}

		// Checking if file meet RFC1034 requirement and getting zone name
		// TODO: implement processing $ORIGIN / $TTL / $GENERATE and $INCLUDE
		if zoneData, err := os.ReadFile(directory + file.Name()); err == nil {
			match := zoneNameRegex.FindSubmatch(zoneData)
			// SOA Record found appending zone to the ZoneList
			if len(match) == 2 {
				fInfo, _ := file.Info()
				zoneFiles = append(zoneFiles, fInfo)
			} else {
				log.Printf("SOA Record not found file (%s) ignored", file.Name())
			}
		} else {
			log.Printf("Error getting file content for %s --> %v", file.Name(), err)
		}
	}

	return zoneFiles, nil
}
