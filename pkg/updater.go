package pkg

import (
	"errors"
	"github.com/hashicorp/go-version"
	"io"
	"log"
	"net/http"
	"regexp"
)

func isAppUpdateAvailable() (bool, error) {
	const releaseUrl = "https://api.github.com/repos/KatodForAnod/simpleTimeTracker/releases"
	resp, err := http.Get(releaseUrl)
	if err != nil {
		log.Println(err)
		return false, err
	}

	var result []byte
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return false, err
		}
		reg := regexp.MustCompile("\"tag_name\":\".*?\"")
		result = reg.Find(bodyBytes)
	}

	if len(result) == 0 {
		newErr := errors.New("not found")
		log.Println(err)
		return false, newErr
	}

	reg := regexp.MustCompile("[0-9].[0-9].[0-9]")
	ver := reg.Find(result)

	newer, err := isVersionNewer(BackendVersion, string(ver))
	if err != nil {
		log.Println(err)
		return false, err
	}

	return newer, nil
}

func isVersionNewer(oldVersion, newVersion string) (bool, error) {
	v1, err := version.NewVersion(newVersion)
	if err != nil {
		log.Println(err)
		return false, err
	}

	constraints, err := version.NewConstraint("> " + oldVersion)
	if constraints.Check(v1) {
		return true, nil
	}
	return false, nil
}
