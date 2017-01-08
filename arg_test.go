package pub

import "testing"

func TestNewArg(t *testing.T) {
	configFilePath := "./cmd/pub/config.toml"
	lineChannelSecret := "ld8njsgn42cbtbnkdff8h9ii7jzpsxua"
	lineAccessToken := "ykx7f1ci90dlart11m7h6uzedtu5ymo0"

	actual := NewArg(configFilePath, lineChannelSecret, lineAccessToken)

	if actual.configFilePath != configFilePath {
		t.Errorf("\nExpect is %s\nActual is %s", configFilePath, actual.configFilePath)
	}
	if actual.lineChannelSecret != lineChannelSecret {
		t.Errorf("\nExpect is %s\nActual is %s", lineChannelSecret, actual.lineChannelSecret)
	}
	if actual.lineAccessToken != lineAccessToken {
		t.Errorf("\nExpect is %s\nActual is %s", lineAccessToken, actual.lineAccessToken)
	}
}
