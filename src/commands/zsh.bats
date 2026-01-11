#!/usr/bin/env bats

bats_load_library test_helper

@test "zsh -c injects set -euo pipefail" {
    # This should fail on 'false' due to set -e
    trun zsh -c 'echo before; false; echo after'
    assert_failure
    assert_output "before"
    refute_output --partial "after"
}

@test "zsh -c with pipefail catches pipe failures" {
    # Without pipefail, this would succeed. With pipefail, it fails.
    trun zsh -c 'false | cat; echo "should not print"'
    assert_failure
    refute_output --partial "should not print"
}

@test "zsh -c with unset variable fails" {
    # set -u should catch undefined variables
    trun zsh -c 'echo $UNDEFINED_VAR_12345; echo after'
    assert_failure
    refute_output --partial "after"
}

@test "zsh -c skips injection if already has set -e" {
    # Should not double-inject
    trun zsh -c 'set -e; echo test'
    assert_success
    assert_output "test"
}

@test "zsh -c skips injection if already has set -euo pipefail" {
    trun zsh -c 'set -euo pipefail; echo test'
    assert_success
    assert_output "test"
}

@test "zsh without -c passes through" {
    trun zsh --version
    assert_success
    assert_output --partial "zsh"
}

@test "zsh -l (login shell) passes through" {
    trun zsh -l -c 'echo login-test'
    assert_success
    assert_output --partial "login-test"
}

@test "zsh -c with successful commands works" {
    trun zsh -c 'echo hello; echo world'
    assert_success
    assert_output --partial "hello"
    assert_output --partial "world"
}
