#!/usr/bin/env bats

bats_load_library test_helper

@test "bash -c injects set -euo pipefail" {
    # This should fail on 'false' due to set -e
    trun bash -c 'echo before; false; echo after'
    assert_failure
    assert_output "before"
    refute_output --partial "after"
}

@test "bash -c with pipefail catches pipe failures" {
    # Without pipefail, this would succeed. With pipefail, it fails.
    trun bash -c 'false | cat; echo "should not print"'
    assert_failure
    refute_output --partial "should not print"
}

@test "bash -c with unset variable fails" {
    # set -u should catch undefined variables
    trun bash -c 'echo $UNDEFINED_VAR_12345; echo after'
    assert_failure
    refute_output --partial "after"
}

@test "bash -c skips injection if already has set -e" {
    # Should not double-inject
    trun bash -c 'set -e; echo test'
    assert_success
    assert_output "test"
}

@test "bash -c skips injection if already has set -euo pipefail" {
    trun bash -c 'set -euo pipefail; echo test'
    assert_success
    assert_output "test"
}

@test "bash without -c passes through" {
    trun bash --version
    assert_success
    assert_output --partial "GNU bash"
}

@test "bash -l (login shell) passes through" {
    trun bash -l -c 'echo login-test'
    assert_success
    assert_output --partial "login-test"
}

@test "bash -c with successful commands works" {
    trun bash -c 'echo hello; echo world'
    assert_success
    assert_output --partial "hello"
    assert_output --partial "world"
}
