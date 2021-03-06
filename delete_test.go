package main

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestDeleteReleaseURL(t *testing.T) {
	RegisterTestingT(t)

	info := &Info{
		ID:        123,
		OwnerName: "taichi",
		RepoName:  "tool",
	}

	url := deleteReleaseURL(info)
	Expect(url).To(Equal("https://api.github.com/repos/taichi/tool/releases/123"))
}

func TestDeleteAssetURL(t *testing.T) {
	RegisterTestingT(t)

	info := &Info{
		OwnerName: "taichi",
		RepoName:  "tool",
	}

	url := deleteAssetURL(info, 12345)
	Expect(url).To(Equal("https://api.github.com/repos/taichi/tool/releases/assets/12345"))
}

func TestDeleteTagURL(t *testing.T) {
	RegisterTestingT(t)

	info := &Info{
		TagName:   "v0.1.2",
		OwnerName: "taichi",
		RepoName:  "tool",
	}

	url := deleteTagURL(info)
	Expect(url).To(Equal("https://api.github.com/repos/taichi/tool/git/refs/tags/v0.1.2"))
}
