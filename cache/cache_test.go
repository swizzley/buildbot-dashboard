package cache

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"

	gc "gopkg.in/check.v1"
)

var _ = gc.Suite(&CacheSuite{})

type CacheSuite struct{}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gc.TestingT(t) }

func (s *CacheSuite) SetUpTest(c *gc.C) {
	usr, err := user.Current()
	if err != nil {
		c.Error(err)
	}

	os.Remove(usr.HomeDir + string(filepath.Separator) + InternalCacheFolder)
}

func (s *CacheSuite) TestSetCacheShouldCreateFile(c *gc.C) {
	cache := NewCache(10)
	cache.SetCache("test", []byte("test content"))

	_, err := os.Stat(cache.GetPath() + "test")
	c.Check(err, gc.IsNil)
}

func (s *CacheSuite) TestNonExistingKeyShouldReturnError(c *gc.C) {
	cache := NewCache(10)
	cache.SetCache("test", []byte("test content"))

	data, err := cache.GetCache("test")

	c.Check(err, gc.IsNil)
	c.Check(string(data), gc.Equals, "test content")
}

func (s *CacheSuite) TestExpiredTimeShouldReturnError(c *gc.C) {
	cache := NewCache(-1)
	cache.SetCache("test", []byte("test content"))

	_, err := cache.GetCache("test")
	c.Check(err, gc.NotNil)
}
