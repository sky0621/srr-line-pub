package pub

import "testing"

func TestNewArg(t *testing.T) {
	configFilePath := "./cmd/pub/config.toml"
	awsAccessKeyID := "j22frhd0ja4j587y36p9p38o5s2d8rx7"
	awsSecretAccessKey := "ro1cyz1w6mol1kbwkwt8jo7jet5gxz7z"
	lineChannelSecret := "ld8njsgn42cbtbnkdff8h9ii7jzpsxua"
	lineAccessToken := "ykx7f1ci90dlart11m7h6uzedtu5ymo0"

	actual := NewArg(configFilePath, awsAccessKeyID, awsSecretAccessKey, lineChannelSecret, lineAccessToken)

	if actual.configFilePath != configFilePath {
		t.Errorf("\nExpect is %s\nActual is %s", configFilePath, actual.configFilePath)
	}
	if actual.awsAccessKeyID != awsAccessKeyID {
		t.Errorf("\nExpect is %s\nActual is %s", awsAccessKeyID, actual.awsAccessKeyID)
	}
	if actual.awsSecretAccessKey != awsSecretAccessKey {
		t.Errorf("\nExpect is %s\nActual is %s", awsSecretAccessKey, actual.awsSecretAccessKey)
	}
	if actual.lineChannelSecret != lineChannelSecret {
		t.Errorf("\nExpect is %s\nActual is %s", lineChannelSecret, actual.lineChannelSecret)
	}
	if actual.lineAccessToken != lineAccessToken {
		t.Errorf("\nExpect is %s\nActual is %s", lineAccessToken, actual.lineAccessToken)
	}
}
