package config

import (
	"os"
	"testing"
	"time"
)

func TestEnvString(t *testing.T) {

	t.Log("VAULT_REST_TEST_STR missing")
	{
		if envString("VAULT_REST_TEST_STR", "x") != "x" {
			t.Errorf("envString did not provide default value")
		}
	}

	t.Log("VAULT_REST_TEST_STR present")
	{
		os.Setenv("VAULT_REST_TEST_STR", "y")
		defer os.Unsetenv("VAULT_REST_TEST_STR")

		if envString("VAULT_REST_TEST_STR", "x") != "y" {
			t.Errorf("envString did not obtain env value")
		}
	}
}

func TestEnvBoolean(t *testing.T) {

	t.Log("VAULT_REST_TEST_BOOL missing")
	{
		if envBoolean("VAULT_REST_TEST_BOOL", true) != true {
			t.Errorf("envBoolean did not provide default value")
		}
	}

	t.Log("VAULT_REST_TEST_BOOL present and valid")
	{
		os.Setenv("VAULT_REST_TEST_BOOL", "false")
		defer os.Unsetenv("VAULT_REST_TEST_BOOL")

		if envBoolean("VAULT_REST_TEST_BOOL", true) != false {
			t.Errorf("envBoolean did not obtain env value")
		}
	}

	t.Log("VAULT_REST_TEST_BOOL present and invalid")
	{
		os.Setenv("VAULT_REST_TEST_BOOL", "x")
		defer os.Unsetenv("VAULT_REST_TEST_BOOL")

		if envBoolean("VAULT_REST_TEST_BOOL", true) != true {
			t.Errorf("envBoolean did not obtain fallback to default value")
		}
	}
}

func TestEnvInteger(t *testing.T) {

	t.Log("VAULT_REST_TEST_INT missing")
	{
		if envInteger("VAULT_REST_TEST_INT", 0) != 0 {
			t.Errorf("envInteger did not provide default value")
		}
	}

	t.Log("VAULT_REST_TEST_INT present and valid")
	{
		os.Setenv("VAULT_REST_TEST_INT", "1")
		defer os.Unsetenv("VAULT_REST_TEST_INT")

		if envInteger("VAULT_REST_TEST_INT", 0) != 1 {
			t.Errorf("envInteger did not obtain env value")
		}
	}

	t.Log("VAULT_REST_TEST_INT present and invalid")
	{
		os.Setenv("VAULT_REST_TEST_INT", "x")
		defer os.Unsetenv("VAULT_REST_TEST_INT")

		if envInteger("VAULT_REST_TEST_INT", 0) != 0 {
			t.Errorf("envInteger did not obtain fallback to default value")
		}
	}
}

func TestEnvDuration(t *testing.T) {

	t.Log("VAULT_REST_TEST_DUR missing")
	{
		if envDuration("VAULT_REST_TEST_DUR", time.Second) != time.Second {
			t.Errorf("envDuration did not provide default value")
		}
	}

	t.Log("VAULT_REST_TEST_DUR present and valid")
	{
		os.Setenv("VAULT_REST_TEST_DUR", "2s")
		defer os.Unsetenv("VAULT_REST_TEST_DUR")

		if envDuration("VAULT_REST_TEST_DUR", time.Second) != 2*time.Second {
			t.Errorf("envDuration did not obtain env value")
		}
	}

	t.Log("VAULT_REST_TEST_DUR present and invalid")
	{
		os.Setenv("VAULT_REST_TEST_DUR", "x")
		defer os.Unsetenv("VAULT_REST_TEST_DUR")

		if envDuration("VAULT_REST_TEST_DUR", time.Second) != time.Second {
			t.Errorf("envDuration did not obtain fallback to default value")
		}
	}
}

func TestEnvFilename(t *testing.T) {

	t.Log("VAULT_REST_TEST_FILE missing")
	{
		if envFilename("VAULT_REST_TEST_FILE", "/tmp/stub") != "/tmp/stub" {
			t.Errorf("envFilename did not provide default value")
		}
	}

	t.Log("VAULT_REST_TEST_FILE present and valid")
	{
		os.Setenv("VAULT_REST_TEST_FILE", "/tmp")
		defer os.Unsetenv("VAULT_REST_TEST_FILE")

		if envFilename("VAULT_REST_TEST_FILE", "/tmp/stub") != "/tmp" {
			t.Errorf("envFilename did not obtain env value")
		}
	}

	t.Log("VAULT_REST_TEST_FILE present and invalid")
	{
		os.Setenv("VAULT_REST_TEST_FILE", "/dev/null")
		defer os.Unsetenv("VAULT_REST_TEST_FILE")

		if envFilename("VAULT_REST_TEST_FILE", "/tmp/stub") != "/tmp/stub" {
			t.Errorf("envFilename did not obtain fallback to default value")
		}
	}
}
