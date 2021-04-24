package config

import "testing"

func TestCleanName(t *testing.T) {
	t.Error(CleanName("@@naAAAAme_a_@"))
}
